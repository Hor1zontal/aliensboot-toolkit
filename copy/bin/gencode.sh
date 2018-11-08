#!/bin/sh
basepath=$(cd `dirname $0`; pwd)


cd ../src/aliens/testserver/protocol/
GOGOPATH=${GOPATH}/src; protoc --proto_path=${GOPATH}:${GOGOPATH}:./ --gogofast_out=plugins=grpc:. *.proto

#生成服务代码
cd $basepath/..

modules=(game gate passport)

for i in "${!modules[@]}"; do
	aliensbot module gen ${modules[$i]}
done