### 生成 liblcl 到 xx.go 字节数组发版包
已规定好的命名规则和文件目录
1. 压缩包
2. 文件目录

### 目录
当前系统用户目录/golcl/

### 命名
#### 文件
所属目录: /xxx/用户目录/golcl/

 规则

| liblcl 文件目录                | Go 字节数组                    |
|----------------------------|----------------------------|
| win32/liblcl.dll           | liblcl_windows_386.go      |
| win64/liblcl.dll           | liblcl_windows_amd64.go    |
| macos64-cocoa/liblcl.dylib | liblcl_darwin_amd64.go     |
| linux64-gtk3/liblcl.so     | liblcl_gtk3_linux_amd64.go |
| linux64-gtk2/liblcl.so     | liblcl_gtk2_linux_amd64.go |