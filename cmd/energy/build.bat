SET CGO_ENABLED=0

set GOARCH=386
set GOOS=windows
go build -ldflags "-s -w" -o energy-windows-32.exe energy.go

set GOARCH=amd64
set GOOS=windows
go build -ldflags "-s -w" -o energy-windows-64.exe energy.go

set GOARCH=amd64
set GOOS=darwin
go build -ldflags "-s -w" -o energy-darwin-64 energy.go

set GOARCH=arm64
set GOOS=darwin
go build -ldflags "-s -w" -o energy-darwinarm-64 energy.go

set GOARCH=amd64
set GOOS=linux
go build -ldflags "-s -w" -o energy-linux-64 energy.go

set GOARCH=arm64
set GOOS=linux
go build -ldflags "-s -w" -o energy-linuxarm-64 energy.go

pause