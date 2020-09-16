# goWindowsService
> what we do?
> 
> Step1.Run Service as SYSTEM
> 
> Step2.Run app.exe(config.go) as an administrator(current user) when your system startup
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

### windows（recommend）
> run build.bat

### mac
```
brew install mingw-w64

env CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc go build -ldflags="-w -s" -o service.exe service/*.go

```
## 5.copy
1. copy your allpication "app.exe"(same as config.go) to output dictionary
2. pack your output dictionary as a zip
3. run output/install.bat
> TIP: add -ldflags="-w -s -H windowsgui" to build windowless app when you biuld your own app.exe.

# LINKS
1. https://gist.github.com/LiamHaworth/1ac37f7fb6018293fc43f86993db24fc
2. https://github.com/joesilva01862/ScreenshotWindowsService