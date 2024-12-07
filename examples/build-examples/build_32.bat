SET CGO_ENABLED=0

currentDIR=%CD%

cd %1

echo %CD%

set GOARCH=386
set GOOS=windows
go build --ldflags="-s -w"

cd %currentDIR%