Code from [go-rod](https://github.com/go-rod/rod)

https://github.com/go-rod/rod

在rod基础上增加了 energy.go 的扩展, 集成自 energy 的 CEF 和 LCL 

通信方式与 rod 不同, energy 直接使用 API `SendDevToolsMessage` 或 `ExecuteDevToolsMethod` 方式操作 devtools, 在 `OnDevToolsRawMessage`, `OnDevToolsEvent`, `OnDevToolsMethodRawResult` 事件回调中接收消息

不需要远程调试端口配置 `SetRemoteDebuggingPort`

使用 SendDevToolsMessage, 因为更方便

已处理 
```
OnDevToolsRawMessage
```

未处理 
```
OnDevToolsEvent 
OnDevToolsMethodRawResult
```

备注: 将来这个示例 rod 将被移除, 使用单独仓库