protoc --gofast_out=plugins=grpc:. *.proto

../codegen -proto protocol.proto -template ../common.template        -output ../../module/passport/service/handle.go     -overwrite true
../codegen -proto protocol.proto -template ../common_json.template   -output ../../module/passport/service/handlejson.go -overwrite true
../codegen -proto protocol.proto -template ../common_handle.template -output ../../module/passport/service/  -prefix 'handle_${}.go'