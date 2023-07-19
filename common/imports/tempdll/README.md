### liblcl.xx 动态库字节码存放方式

### 介绍
> 1. 该方式可不必下载`liblcl.xx`动态库, 使用`energy`时,只需安装CEF框架
> 2. 原理是将`liblcl.xx`动态库写入到`liblclbinres/xxx.go`文件中以字节码形式存放
> 3. 使用该模式在`liblclbinres/xxx.go`读取`liblcl.xx`动态库字节码并释放到本机目录

### 使用
> 1. 编译时增加编译命令参数 `-tags="tempdll"`
>> `go build -tags="tempdll"`
> 2. 开发工具中以Goland为例
>> 在运行配置中 `Go tool arguments` 中配置 `-tags="tempdll"`
> 3. `TempDLL` 全局变量配置`liblcl.xx`动态库保存目录
> 4. `TempDLL` 全局变量根据编译参数`-tags="tempdll"`动态初始化
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
> 1. `github.com/energye/liblclbinres` 动态库字节码与发行版同步
> 2. [genbinres](..%2F..%2Ftools%2Fgenbinres) 生成`github.com/energye/liblclbinres`动态库字节码