gogopath=${GOPATH}/src/github.com/gogo/protobuf/protobuf
echo $gogopath
protoc --gogo_out=plugins=grpc:. *.proto