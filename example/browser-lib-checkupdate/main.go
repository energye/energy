// auto update application

package main

import (
	"archive/zip"
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef/i18n"
	"github.com/energye/energy/v2/cmd/autoupdate"
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/example/browser-lib-checkupdate/form"
	"github.com/energye/golcl/energy/inits"
	"github.com/energye/golcl/energy/tools"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api/dllimports"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/colors"
	"github.com/energye/golcl/pkgs/libname"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

/*
  检测liblcl库的更新
  下面的代码演示是未通过cef库导入的liblcl,而是单独导入,这样可以减少执行文件的大小
    []*dllimports.ImportTable
  通过imports.SetEnergyImportDefs(version)将proc导入到执行文件中, 为保持和cef的proc下标一致,第0个下标为空导入
*/

var (
	//go:embed resources
	resources embed.FS
	// form
	updateForm *form.UpdateForm
	version    = []*dllimports.ImportTable{
		dllimports.NewEnergyImport("", 0),                //空导入
		dllimports.NewEnergyImport("LibVersion", 0),      //获取lib库的版本号
		dllimports.NewEnergyImport("LibBuildVersion", 0), //获取lib库的构建工具版本
	}
)

func main() {
	// 针对 liblcl
	energyPath := filepath.Join(os.TempDir(), "energy")
	if !tools.IsExist(energyPath) {
		os.Mkdir(energyPath, fs.ModePerm)
	}
	var (
		libPath    = libname.LibPath()
		dstLibPath = filepath.Join(energyPath, libname.GetDLLName())
	)
	if tools.IsExist(dstLibPath) {
		os.Remove(dstLibPath)
	}

	// 在初始化之前 先把 liblcl 复制一份到临时目录中
	// 然后我们加载临时目录中的 liblcl
	// 之后升级时使用升级新文件替换实际的 liblcl
	// 可以在程度内直接替换
	src, err := os.Open(libPath)
	if err != nil {
		panic(err)
	}
	dst, err := os.OpenFile(dstLibPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	// 复制一份
	_, err = io.Copy(dst, src)
	if err != nil {
		panic(err)
	}
	dst.Close()
	src.Close()
	libname.LibName = dstLibPath // 设置临时目录的 liblcl
	defer os.Remove(dstLibPath)

	// proc 注入到 imports
	imports.SetEnergyImportDefs(version)
	// 初始化golcl
	inits.Init(nil, &resources)
	i18n.SetLocalFS(&resources, "resources")
	//i18n.Switch(consts.LANGUAGE_en_US)
	i18n.Switch(consts.LANGUAGE_zh_CN)
	// 开启自动更新检查
	autoupdate.IsCheckUpdate(true)
	// 如果 liblcl 有更新该函数将被回调,并创建更新窗口
	autoupdate.CanUpdateLiblcl = func(model *autoupdate.Model, level int, canUpdate bool) {
		updateVersion := model.Versions[model.Latest]
		var energyLiblcl = func() (string, bool) {
			if common.IsWindows() {
				return fmt.Sprintf("Windows %d bits", strconv.IntSize), true
			} else if common.IsLinux() {
				return "Linux x86 64 bits", true
			} else if common.IsDarwin() {
				return "MacOSX x86 64 bits", true
			}
			//not support download
			return fmt.Sprintf("%v %v", runtime.GOOS, runtime.GOARCH), false
		}

		// 这里使用窗口形式展示更新
		// 运行应用后窗口创建时回调
		// 通过代码设计窗口上的UI组件
		form.OnCreate = func(m *form.UpdateForm) {
			// 应用图标
			lcl.Application.Icon().LoadFromFSFile("resources/icon.ico")

			// 窗口一些属性配置
			m.SetDoubleBuffered(true)
			m.EnabledMinimize(false)
			m.EnabledMaximize(false)
			m.SetShowHint(true)
			//m.SetFormStyle(types.FsSystemStayOnTop)
			m.SetPosition(types.PoDesktopCenter)
			m.SetBorderStyle(types.BsSingle)
			//m.SetBorderStyle(types.BsNone)
			//m.SetShowInTaskBar(types.StNever)
			m.SetColor(colors.ClWhite)
			m.SetWidth(590)
			m.SetHeight(390)
			m.SetCaption(i18n.Resource("title")) // 自定义窗口标题
			var titleHeight int32 = 0
			if runtime.GOOS == "windows" {
				m.SetBorderStyle(types.BsNone)
				titleHeight = 32
				// 自定义窗口标题栏
				m.TitlePanel = m.NewPanel()
				m.TitlePanel.SetColor(colors.ClTeal)
				m.TitlePanel.SetHeight(titleHeight) // icon 的高
				m.TitlePanel.SetWidth(m.Width())
				// 模拟标题栏移动窗口
				var (
					isDown bool
					dx, dy int32
				)
				m.TitlePanel.SetOnMouseDown(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
					isDown = true
					dx, dy = x, y
				})
				m.TitlePanel.SetOnMouseUp(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
					isDown = false
				})
				m.TitlePanel.SetOnMouseMove(func(sender lcl.IObject, shift types.TShiftState, x, y int32) {
					if isDown { //鼠标按下时计算移动坐标
						m.SetLeft(m.Left() - (dx - x))
						m.SetTop(m.Top() - (dy - y))
					}
				})

				// title -> icon
				titleIcon := lcl.NewImage(m.TitlePanel)
				titleIcon.SetParent(m.TitlePanel)
				titleIcon.SetWidth(32)  // icon 的宽
				titleIcon.SetHeight(32) // icon 的高
				titleIcon.Picture().LoadFromFSFile("resources/icon.png")

				// title -> text
				titleText := lcl.NewLabel(m.TitlePanel)
				titleText.SetParent(m.TitlePanel)
				titleText.SetTop(3)
				titleText.SetLeft(40)
				titleText.Font().SetSize(12)
				titleText.Font().SetColor(colors.ClWhite)
				titleText.SetCaption(m.Caption())

				// title -> close button
				titleClose := lcl.NewImageButton(m.TitlePanel)
				titleClose.SetParent(m.TitlePanel)
				titleClose.SetImageCount(4)
				titleClose.SetAutoSize(true)
				titleClose.SetCursor(types.CrHandPoint)
				titleClose.Picture().LoadFromFSFile("resources/btn_close.png")
				titleClose.SetLeft(m.Width() - 40) //关闭按钮位置 left = 窗口宽 - 按钮图片宽
				titleClose.SetHint(i18n.Resource("close"))
				titleClose.SetOnClick(func(lcl.IObject) {
					m.Close() // 关闭窗口
				})
			}

			// background
			bgImage := lcl.NewImage(m)
			bgImage.SetParent(m)
			bgImage.SetTop(titleHeight + 10)
			bgImage.SetWidth(271)                              // 图片宽
			bgImage.SetHeight(60)                              // 图片高
			bgImage.SetLeft((m.Width() - bgImage.Width()) / 2) // 设置以窗口居中
			bgImage.Picture().LoadFromFSFile("resources/bg.png")

			var updatePanelHeight int32 = 200
			// 更新提醒 panel
			m.UpdatePromptPanel = m.NewPanel()
			m.UpdatePromptPanel.SetTop(bgImage.Top() + bgImage.Height() + 20)
			m.UpdatePromptPanel.SetWidth(m.Width())
			m.UpdatePromptPanel.SetHeight(updatePanelHeight)

			// 更新内容
			m.UpdateContentMemo = lcl.NewMemo(m.UpdatePromptPanel)
			m.UpdateContentMemo.SetParent(m.UpdatePromptPanel)
			ucw := m.Width() / 4
			m.UpdateContentMemo.SetWidth(ucw * 3)
			m.UpdateContentMemo.SetLeft((m.Width() - m.UpdateContentMemo.Width()) / 2)
			m.UpdateContentMemo.SetHeight(updatePanelHeight)
			m.UpdateContentMemo.SetReadOnly(true)
			m.UpdateContentMemo.SetColor(colors.ClWhite)
			if !canUpdate {
				m.UpdateContentMemo.Lines().Add(i18n.Resource("currentVersion") + ": " + model.CurrentVersion)
				m.UpdateContentMemo.Lines().Add(i18n.Resource("latestVersion") + ": " + model.Latest)
				m.UpdateContentMemo.SetEnabled(false)
			} else if canUpdate { // 有更新
				m.UpdateContentMemo.SetScrollBars(types.SsAutoBoth)
				m.UpdateContentMemo.Lines().Add(i18n.Resource("updateContent") + " " + model.Latest)
				for i, content := range updateVersion.Content {
					m.UpdateContentMemo.Lines().Add("  " + strconv.Itoa(i+1) + ". " + content)
				}
				// liblcl 下载版本URL
				liblclZipName, _ := energyLiblcl()
				downUrl := strings.Replace(model.Download.Url, "{url}", model.Download.Source[model.Download.SourceSelect], -1) // 使用配置的下载源
				downUrl = strings.Replace(downUrl, "{version}", updateVersion.EnergyVersion, -1)                                // liblcl 所属的 enregy 版本
				downUrl = strings.Replace(downUrl, "{OSARCH}", liblclZipName, -1)                                               // 根据系统架构获取对应的文件名
				m.UpdateContentMemo.Lines().Add("")
				m.UpdateContentMemo.Lines().Add(i18n.Resource("downloadURL"))
				m.UpdateContentMemo.Lines().Add(downUrl)

				// 取消按钮
				//cancelBtn := lcl.NewImageButton(m)
				//cancelBtn.SetParent(m)
				//cancelBtn.SetImageCount(3)
				//cancelBtn.SetAutoSize(true)
				//cancelBtn.SetCursor(types.CrHandPoint)
				//cancelBtn.Picture().LoadFromFSFile("resources/btn-cancel.png")
				//cancelBtn.SetLeft(300)
				//cancelBtn.SetTop(290)
				//cancelBtn.SetHint(i18n.Resource("cancel"))
				//cancelBtn.SetOnClick(func(lcl.IObject) {
				//	//m.Close()
				//})

				// 下载
				var isNotDownload = true
				// 更新下载进度条
				updateProgressBar := lcl.NewProgressBar(m)
				updateProgressBar.SetParent(m)
				updateProgressBar.SetTop(updatePanelHeight + m.UpdatePromptPanel.Top() + 5)
				updateProgressBar.SetLeft(m.UpdateContentMemo.Left())
				updateProgressBar.SetMin(0)
				updateProgressBar.SetMax(100)
				updateProgressBar.SetWidth(m.UpdateContentMemo.Width())
				updateProgressBar.SetHeight(10)
				updateProgressBar.SetVisible(false)
				// 更新按钮
				updateBtn := lcl.NewImageButton(m)
				updateBtn.SetParent(m)
				updateBtn.SetImageCount(3)
				//updateBtn.SetAutoSize(true)
				updateBtn.SetWidth(80)
				updateBtn.SetHeight(50)
				updateBtn.SetCursor(types.CrHandPoint)
				updateBtn.Picture().LoadFromFSFile("resources/btn-update.png")
				updateBtn.SetLeft(400)
				updateBtn.SetTop(325)
				updateBtn.SetHint(i18n.Resource("update"))
				updateBtn.SetOnClick(func(lcl.IObject) {
					if isNotDownload {
						updateProgressBar.SetPosition(0)
						updateProgressBar.SetVisible(isNotDownload)
						isNotDownload = false
						updateBtn.SetEnabled(isNotDownload)

						var savePath, _ = filepath.Split(libPath)
						var fileName, _ = energyLiblcl()
						var saveFilePath = filepath.Join(savePath, fileName) + ".zip" // zip
						m.UpdateContentMemo.Lines().Add("---------------------------------------------------------------")
						m.UpdateContentMemo.Lines().Add(i18n.Resource("beginDownload"))
						m.UpdateContentMemo.Lines().Add("  " + i18n.Resource("downloadURL") + ": " + downUrl)
						m.UpdateContentMemo.Lines().Add("  " + i18n.Resource("savePath") + ": " + saveFilePath)
						go func() {
							// 下载地址, 保存目录, 回调更新进程函数
							downloadFile(downUrl, saveFilePath, func(totalLength, processLength int64, err error) {
								var process = int32((float64(processLength) / float64(totalLength)) * 100)
								lcl.ThreadSync(func() {
									if err != nil {
										m.UpdateContentMemo.Lines().Add(i18n.Resource("downloadError") + ":\n  " + err.Error())
										m.UpdateContentMemo.Lines().Add(i18n.Resource("downloadAgain"))
									} else {
										updateProgressBar.SetPosition(process)
										if process >= 100 {
											m.UpdateContentMemo.Lines().Add(i18n.Resource("endDownload") + " -> 100")
										}
									}
								})
							})
							m.UpdateContentMemo.Lines().Add("---------------------------------------------------------------")
							updateProgressBar.SetPosition(0)
							m.UpdateContentMemo.Lines().Add(i18n.Resource("extractUnZip") + ": " + filepath.Join(savePath, libname.GetDLLName()))
							// 解压释放文件
							extractUnZip(saveFilePath, savePath, func(totalLength, processLength int64, err error) {
								var process = int32((float64(processLength) / float64(totalLength)) * 100)
								lcl.ThreadSync(func() {
									if err != nil {
										m.UpdateContentMemo.Lines().Add(i18n.Resource("updateError") + ":\n  " + err.Error())
										m.UpdateContentMemo.Lines().Add(i18n.Resource("updateAgain"))
									} else {
										updateProgressBar.SetPosition(process)
										if process >= 100 {
											m.UpdateContentMemo.Lines().Add(i18n.Resource("extractUnZip") + ": success -> 100")
											m.UpdateContentMemo.Lines().Add(i18n.Resource("updateSuccess"))
										}
									}
								})
							}, libname.GetDLLName())
							isNotDownload = true
							updateBtn.SetEnabled(isNotDownload)
						}()
					}
				})
			}
		}
		// run and create update form
		lcl.RunApp(&updateForm)
	}
	// 检查 liblcl 库是否有更新
	autoupdate.CheckUpdate()
}

// 下载文件
func downloadFile(url string, localPath string, callback func(totalLength, processLength int64, err error)) error {
	var (
		fsize   int64
		buf     = make([]byte, 1024*10)
		written int64
	)
	tmpFilePath := localPath + ".download"
	client := new(http.Client)
	resp, err := client.Get(url)
	if err != nil {
		callback(0, 0, err)
		return err
	}
	fsize, err = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 32)
	if err != nil {
		callback(0, 0, err)
		return err
	}
	println("Save path: [", localPath, "] file size:", fsize)
	file, err := os.Create(tmpFilePath)
	if err != nil {
		callback(0, 0, err)
		return err
	}
	defer file.Close()
	if resp.Body == nil {
		callback(0, 0, err)
		return nil
	}
	defer resp.Body.Close()
	for {
		nr, er := resp.Body.Read(buf)
		if nr > 0 {
			nw, err := file.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			callback(fsize, written, err)
			if err != nil {
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
		file.Sync()
		file.Close()
		err = os.Rename(tmpFilePath, localPath)
		if err != nil {
			return err
		}
	}
	return err
}

func writeFile(r io.Reader, w *os.File, totalLength int64, callback func(totalLength, processLength int64, err error)) {
	buf := make([]byte, 1024*10)
	var written int64
	for {
		nr, err := r.Read(buf)
		if nr > 0 {
			nw, err := w.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			callback(totalLength, written, err)
			if err != nil {
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if err != nil {
			break
		}
	}
}

func extractUnZip(filePath, targetPath string, callback func(totalLength, processLength int64, err error), files ...interface{}) {
	if rc, err := zip.OpenReader(filePath); err == nil {
		defer rc.Close()
		for i := 0; i < len(files); i++ {
			if f, err := rc.Open(files[i].(string)); err == nil {
				defer f.Close()
				st, _ := f.Stat()
				targetFileName := filepath.Join(targetPath, st.Name())
				if st.IsDir() {
					os.MkdirAll(targetFileName, st.Mode())
					continue
				}
				if targetFile, err := os.Create(targetFileName); err == nil {
					writeFile(f, targetFile, st.Size(), callback)
					targetFile.Close()
				} else {
					callback(0, 0, err)
				}
			} else {
				callback(0, 0, err)
				return
			}
		}
	} else {
		if err != nil {
			callback(0, 0, err)
		}
	}
}
