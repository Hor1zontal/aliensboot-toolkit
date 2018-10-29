GOGOPATH=${GOPATH}/src
protoc --proto_path=${GOPATH}:${GOGOPATH}:./ --gogofast_out=plugins=grpc:. *.proto


#code auto generate
./codegen -configPath config.yml