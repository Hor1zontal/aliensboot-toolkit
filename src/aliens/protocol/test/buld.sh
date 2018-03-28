export GOPATH=/Users/hejialin/git/server/kylin;/Users/hejialin/git/server/alienslib

protoc -i /usr/local/include -i. \ -I$GOaPATH/src \ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \ --go_out=google/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \test.proto