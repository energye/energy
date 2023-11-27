### liblcl.xx 动态库字节码存放方式

### 介绍
> 1. 该方式可不必下载`liblcl.xx`动态库, 使用`energy`时,只需安装CEF框架
> 2. energy将`liblcl.xx`动态库写入到`liblclbinres/xxx.go`文件中以字节码形式存放
> 3. 使用该模式在`liblclbinres/xxx.go`读取`liblcl.xx`动态库字节码并释放到本机目录

### 使用
> 1. 编译时增加编译命令参数 `-tags="tempdll [cef version]"`
>> `go build -tags="tempdll latest"`
> 2. 开发工具中以Goland为例
>> 命令行参数 `Go tool arguments` 中配置 `-tags="tempdll latest"`
> 3. `TempDLL` 全局变量配置`liblcl.xx`动态库保存目录
> 4. `TempDLL` 全局变量根据编译参数`-tags="tempdll latest"`动态初始化

### 系统
```go
     windows:
         386: -tags="tempdll latest"
         amd64: -tags="tempdll latest"
     windows(Windows 7, 8/8.1 and Windows Server 2012):
         386: -tags="tempdll 109"
         amd64: -tags="tempdll 109"
     linux(gtk3):
         amd64: -tags="tempdll latest"
         arm64: -tags="tempdll latest"
     linux(gtk2):
         amd64: -tags="tempdll 106"
         arm64: -tags="tempdll 106"
     macos:
         amd64: -tags="tempdll latest"
         arm64: -tags="tempdll latest"
```

```go
// TempDllDIR
//  DLL存放目录
type TempDllDIR int8

const (
    TddInvalid    TempDllDIR = iota - 1 // 无效
    TddTmp                              // 系统临时目录
    TddCurrent                          // 当前执行文件目录
    TddEnergyHome                       // Energy环境变量目录, 如果为空，则为系统临时目录
    TddCustom                           // 自定义目录, 如果为空，则为系统临时目录
)
```

### 发版
> 1. `github.com/energye/liblclbinres` vx.x.x 动态库字节码与发行版同步
> 2. [github.com/energye/liblclbinres/genbinres](https://github.com/energye/liblclbinres/genbinres) 生成`github.com/energye/liblclbinres`动态库字节码