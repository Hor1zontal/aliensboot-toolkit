GOGOPATH=${GOPATH}/src
MODULENAME=passport
protoc --proto_path=${GOPATH}:${GOGOPATH}:./ --gofast_out=plugins=grpc:. *.proto

../codegen -proto protocol.proto -template ../templates/common.template        -output ../../module/${MODULENAME}/service/handle.go     -overwrite true
../codegen -proto protocol.proto -template ../templates/common_json.template   -output ../../module/${MODULENAME}/service/handlejson.go -overwrite true
../codegen -proto protocol.proto -template ../templates/common_handle.template -output ../../module/${MODULENAME}/service/  -prefix 'handle_${}.go'
../codegen -proto protocol.proto -template ../templates/service.template       -output ../../module/${MODULENAME}/service/service.go   -overwrite true
