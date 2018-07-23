GOGOPATH=${GOPATH}/src
MODULENAME=room
protoc --proto_path=${GOPATH}:${GOGOPATH}:./ --gofast_out=plugins=grpc:. *.proto

../codegen -proto room.proto -template ../agent_common.template        -output ../../module/${MODULENAME}/service/handle.go     -overwrite true
../codegen -proto room.proto -template ../agent_common_json.template   -output ../../module/${MODULENAME}/service/handlejson.go -overwrite true
../codegen -proto room.proto -template ../agent_common_handle.template -output ../../module/${MODULENAME}/service/  -prefix 'handle_${}.go'

../codegen -proto room.proto -template ../common_rpc.template -output ../../module/cluster/rpc/${MODULENAME}.go  -overwrite true

../codegen -proto room.proto -template ../service.template       -output ../../module/${MODULENAME}/service/service.go   -overwrite true

