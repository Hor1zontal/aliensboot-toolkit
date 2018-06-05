GOGOPATH=${GOPATH}/src
MODULENAME=scene
protoc --proto_path=${GOPATH}:${GOGOPATH}:./ --gofast_out=plugins=grpc:. *.proto

../codegen -proto protocol.proto -template ../common.template        -output ../../module/${MODULENAME}/service/handle.go     -overwrite true
../codegen -proto protocol.proto -template ../common_json.template   -output ../../module/${MODULENAME}/service/handlejson.go -overwrite true
../codegen -proto protocol.proto -template ../common_handle.template -output ../../module/${MODULENAME}/service/  -prefix 'handle_${}.go'
../codegen -proto protocol.proto -template ../service.template       -output ../../module/${MODULENAME}/service/service.go   -overwrite true
