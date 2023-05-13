### i18n 多语言资源

### 资源文件
#### 文件名格式
- locale.[lang].json | locale.[lang].ini => locale.en-US.json | locale.zh-CN.ini
#### 内容格式
- locale.[lang].json
```json
{
  "name": "value",
  "name2": "value2",
  ...
}
```
- locale.[lang].ini
```ini
name=value
name2=value2
...
```

#### 使用本地加载资源
```go
i18n.SetLocalPath(localPath string)
```

#### 使用内置FS加载资源
```go
i18n.SetLocalFS(localFS *embed.FS, localFSPath string)
```

#### 语言切换
```go
i18n.Switch(lang consts.LANGUAGE)
```

#### 静态资源注册
```go
i18n.RegisterResource(name, value string)
```

#### 变量资源注册
```go
i18n.RegisterVarResource(name string, value *string)
```

#### 获取资源
```go
i18n.Resource(name string) string
```