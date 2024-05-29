
这个示例演示在其它语言调用energy编译为dll库调用示例

简单起见，提供了两种语言. 分别为: lazarus和python

为运行成功，需要安装对应语言环境

go编译为dll需要使用cgo
- mingw64 (cgo)
- golang (go开发环境)

示例运行，两种语言任意
- lazarus 下载 https://www.lazarus-ide.org/ 
- python 环境

步骤

以python为例（因为它最简单）

准备好环境后
1. 运行 build_x64.bat 编译出 libenergy.dll
2. cmd 进入dll示例目录: 执行命令 python pyLoadLibenergy.py