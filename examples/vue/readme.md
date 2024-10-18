在前端 [vite.config.ts](frontend%2Fvite.config.ts) 配置构建输出目录
```json
build:{
  outDir:"../resources"
}
```

1. node:
构建前端静态资源文件到配置好的 resources 目录
```cmd
cd frontend

npm run build
```

2. golang: 编译Go执行文件
```cmd
go build -ldflags="-s -w" -tags="prod"
```

