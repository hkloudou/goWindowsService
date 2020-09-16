# goWindowsService
> what we do?
> 
> Step1.Run Serice as SYSTEM
> 
> Step2.Run a program as an administrator(current user)
## 1.Star this project ,haha
## 2.clone to your workspace
```
go get github.com/hkloudou/goWindowsService
```
## 3.config.go
```
package main

var serviceName = "Your Service Name"
var serviceDisplayName = "Your ServiceDisplayName"
var serviceDescription = "https://github.com/hkloudou/goWindowsService"

var appPath = "app.exe"   // your app path

```
## 4.build

### windos（recommend）
> run build.bat

### mac
```
brew install mingw-w64

env CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc go build -ldflags="-w -s" -o service.exe service/*.go
```
## 5.copy
> copy your allpication "app.exe"(same as config.go) to output dictionary
>
> pack your output dictionary as a zip
>
> run output/install.bat