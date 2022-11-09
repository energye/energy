package cmd

import (
	"compress/bzip2"
	"encoding/json"
	"errors"
	"fmt"
	progressbar "github.com/energye/energy/cmd/progress-bar"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

var CmdInstall = &Command{
	UsageLine: "install [path] [version]",
	Short:     "Automatically configure the CEF and Energy framework",
	Long: `Automatically configure the CEF and Energy framework
During this process, CEF and Energy are downloaded
Default framework name is "CEFEnergy"`,
}

const download_version_config_url = "https://energy.yanghy.cn/energy_download.json"

type downloadInfo struct {
	fileName string
	savePath string
	url      string
	success  bool
}

func init() {
	CmdInstall.Run = runInstall
}

//运行安装
func runInstall(c *CommandConfig) error {
	if c.Install.Path == "" {
		c.Install.Path = c.Wd
	}
	c.Install.Path = filepath.Join(c.Install.Path, c.Install.Name)
	if c.Install.Version == "" {
		c.Install.Version = "latest"
	}
	os.MkdirAll(c.Install.Path, fs.ModeDir)
	println("Start downloading CEF and Energy dependency")
	downloadJSON, err := downloadConfig()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
	var downloadVersion map[string]interface{}
	if err := json.Unmarshal(downloadJSON, &downloadVersion); err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
	version, ok := downloadVersion[c.Install.Version]
	if !ok {
		println("Invalid version number:", c.Install.Version)
		os.Exit(1)
	}
	osConfig := version.(map[string]interface{})[runtime.GOOS]
	var osVersion map[string]interface{}
	if runtime.GOOS == "windows" {
		//区分windows系统位数
		bits := osConfig.(map[string]interface{})
		osVersion = bits[fmt.Sprintf("%d", strconv.IntSize)].(map[string]interface{})
	}
	//下载地址
	var (
		cef, cefOk       = osVersion["cef"].(string)
		energy, energyOk = osVersion["energy"].(string)
	)
	if cefOk && energyOk {
		var downloads = make(map[string]*downloadInfo)
		downloads["cef"] = &downloadInfo{fileName: urlName(cef), savePath: filepath.Join(c.Install.Path, urlName(cef)), url: cef}
		downloads["energy"] = &downloadInfo{fileName: urlName(energy), savePath: filepath.Join(c.Install.Path, urlName(energy)), url: energy}
		for key, dl := range downloads {
			fmt.Printf("start download %s url: %s\n", key, dl.url)
			bar := progressbar.NewBar(100)
			bar.SetNotice("downloading " + dl.fileName + ": ")
			bar.SetGraph("█")
			bar.HideRatio()
			err = downloadFile(dl.url, dl.savePath, func(totalLength, processLength int64) {
				bar.PrintBar(int((float64(processLength) / float64(totalLength)) * 100))
			})
			bar.PrintEnd("download " + dl.fileName + " end")
			if err != nil {
				println("download", dl.fileName, "error", err)
			}
			dl.success = err == nil
		}
		println("Release files")
		for key, dl := range downloads {
			if dl.success {
				if key == "cef" {
					bar := progressbar.NewBar(0)
					bar.SetNotice("Unpack file " + dl.fileName + ": ")
					tarName := UnBz2ToTar(dl.savePath, func(processLength int64) {
						bar.PrintSizeBar(processLength)
					})
					bar.PrintEnd()
					println("Unpack file", dl.fileName, "end.")
					ExtractFiles(key, tarName)
				} else if key == "energy" {
					ExtractFiles(key, dl.savePath)
				}
			}
		}
	} else {
		println("Invalid version number:", c.Install.Version)
		os.Exit(1)
	}
	return nil
}

//提取文件
func ExtractFiles(keyName, path string) {
	println("extract", keyName, "file path:", path)
}

//释放bz2文件到tar
func UnBz2ToTar(name string, callback func(processLength int64)) string {
	fileBz2, err := os.Open(name)
	if err != nil {
		fmt.Errorf("%s", err.Error())
		os.Exit(1)
	}
	defer fileBz2.Close()
	dirName := fileBz2.Name()
	dirName = dirName[:strings.LastIndex(dirName, ".")]
	r := bzip2.NewReader(fileBz2)
	w, err := os.Create(dirName)
	if err != nil {
		fmt.Errorf("%s", err.Error())
		os.Exit(1)
	}
	defer w.Close()
	buf := make([]byte, 1024*10)
	var written int64
	for {
		nr, err := r.Read(buf)
		if nr == 0 || err != nil {
			break
		}
		nw, err := w.Write(buf[:nr])
		if nw > 0 {
			written += int64(nw)
		}
		callback(written)
		if err != nil {
			break
		}
		if nr != nw {
			err = io.ErrShortWrite
			break
		}
	}
	return dirName
}

//url文件名
func urlName(downloadUrl string) string {
	if u, err := url.QueryUnescape(downloadUrl); err != nil {
		return ""
	} else {
		u = u[strings.LastIndex(u, "/")+1:]
		return u
	}
}

//下载文件配置
func downloadConfig() ([]byte, error) {
	client := new(http.Client)
	resp, err := client.Get(download_version_config_url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func isFileExist(filename string, filesize int64) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if filesize == info.Size() {
		return true
	}
	del := os.Remove(filename)
	if del != nil {
		println(del)
	}
	return false
}

//下载文件
func downloadFile(url string, localPath string, callback func(totalLength, processLength int64)) error {
	var (
		fsize   int64
		buf     = make([]byte, 32*1024)
		written int64
	)
	tmpFilePath := localPath + ".download"
	client := new(http.Client)
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	fsize, err = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 32)
	if err != nil {
		return err
	}
	if isFileExist(localPath, fsize) {
		return err
	}
	println("save path: [", localPath, "] file size:", fsize)
	file, err := os.Create(tmpFilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	if resp.Body == nil {
		return errors.New("body is null")
	}
	defer resp.Body.Close()
	for {
		nr, er := resp.Body.Read(buf)
		if nr > 0 {
			nw, ew := file.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			callback(fsize, written)
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	if err == nil {
		file.Close()
		err = os.Rename(tmpFilePath, localPath)
		if err != nil {
			return err
		}
	}
	return err
}
