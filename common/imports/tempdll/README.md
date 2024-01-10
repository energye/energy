### liblcl.xx 动态库字节码存放方式

### 介绍
> 将liblcl打包进执行文件中, 运行时根据TempDllDIR配置释放到指定目录

### 使用
#### 该方式在Go编译时将执行文件内置到exe中

#### 前提条件
1. 在Go main函数初始化全局配置时[cef.GlobalInit(libs, resources)]设置libs内嵌对象参数
2. 目录名默认 libs

### 系统
```go
     除MacOS都可使用该方式, MacOS采用固定的xxx.app包内加载
```

### 配置
参考: tempdll.TempDLL