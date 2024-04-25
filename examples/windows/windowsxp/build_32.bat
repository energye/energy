SET CGO_ENABLED=0

set GOARCH=386
set GOOS=windows
go build -ldflags "-s -w"