### Go语言实现的多进程IPC通信

### 简介
```text
基于 net socket 和 unix domain socket
```
### 通信-通道选择
| 系统      | 版本             | IPC        |
|---------|----------------|------------|
| windows | &lt; 10.17063  | net socket |
| windows | &gt;= 10.17063 | unix       |
| linux   | all            | unix       |
| macosx  | all            | unix       |
