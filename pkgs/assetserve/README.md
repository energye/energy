### 简易的内置静态资源http服务
```go
由于使用go http server该服务可能会引起一些安全软件报毒
报毒解决方案
    1. 配置 ssl 证书, 未测试
    2. 编译时 使用 -ldflags "-s -w" 去除调试信息和符号
    3. 使用 upx 工具压缩执行文件
最好是 2 和 3 配合使用
```

### 示例
```go
// 创建http服务
server := assetserve.NewAssetsHttpServer()
server.PORT = 22022 //设置端口号
server.AssetsFSName = "resources"   // 使用go内置资源go:embed, 指定资源目录名
server.Assets = &resources          // 使用go内置资源go:embed, 设置embed.FS引用
//server.LocalAssets = "/to/path/"  // 使用本地资源目录, 指定本地目录
// 启动http服务
server.StartHttpServer()            // 阻塞进程
```

### 安全配置
#### 防止应用外访问内置资源
```go
使用http请求头参数验证资源请求源有效性
该配置应在全局, 存在于主进程和子进程中

// 设置自定义头
assetserve.AssetsServerHeaderKeyName = "energy"
// 设置自定义值
assetserve.AssetsServerHeaderKeyValue = "energy"
```