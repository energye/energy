SET CGO_ENABLED=0

set GOARCH=amd64
set GOOS=windows
go build -ldflags "-s -w"