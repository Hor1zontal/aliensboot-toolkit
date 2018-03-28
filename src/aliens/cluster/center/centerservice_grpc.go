/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2017/3/24
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package center

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"net"
	"strconv"
	"aliens/log"
	"golang.org/x/net/context"
	"aliens/common/util"
	"reflect"
)

var rpcClientFactories = make(map[string]func(cc *grpc.ClientConn) interface{})

const method = "Request"


func RegisterRPCClientFactory(serviceType string, clientFactory func(cc *grpc.ClientConn) interface{}) {
	rpcClientFactories[serviceType] = clientFactory
}

func PublicRPCService(serviceType string, port int, server *grpc.Server) *gRPCService {
	if !ClusterCenter.IsConnect() {
		panic(serviceType + " cluster center is not connected")
		return nil
	}
	service := &gRPCService{
		centerService: &centerService{
			id:          GetServerNode(),
			serviceType: serviceType,
			Address:     util.GetAddress(port),
			Protocol: GRPC,
		},
		port: port,
		server : server,
	}
	if !service.Start() {
		panic(service.serviceType + " rpc service can not be start")
	}
	//RPC启动成功,则发布到中心服务器
	if !ClusterCenter.PublicService(service) {
		panic(service.serviceType + " rpc service can not be start")
	}
	return service
}


type gRPCService struct {
	*centerService

	//调用服务参数
	client *grpc.ClientConn
	caller reflect.Value

	//启动服务参数
	port    int
	server  *grpc.Server      //
}

func (this *gRPCService) GetDesc() string {
	return "rpc service"
}

func (this *gRPCService) GetID() string {
	return this.id
}

func (this *gRPCService) GetType() string {
	return this.serviceType
}

func (this *gRPCService) SetID(id string) {
	this.id = id
}

func (this *gRPCService) SetType(serviceType string) {
	this.serviceType = serviceType
}

//启动服务
func (this *gRPCService) Start() bool {
	if this.server == nil {
		log.Error("invalid service param")
		return false
	}
	address := ":" + strconv.Itoa(this.port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Error("failed to listen: %v", err)
		return false
	}
	go func() {
		this.server.Serve(lis)
		log.Info("rpc service %v stop", this.serviceType)
	}()
	return true
}

//连接服务
func (this *gRPCService) Connect() bool {
	rpc := rpcClientFactories[this.serviceType]
	if rpc == nil {
		log.Error("grpc mapping not register %v", this.serviceType)
		return false
	}

	conn, err := grpc.Dial(this.Address, grpc.WithInsecure())
	if err != nil {
		log.Error("did not connect: %v", err)
		return false
	}
	this.client = conn

	caller := rpc(this.client)
	mutable := reflect.ValueOf(caller).Elem()
	this.caller = mutable.Addr().MethodByName(method)
	if !this.caller.IsValid() {
		log.Error("grpc %v request method not found", this.serviceType)
		this.client.Close()
		return false
	}
	return true
}

//比较服务是否冲突
func (this *gRPCService) Equals(other IService) bool {
	otherService, ok := other.(*gRPCService)
	if !ok {
		return false
	}
	return this.serviceType == otherService.serviceType && this.Address == otherService.Address
}

//服务是否本进程启动的
func (this *gRPCService) IsLocal() bool {
	return this.server != nil
}

//关闭服务
func (this *gRPCService) Close() {
	if this.server != nil {
		this.server.Stop()
		this.server = nil
	}
	if this.client != nil {
		this.client.Close()
		this.client = nil
	}
}

//向服务请求消息
func (this *gRPCService) Request(request interface{}) (interface{}, error) {
	//服务为本机，直接处理
	if this.client == nil {
		return nil, errors.New("service is not initial")
	}
	params := make([]reflect.Value, 2)
	params[0] = reflect.ValueOf(context.Background())
	params[1] = reflect.ValueOf(request)
	results := this.caller.Call(params)

	response := results[0].Interface()
	/*if len(results) == 2 {
		return response, re
	}*/
	return response, nil
}