::GOROOT=/usr/local/Cellar/go/1.11/libexec #gosetup
set PROJECT=G:\Server\aliensbot
set GOLIB=G:\botLibrary
set GOPATH=%PROJECT%;%GOLIB%
go build -o aliensbot.exe %PROJECT%/src/aliens/toolkit/main.go
move aliensbot.exe %ALIENSBOT_HOME%/bin