# {{.Name}} Project

## 目录说明 - Directory description
> resources 
>> cn: 资源存放目录, 可自定义目录名
>>
>> en: Resource storage directory, customizable directory name
>
>  build
>> 自动生成目录: 用于编译、构建、生成安装包
>>
>> Automatically generate directory: used for compiling, building, and generating installation packages

## 文件说明 - File description
> energy.json
>> cn: 项目配置文件, 用于构建和生成安装程序, 文件名不可更改.
>>
>> en: project configuration file, which is used to build and generate the installer, has an unchangeable file name.
> 
> go.mod
>> cn: go模块依赖管理, 文件名不可更改.
>>
>> en: go module dependency management, has an unchangeable file name.
> 
> main.go
>> cn: go energy应用启动入口文件
>>
>> en: go energy application enables the entry file
>
> go.sum No need to change

### cn: 以上可以根据实际情况变更
### en: The above can be changed according to the actual situation

## 运行应用 - Run Application
> Windows, Linux: `go run main.go`
> 
> MacOS: `go run main.go env=dev`

## 构建应用 - Building Applications
> Use Go: 
>> Windows: `go build -ldflags "-H windowsgui -s -w"`
>> 
>> Linux: `go build -ldflags "-s -w"`
>>
>> MacOS: `go build -ldflags "-s -w"`
> 
> Use Energy
>> `energy build`
> 

## 制作安装包 - Making installation packages
> 1. 构建应用 - Building Applications
>> `energy build`
> 
> 2. 执行制作安装包命令 - Run the create installation package command
>> `energy package`
>>
>> [Reference link](https://energye.github.io/course/build-package)