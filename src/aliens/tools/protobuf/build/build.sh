GOROOT=/usr/local/Cellar/go/1.10/libexec #gosetup
PROJECT=/Users/hejialin/git/server/aliensbot
GOLIB=/Users/hejialin/git/server/alienslib
GOPATH=${PROJECT}:${GOLIB} #gosetup
${GOROOT}/bin/go build -o codegen ${PROJECT}/src/aliens/tools/protobuf/build/service_builder.go  #gosetup
