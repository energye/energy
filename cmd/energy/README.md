# Energy 命令行工具
## 自动配置开发环境
## 使用方式
> 需要安装[Golang](https://golang.google.cn/dl/)环境
>> go get -u github.com/energye/energy
>>
>> 进入 cmd/energy 目录
>> 
>> 执行命令 "go install" 安装命令行工具
>> 
>> go install
>>
>
> 三、命令参数
> 
>> energy install [path] [version]
>>

| 名称      | 描述                                      |
|---------|-----------------------------------------|
| install | 安装energy框架, 需要连接互联网, 自动下载CEF和Energy框架环境 |
| name    | 可选参数, 目录名称, 默认EnergyFramework                 |
| path    | 可选参数, 安装目录, 默认当前目录生成[name]文件夹           |
| version | 可选参数, 版本号v1.1.0, 默认最新版本latest           |

>示例
>> 1. `energy install .`
>>
>> 2. `energy install --path=/app/energyFramework -v x.x.x .`
>>