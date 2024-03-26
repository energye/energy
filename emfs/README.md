
----
embed.FS 内部静态资源集成
```golang
//resource 静态资源
//libs 库资源

//go:embed resource
var resource embed.FS

//go:embed libs
var libs embed.FS

func main() {
	inits.Init(&libs, &resource)
	...
	...
}
```