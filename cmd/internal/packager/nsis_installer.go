package packager

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"github.com/energye/golcl/energy/tools"
	"github.com/energye/golcl/tools/command"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"text/template"
)

//go:embed assets
var assets embed.FS

const (
	windowsNsis      = "windows/installer-nsis.nsi"
	windowsNsisTools = "windows/installer-tools.nsh"
)

func GeneraNSISInstaller(projectData *Project) error {
	switch runtime.GOOS {
	case "windows":
		if err := windows(projectData); err != nil {
			return err
		}
	case "linux":
	case "darwin":
	default:
		return errors.New("unsupported system")
	}
	makeNSIS(projectData)
	return nil
}

func windows(projectData *Project) error {
	// 创建构建输出目录
	buildOutDir := buildOutPath(projectData)
	if !tools.IsExist(buildOutDir) {
		if err := os.MkdirAll(buildOutDir, 0755); err != nil {
			return fmt.Errorf("unable to create directory: %w", err)
		}
	}
	// 生成安装生成配置文件
	if nsisData, err := readFile(projectData, windowsNsis); err != nil {
		return err
	} else {
		if err = writeFile(projectData, windowsNsis, nsisData); err != nil {
			return err
		}
	}
	if toolsData, err := readFile(projectData, windowsNsisTools); err != nil {
		return err
	} else {
		tmpl, err := template.New("tools").Parse(string(toolsData))
		if err != nil {
			return err
		}
		data := make(map[string]any)
		data["Name"] = projectData.Name
		data["Info"] = projectData.Info
		var out bytes.Buffer
		if err = tmpl.Execute(&out, data); err != nil {
			return err
		}
		if err = writeFile(projectData, windowsNsisTools, out.Bytes()); err != nil {
			return err
		}
	}
	return nil
}

// 使用nsis生成安装包
func makeNSIS(projectData *Project) error {
	var args []string
	cmd := command.NewCMD()
	cmd.Dir = projectData.ProjectPath
	cmd.MessageCallback = func(bytes []byte, err error) {
		println("makensis:", string(bytes))
	}
	nsisScriptPath := filepath.Join(buildOutPath(projectData), windowsNsis)
	args = append(args, "-DARG_ENERGY_AMD64_BINARY=E:\\SWT\\gopath\\src\\github.com\\energye\\energy\\cmd\\internal\\test\\simple.exe")
	if projectData.Info.License != "" {
		// 授权信息文本目录: ..\LICENSE.txt
		args = append(args, "-DARG_ENERGY_PAGE_LICENSE="+projectData.Info.License)
	}
	if projectData.Info.Language != "" {
		// default English
		// 可选多种语言: SimpChinese, 参考目录: NSIS\Contrib\Language files
		args = append(args, "-DARG_ENERGY_LANGUAGE="+projectData.Info.Language)
	}
	args = append(args, nsisScriptPath)
	cmd.Command("makensis", args...)

	return nil
}

// CommandExists 命令是否存在
func CommandExists(name string) bool {
	_, err := exec.LookPath(name)
	if err != nil {
		return false
	}
	return true
}

// 返回根据配置的资源目录
func assetsPath(projectData *Project, file string) string {
	return filepath.ToSlash(filepath.Join(projectData.BuildAssetsDir, file))
}

// 返回固定的构建输出目录 $current/build
func buildOutPath(projectData *Project) string {
	return filepath.Join(projectData.ProjectPath, "build")
}

// ReadFile
//  读取文件，根据项目配置先在本地目录读取，如果读取失败，则在内置资源目录读取
func readFile(projectData *Project, file string) ([]byte, error) {
	localFilePath := assetsPath(projectData, file)
	content, err := os.ReadFile(localFilePath)
	if errors.Is(err, fs.ErrNotExist) {
		content, err = fs.ReadFile(assets, localFilePath)
		if err != nil {
			return nil, err
		}
		if err := writeFile(projectData, file, content); err != nil {
			return nil, fmt.Errorf("unable to create file in build folder: %s", err)
		}
		return content, nil
	}

	return content, err
}

// 写文件
func writeFile(projectData *Project, file string, content []byte) error {
	buildOutDir := buildOutPath(projectData)
	if !tools.IsExist(buildOutDir) {
		if err := os.MkdirAll(buildOutDir, 0755); err != nil {
			return fmt.Errorf("unable to create directory: %w", err)
		}
	}
	targetPath := filepath.Join(buildOutDir, file)
	if !tools.IsExist(filepath.Dir(targetPath)) {
		if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
			return fmt.Errorf("unable to create directory: %w", err)
		}
	}
	if err := os.WriteFile(targetPath, content, 0644); err != nil {
		return err
	}
	return nil
}
