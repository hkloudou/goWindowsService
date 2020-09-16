set CGO_ENABLED=1
set GOARCH=386
set GOOS=windows

rsrc -manifest nac.manifest -o nac.syso -ico x.ico

go build -ldflags="-w -s" -o EtaxService.exe
pause
echo press any key continue