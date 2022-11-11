**# Energy 命令行工具

## 使用方式
> 一、使用已编译
>
>> [下载地址]()
>>
>> 配置到环境变量 或 直接在命令行中执行
> 
> 二、自行编译
>
> 需要安装[Go环境]()
>>go get -u github.com/energye/energy
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
| path    | 可选参数, 指定安装目录, 默认当前目录生成CEFEnergy文件夹      |
| version | 可选参数, 指定版本号v1.1.0, 默认最新版本latest               |

>示例
>> 1. energy install
>>
>> 2. energy install --path=/app/energyFramework --version=v1.1.0

----
### 授权
### [![License GPL 3.0](https://img.shields.io/badge/License%20GPL3.0-green)](https://opensource.org/licenses/GPL-3.0)
