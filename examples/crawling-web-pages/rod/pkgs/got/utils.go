package got

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"math/big"
	"mime"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"text/template"
	"time"
)

// Context helper
type Context struct {
	context.Context
	Cancel func()
}

// Utils for commonly used methods
type Utils struct {
	Testable
}

// Fatal is the same as [testing.common.Fatal]
func (ut Utils) Fatal(args ...interface{}) {
	ut.Helper()
	ut.Log(args...)
	ut.FailNow()
}

// Fatalf is the same as [testing.common.Fatalf]
func (ut Utils) Fatalf(format string, args ...interface{}) {
	ut.Helper()
	ut.Logf(format, args...)
	ut.FailNow()
}

// Log is the same as [testing.common.Log]
func (ut Utils) Log(args ...interface{}) {
	ut.Helper()
	ut.Logf("%s", fmt.Sprintln(args...))
}

// Error is the same as [testing.common.Error]
func (ut Utils) Error(args ...interface{}) {
	ut.Helper()
	ut.Log(args...)
	ut.Fail()
}

// Errorf is the same as [testing.common.Errorf]
func (ut Utils) Errorf(format string, args ...interface{}) {
	ut.Helper()
	ut.Logf(format, args...)
	ut.Fail()
}

// Skipf is the same as [testing.common.Skipf]
func (ut Utils) Skipf(format string, args ...interface{}) {
	ut.Helper()
	ut.Logf(format, args...)
	ut.SkipNow()
}

// Skip is the same as [testing.common.Skip]
func (ut Utils) Skip(args ...interface{}) {
	ut.Helper()
	ut.Log(args...)
	ut.SkipNow()
}

// Run f as a sub-test
func (ut Utils) Run(name string, f func(t G)) bool {
	runVal := reflect.ValueOf(ut.Testable).MethodByName("Run")
	return runVal.Call([]reflect.Value{
		reflect.ValueOf(name),
		reflect.MakeFunc(runVal.Type().In(1), func(args []reflect.Value) []reflect.Value {
			f(New(args[0].Interface().(Testable)))
			return nil
		}),
	})[0].Interface().(bool)
}

// Parallel is the same as [testing.T.Parallel]
func (ut Utils) Parallel() Utils {
	reflect.ValueOf(ut.Testable).MethodByName("Parallel").Call(nil)
	return ut
}

// DoAfter d duration if the test is still running
func (ut Utils) DoAfter(d time.Duration, do func()) (cancel func()) {
	ctx := ut.Context()
	go func() {
		ut.Helper()
		tmr := time.NewTimer(d)
		defer tmr.Stop()
		select {
		case <-ctx.Done():
		case <-tmr.C:
			do()
		}
	}()
	return ctx.Cancel
}

// PanicAfter d duration if the test is still running
func (ut Utils) PanicAfter(d time.Duration) (cancel func()) {
	return ut.DoAfter(d, func() {
		ut.Helper()
		panicWithTrace(fmt.Sprintf("%s timeout after %v", ut.Name(), d))
	})
}

// Context that will be canceled after the test
func (ut Utils) Context() Context {
	ctx, cancel := context.WithCancel(context.Background())
	ut.Cleanup(cancel)
	return Context{ctx, cancel}
}

// Timeout context that will be canceled after the test
func (ut Utils) Timeout(d time.Duration) Context {
	ctx, cancel := context.WithTimeout(context.Background(), d)
	ut.Cleanup(cancel)
	return Context{ctx, cancel}
}

// RandStr generates a random string with the specified length
func (ut Utils) RandStr(l int) string {
	ut.Helper()
	b := ut.RandBytes((l + 1) / 2)
	return hex.EncodeToString(b)[:l]
}

// RandInt generates a random integer within [min, max)
func (ut Utils) RandInt(min, max int) int {
	ut.Helper()
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	ut.err(err)
	return int(n.Int64()) + min
}

// RandBytes generates a random byte array with the specified length
func (ut Utils) RandBytes(l int) []byte {
	ut.Helper()
	b := make([]byte, l)
	_, err := rand.Read(b)
	ut.err(err)
	return b
}

// Render template. It will use [Utils.Read] to read the value as the template string.
func (ut Utils) Render(value interface{}, data interface{}) *bytes.Buffer {
	ut.Helper()
	out := bytes.NewBuffer(nil)
	t := template.New("")
	t, err := t.Parse(ut.Read(value).String())
	ut.err(err)
	ut.err(t.Execute(out, data))
	return out
}

// WriteFile at path with content, it uses [Utils.Open] to open the file.
func (ut Utils) WriteFile(path string, content interface{}) {
	f := ut.Open(true, path)
	defer func() { ut.err(f.Close()) }()
	ut.Write(content)(f)
}

// PathExists checks if path exists
func (ut Utils) PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Chdir is like [os.Chdir] but will restore the dir after test.
func (ut Utils) Chdir(dir string) {
	ut.Helper()
	cwd, err := os.Getwd()
	ut.err(err)
	ut.err(os.Chdir(dir))
	ut.Cleanup(func() { _ = os.Chdir(cwd) })
}

// Setenv is like [os.Setenv] but will restore the env after test.
func (ut Utils) Setenv(key, value string) {
	ut.Helper()
	old := os.Getenv(key)
	ut.err(os.Setenv(key, value))
	ut.Cleanup(func() { _ = os.Setenv(key, old) })
}

// MkdirAll is like [os.MkdirAll] but will remove the dir after test and fail the test if error.
// The default perm is 0755.
func (ut Utils) MkdirAll(perm fs.FileMode, path string) {
	if perm == 0 {
		perm = 0755
	}

	dir := filepath.Dir(path)

	if !ut.PathExists(dir) {
		ut.MkdirAll(perm, dir)
	}

	if ut.PathExists(path) {
		return
	}

	ut.err(os.Mkdir(path, perm))
	ut.Cleanup(func() { _ = os.RemoveAll(path) })
}

// Open a file. Override it if create is true. Directories will be auto-created.
// If the directory and file doesn't exist, it will be removed after the test.
func (ut Utils) Open(create bool, path string) (f *os.File) {
	ut.Helper()

	var err error
	if create {
		ut.MkdirAll(0, filepath.Dir(path))
		f, err = os.Create(path)
		if err == nil {
			ut.Cleanup(func() { _ = os.Remove(path) })
		}
	} else {
		f, err = os.Open(path)
	}
	ut.err(err)

	return f
}

// Read all from value. If the value is string and it's a file path,
// the file content will be read, or the string will be returned.
// If the value is [io.Reader], the reader will be read. If the value is []byte, the value will be returned.
// Others will be converted to string and returned.
func (ut Utils) Read(value interface{}) *bytes.Buffer {
	ut.Helper()

	var r io.Reader

	switch v := value.(type) {
	case string:
		if !ut.PathExists(v) {
			return bytes.NewBufferString(v)
		}
		f := ut.Open(false, v)
		defer func() { ut.err(f.Close()) }()
		r = f
	case io.Reader:
		r = v
	case []byte:
		return bytes.NewBuffer(v)
	default:
		return bytes.NewBufferString(fmt.Sprintf("%v", v))
	}

	b := bytes.NewBuffer(nil)
	_, err := io.Copy(b, r)
	ut.err(err)
	return b
}

// JSON from string, []byte, or io.Reader
func (ut Utils) JSON(src interface{}) (v interface{}) {
	ut.Helper()

	var b []byte
	switch obj := src.(type) {
	case []byte:
		b = obj
	case string:
		b = []byte(obj)
	case io.Reader:
		var err error
		b, err = ioutil.ReadAll(obj)
		ut.err(err)
	}
	ut.err(json.Unmarshal(b, &v))
	return
}

// ToJSON convert obj to JSON bytes
func (ut Utils) ToJSON(obj interface{}) *bytes.Buffer {
	ut.Helper()
	b, err := json.MarshalIndent(obj, "", "  ")
	ut.err(err)
	return bytes.NewBuffer(b)
}

// ToJSONString convert obj to JSON string
func (ut Utils) ToJSONString(obj interface{}) string {
	ut.Helper()
	return ut.ToJSON(obj).String()
}

// Write obj to the writer. Encode obj to []byte and cache it for writer.
// If obj is not []byte, string, or [io.Reader], it will be encoded as JSON.
func (ut Utils) Write(obj interface{}) (writer func(io.Writer)) {
	lock := sync.Mutex{}
	var cache []byte
	return func(w io.Writer) {
		lock.Lock()
		defer lock.Unlock()

		ut.Helper()

		if cache != nil {
			_, err := w.Write(cache)
			ut.err(err)
			return
		}

		buf := bytes.NewBuffer(nil)
		w = io.MultiWriter(buf, w)

		var err error
		switch v := obj.(type) {
		case []byte:
			_, err = w.Write(v)
		case string:
			_, err = w.Write([]byte(v))
		case io.Reader:
			_, err = io.Copy(w, v)
		default:
			err = json.NewEncoder(w).Encode(v)
		}
		ut.err(err)
		cache = buf.Bytes()
	}
}

// HandleHTTP handles a request. If file exists serve the file content. The file will be used to set the Content-Type header.
// If the file doesn't exist, the value will be encoded by [Utils.Write] and used as the response body.
func (ut Utils) HandleHTTP(file string, value ...interface{}) func(http.ResponseWriter, *http.Request) {
	var obj interface{}
	if len(value) > 1 {
		obj = value
	} else if len(value) == 1 {
		obj = value[0]
	}

	write := ut.Write(obj)

	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(file); err == nil {
			http.ServeFile(w, r, file)
			return
		}

		if obj == nil {
			return
		}

		w.Header().Add("Content-Type", mime.TypeByExtension(filepath.Ext(file)))

		write(w)
	}
}

// Serve http on a random port. The server will be auto-closed after the test.
func (ut Utils) Serve() *Router {
	ut.Helper()
	return ut.ServeWith("tcp4", "127.0.0.1:0")
}

// ServeWith specified network and address
func (ut Utils) ServeWith(network, address string) *Router {
	ut.Helper()

	mux := http.NewServeMux()
	srv := &http.Server{Handler: mux}

	l, err := net.Listen(network, address)
	ut.err(err)

	ut.Cleanup(func() { _ = srv.Close() })

	go func() { _ = srv.Serve(l) }()

	u, err := url.Parse("http://" + l.Addr().String())
	ut.err(err)

	return &Router{ut, u, srv, mux}
}

// Router of a http server
type Router struct {
	ut      Utils
	HostURL *url.URL
	Server  *http.Server
	Mux     *http.ServeMux
}

// URL will prefix the path with the server's host
func (rt *Router) URL(path ...string) string {
	return rt.HostURL.String() + strings.Join(path, "")
}

// Route on the pattern. Check the doc of [http.ServeMux] for the syntax of pattern.
// It will use [Utils.HandleHTTP] to handle each request.
func (rt *Router) Route(pattern, file string, value ...interface{}) *Router {
	h := rt.ut.HandleHTTP(file, value...)

	rt.Mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		h(w, r)
	})

	return rt
}

// ReqMIME option type, it should be like ".json", "test.json", "a/b/c.jpg", etc
type ReqMIME string

// Req the url. The method is the http method, default value is "GET".
// If an option is http.Header, it will be used as the request header.
// If an option is [Utils.ReqMIME], it will be used to set the Content-Type header.
// Other option type will be treat as request body, it will be encoded by Utils.Write .
func (ut Utils) Req(method, url string, options ...interface{}) *ResHelper {
	ut.Helper()

	header := http.Header{}
	var contentType string
	var body io.Reader

	for _, item := range options {
		switch val := item.(type) {
		case http.Header:
			header = val
		case ReqMIME:
			contentType = mime.TypeByExtension(filepath.Ext(string(val)))
		default:
			buf := bytes.NewBuffer(nil)
			ut.Write(val)(buf)
			body = buf
		}
	}

	req, err := http.NewRequest(method, url, body)
	ut.err(err)

	if header != nil {
		req.Header = header
	}

	req.Header.Set("Content-Type", contentType)

	res, err := http.DefaultClient.Do(req)
	ut.err(err)

	return &ResHelper{ut, res}
}

// ResHelper of the request
type ResHelper struct {
	ut Utils
	*http.Response
}

// Bytes body
func (res *ResHelper) Bytes() *bytes.Buffer {
	res.ut.Helper()
	return res.ut.Read(res.Body)
}

// String body
func (res *ResHelper) String() string {
	res.ut.Helper()
	return res.Bytes().String()
}

// JSON body
func (res *ResHelper) JSON() (v interface{}) {
	res.ut.Helper()
	return res.ut.JSON(res.Body)
}

func (ut Utils) err(err error) {
	ut.Helper()

	if err != nil {
		ut.Fatal(err)
	}
}

// there no way to stop a blocking test from outside
var panicWithTrace = func(v interface{}) {
	panic(v)
}
