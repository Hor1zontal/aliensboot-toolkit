#GOROOT=/usr/local/Cellar/go/1.11/libexec #gosetup
PROJECT=/Users/hejialin/git/server/aliensbot
GOLIB=/Users/hejialin/git/server/alienslib
GOPATH=${PROJECT}:${GOLIB} #gosetup
${GOROOT}/bin/go build -o aliensbot ${PROJECT}/src/aliens/toolkit/main.go  #gosetup
mv aliensbot ${ALIENSBOT_HOME}/bin