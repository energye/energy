SET CGO_ENABLED=0


set GOARCH=amd64
set GOOS=linux
go build -trimpath -ldflags "-s -w" -o energy-linux64 energy.go
md5.exe energy-linux64

pause