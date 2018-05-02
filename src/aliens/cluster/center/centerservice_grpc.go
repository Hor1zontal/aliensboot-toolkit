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
	"aliens/protocol"

)

func PublicGRPCService(config ServiceConfig, handle protocol.RPCServiceServer) *GRPCService {
	if !ClusterCenter.IsConnect() {
		panic(config.Name + " cluster center is not connected")
		return nil
	}

	server := grpc.NewServer()
	protocol.RegisterRPCServiceServer(server,  handle)

	if config.Address == "" {
		config.Address = util.GetAddress(config.Port)
	}

	service := &GRPCService{
		centerService: &centerService{
			id:          ClusterCenter.GetNodeID(),
			serviceType: config.Name,
			Address:     config.Address,
			Protocol: GRPC,
		},
		port: config.Port,
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


type GRPCService struct {
	*centerService

	//调用服务参数
	client *grpc.ClientConn
	caller protocol.RPCServiceClient

	//启动服务参数
	port    int
	server  *grpc.Server      //
}

func (this *GRPCService) GetDesc() string {
	return "rpc service"
}

func (this *GRPCService) GetID() string {
	return this.id
}

func (this *GRPCService) GetType() string {
	return this.serviceType
}

func (this *GRPCService) SetID(id string) {
	this.id = id
}

func (this *GRPCService) SetType(serviceType string) {
	this.serviceType = serviceType
}

//启动服务
func (this *GRPCService) Start() bool {
	if this.server == nil {
		log.Error("invalid service param")
		return false
	}
	address := ":" + strconv.Itoa(this.port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Errorf("failed to listen: %v", err)
		return false
	}
	go func() {
		this.server.Serve(lis)
		log.Infof("rpc service %v stop", this.serviceType)
	}()
	return true
}

//连接服务
func (this *GRPCService) Connect() bool {
	//rpc := rpcClientFactories[this.serviceType]
	//if rpc == nil {
	//	log.Error("grpc mapping not register %v", this.serviceType)
	//	return false
	//}

	conn, err := grpc.Dial(this.Address, grpc.WithInsecure())
	if err != nil {
		log.Errorf("did not connect: %v", err)
		return false
	}
	this.client = conn
	this.caller = protocol.NewRPCServiceClient(this.client)
	//rpc(this.client)
	//mutable := reflect.ValueOf(caller).Elem()
	//this.caller = mutable.Addr().MethodByName(method)
	//if !this.caller.IsValid() {
	//	log.Error("grpc %v request method not found", this.serviceType)
	//	this.client.Close()
	//	return false
	//}
	return true
}

//比较服务是否冲突
func (this *GRPCService) Equals(other IService) bool {
	otherService, ok := other.(*GRPCService)
	if !ok {
		return false
	}
	return this.serviceType == otherService.serviceType && this.Address == otherService.Address
}

//服务是否本进程启动的
func (this *GRPCService) IsLocal() bool {
	return this.server != nil
}

//关闭服务
func (this *GRPCService) Close() {
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
func (this *GRPCService) Request(request interface{}) (interface{}, error) {
	//服务为本机，直接处理
	if this.client == nil {
		return nil, errors.New("service is not initial")
	}
	//params := make([]reflect.Value, 2)
	//params[0] = reflect.ValueOf(context.Background())
	//params[1] = reflect.ValueOf(request)
	//results := this.caller.Call(params)
	//
	//response := results[0].Interface()
	requestAny, error := request.(*protocol.Any)
	if !error {
		return nil, errors.New("invalid request type")
	}
	return this.caller.Request(context.Background(), requestAny)
}