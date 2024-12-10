SET CGO_ENABLED=0

set GOARCH=386
set GOOS=windows
go build -trimpath -ldflags "-s -w" -o energy-windows-32.exe energy.go
md5.exe energy-windows-32.exe

set GOARCH=amd64
set GOOS=windows
go build -trimpath -ldflags "-s -w" -o energy-windows-64.exe energy.go
md5.exe energy-windows-64.exe

set GOARCH=amd64
set GOOS=darwin
go build -trimpath -ldflags "-s -w" -o energy-darwin-64 energy.go
md5.exe energy-darwin-64

set GOARCH=arm64
set GOOS=darwin
go build -trimpath -ldflags "-s -w" -o energy-darwinarm-64 energy.go
md5.exe energy-darwinarm-64

set GOARCH=386
set GOOS=linux
go build -trimpath -ldflags "-s -w" -o energy-linux-32 energy.go
md5.exe energy-linux-32

set GOARCH=amd64
set GOOS=linux
go build -trimpath -ldflags "-s -w" -o energy-linux-64 energy.go
md5.exe energy-linux-64

set GOARCH=arm64
set GOOS=linux
go build -trimpath -ldflags "-s -w" -o energy-linuxarm-64 energy.go
md5.exe energy-linuxarm-64

set GOARCH=arm
set GOOS=linux
go build -trimpath -ldflags "-s -w" -o energy-linuxarm energy.go
md5.exe energy-linuxarm

pause