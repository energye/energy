SET CGO_ENABLED=1

set GOARCH=386
set GOOS=windows
go build -o energy-32.exe energy.go

set GOARCH=amd64
set GOOS=windows
go build -o energy-64.exe energy.go

set GOARCH=amd64
set GOOS=darwin
go build -o energy-darwin-64 energy.go

set GOARCH=amd64
set GOOS=linux
go build -o energy-linux-64 energy.go

pause