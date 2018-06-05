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

type GRPCService struct {
	*ServiceConfig

	//调用服务参数
	client *grpc.ClientConn
	caller protocol.RPCServiceClient

	//启动服务参数
	server  *grpc.Server      //

	handler protocol.RPCServiceServer //处理句柄
}

func (this *GRPCService) GetConfig() *ServiceConfig {
	return this.ServiceConfig
}

func (this *GRPCService) GetDesc() string {
	return "rpc service"
}

func (this *GRPCService) SetHandler(handler interface{}) {
	result, ok := handler.(protocol.RPCServiceServer)
	if !ok {

	}

	this.handler = result
}

//启动服务
func (this *GRPCService) Start() bool {
	//var server *grpc.Server = nil
	//if ClusterCenter.certFile != "" {
	//	creds, err := credentials.NewServerTLSFromFile(ClusterCenter.certFile, ClusterCenter.keyFile)
	//	if err != nil {
	//		log.Fatalf("Failed to generate credentials %v", err)
	//		return false
	//	}
	//	server = grpc.NewServer(grpc.Creds(creds))
	//} else {
	//	server = grpc.NewServer()
	//}

	server := grpc.NewServer()
	protocol.RegisterRPCServiceServer(server, this.handler)
	if this.Address == "" {
		this.Address = util.GetIP()
	}
	this.server = server

	address := ":" + strconv.Itoa(this.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Errorf("failed to listen: %v", err)
		return false
	}
	go func() {
		this.server.Serve(lis)
		log.Infof("rpc service %v stop", this.Name)
	}()
	return true
}

//连接服务
func (this *GRPCService) Connect() bool {
	var option grpc.DialOption = nil
	//if ClusterCenter.certFile != "" {
	//	creds, err := credentials.NewClientTLSFromFile(ClusterCenter.certFile, ClusterCenter.commonName)
	//	if err != nil {
	//		log.Errorf("Failed to create TLS credentials %v", err)
	//		return false
	//	}
	//	option = grpc.WithTransportCredentials(creds)
	//} else {
	//	option = grpc.WithInsecure()
	//}
	option = grpc.WithInsecure()

	conn, err := grpc.Dial(this.Address + ":" + util.IntToString(this.Port), option)
	if err != nil {
		log.Errorf("did not connect: %v", err)
		return false
	}
	this.client = conn
	this.caller = protocol.NewRPCServiceClient(this.client)
	return true
}

//比较服务是否冲突
func (this *GRPCService) Equals(other IService) bool {
	otherService, ok := other.(*GRPCService)
	if !ok {
		return false
	}
	return this.Name == otherService.Name && this.Address == otherService.Address
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
	requestAny, error := request.(*protocol.Any)
	if !error {
		return nil, errors.New("invalid request type")
	}
	return this.caller.Request(context.Background(), requestAny)
}