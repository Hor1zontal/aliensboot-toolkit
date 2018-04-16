GOROOT=/usr/local/Cellar/go/1.10/libexec #gosetup
PROJECT=/Users/hejialin/git/server/kylin
GOPATH=${PROJECT}:/Users/hejialin/git/server/alienslib #gosetup
${GOROOT}/bin/go build ${PROJECT}/src/aliens/tools/zookeeper/main/main.go #gosetup
