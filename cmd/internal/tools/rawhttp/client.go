package rawhttp

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

// Client is a client for making raw http requests with go
type Client struct {
	Options  *Options
	jar      *cookiejar.Jar
	proxyURL *url.URL
}

// NewClient creates a new rawhttp client with provided options
func NewClient(options *Options) *Client {
	jar, _ := cookiejar.New(nil)
	client := &Client{
		Options: options,
		jar:     jar,
	}
	return client
}

// Get makes a GET request to a given URL
func (c *Client) Get(url string) (*http.Response, error) {
	return c.DoRaw("GET", url, "", nil, nil)
}

// DoRaw does a raw request with some configuration
func (c *Client) DoRaw(method, url, uripath string, headers map[string][]string, body io.Reader) (*http.Response, error) {
	rs := &RedirectStatus{
		FollowRedirects: true,
		MaxRedirects:    c.Options.MaxRedirects,
	}
	return c.send(method, url, headers, body, rs, c.Options)
}

func (c *Client) send(method, uri string, headers map[string][]string, body io.Reader, redirectStatus *RedirectStatus, options *Options) (*http.Response, error) {
	protocol := "http"
	if strings.HasPrefix(strings.ToLower(uri), "https://") {
		protocol = "https"
	}
	if options.Proxy != "" {
		var err error
		c.proxyURL, err = url.Parse(options.Proxy)
		if err != nil {
			return nil, err
		}
	}
	httpClient := &http.Client{
		Jar: c.jar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) > redirectStatus.MaxRedirects {
				return errors.New(fmt.Sprintf("已达到最大重定向数: %v", redirectStatus.MaxRedirects))
			}
			if req.Response != nil {
				redirectUrl := headerValue(req.Response.Header, "Location")
				if redirectUrl != "" {
					return http.ErrUseLastResponse
				}
			}
			return nil
		},
		Transport: &http.Transport{
			//TLSClientConfig: &tls.Config{InsecureSkipVerify: true },
			Proxy: http.ProxyURL(c.proxyURL),
		},
	}
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return nil, err
	}
	for key, val := range headers {
		req.Header.Set(key, val[0])
	}
	host := req.Host
	if options.AutomaticHostHeader {
		req.Header.Set("Host", fmt.Sprintf(" %s", host))
	}
	req.Header.Set("User-Agent", "energy raw http/1.0.0")
	req.Header.Set("Accept", "*/*")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	code := resp.StatusCode
	var isRedirect = func() bool {
		if code == REDIRECTION_NOT_MODIFIED {
			return false
		}
		return code >= REDIRECTION_MULTIPLE_CHOICES && code < CLIENT_ERROR_BAD_REQUEST
	}
	// 301 302 303 307 308
	if isRedirect() && redirectStatus.FollowRedirects && redirectStatus.Current <= redirectStatus.MaxRedirects {
		_, err := io.Copy(io.Discard, resp.Body)
		if err := firstErr(err, resp.Body.Close()); err != nil {
			return nil, err
		}
		redirectUrl := headerValue(resp.Header, "Location")
		if strings.HasPrefix(redirectUrl, "/") {
			redirectUrl = fmt.Sprintf("%s://%s%s", protocol, host, redirectUrl)
		}
		redirectStatus.Current++
		return c.send(method, redirectUrl, headers, body, redirectStatus, options)
	}
	return resp, nil
}

func firstErr(err1, err2 error) error {
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	return nil
}

func headerValue(headers map[string][]string, key string) string {
	key = strings.ToLower(key)
	for tmpKey, vals := range headers {
		tmpKey = strings.ToLower(tmpKey)
		if tmpKey == key {
			return strings.Join(vals, " ")
		}
	}
	return ""
}

// RedirectStatus is the current redirect status for the request
type RedirectStatus struct {
	FollowRedirects bool
	MaxRedirects    int
	Current         int
}
