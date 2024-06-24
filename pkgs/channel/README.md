### Multi process IPC communication implemented in Go

### introduction
```text
based net socket & unix domain socket
```
### Communication - Channel Selection
| OS Version | version        | IPC        |
|------------|----------------|------------|
| windows    | &lt; 10.17063  | net socket |
| windows    | &gt;= 10.17063 | unix       |
| linux      | all            | unix       |
| macosx     | all            | unix       |
