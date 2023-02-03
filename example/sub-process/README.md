```text
这个示例演示了 主进程和 子进程相互独立出来
windows 和 linux 配置和使用上没有区别
macos 有明显的区别

windows linux MacOS 的使用
步骤
 1. 先编译好子进程程序 
    cd sub-process 
    go build
    sub-process.exe 
 2. 目录配置
    windows linux: 将子进程执行文件（sub-process.exe）在主进程SetBrowseSubprocessPath配置，
        如果在 FrameworkDirPath 可以直接写文件名
    MacOS: 将子进程执行文件（sub-process.exe）在主进程 macapp.MacApp.SetBrowseSubprocessPath 配置绝对路径
 3. 运行主程序
 
 示例目录说明
     main-process 主进程
        resources 资源目录
     sub-process 子进程
 
 
 二 MacOS 的使用 - 手动创建独立子进程 app 包
 说明:
   1. MacOS 下多进程默认是主子进程，在开发时生成的MacOS程序包是以主子进程方式
   2. 如果在程序中明显区别出主子进程，需要将编译好的子进程分别手动放入macos程序包中
   3. 不需要指定 IsCEF = true , SetBaseCefFrameworksDir 和 SetBrowseSubprocessPath
 手动创建-MacOS App目录结构参考 以goland开发工具示例:
   1. 在程序启动后，控制台输出的前几行中，有程序临时输出目录 类似（/private/var/folders/c1/6d_0g9xj68g0fzkm06xtx84r0000gn/T/GoLand/xxxx_go）
   2. 找到 xxxx_go 程序包
   3. 进入frameworks目录
      其中 以下几个子包即是子程序
         xxx_go Helper 
         xxx_go Helper(GPU) 
         xxx_go Helper(Plugin) 
         xxx_go Helper(Render) 
   4. 将自己编译的子程序以这上面的包名格式命名，并放入替换到每个程序包的MacOS目录中即可
```