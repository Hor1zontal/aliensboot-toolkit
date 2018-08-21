/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2017/3/24
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package service

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"net"
	"strconv"
	"aliens/log"
	"golang.org/x/net/context"
	"aliens/common/util"
	"reflect"
	"aliens/protocol/base"
)

type GRPCService struct {
	*CenterService

	//调用服务参数
	client *grpc.ClientConn

	//启动服务参数
	server  *grpc.Server      //

	requestClient base.RPCServiceClient
	receiveClient base.RPCService_ReceiveClient

	handler *rpcHandler //处理句柄
}

func (this *GRPCService) GetConfig() *CenterService {
	return this.CenterService
}

func (this *GRPCService) GetDesc() string {
	return "rpc service"
}

func (this *GRPCService) SetHandler(handler interface{}) {
	result, ok := handler.(*rpcHandler)
	if !ok {
		log.Fatalf("invalid grpc service request %v", reflect.TypeOf(handler))
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
	base.RegisterRPCServiceServer(server, this.handler)
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
	address := this.Address + ":" + util.IntToString(this.Port)
	conn, err := grpc.Dial(address, option)
	if err != nil {
		log.Errorf("did not connect: %v", err)
		return false
	}
	this.client = conn
	this.requestClient = base.NewRPCServiceClient(this.client)

	receiveClient, err := this.requestClient.Receive(context.Background())
	if err != nil {
		log.Errorf("init receive handler error : %v", err)
		return false
	}
	this.receiveClient = receiveClient
	log.Infof("connect grpc service %v-%v success", this.Name, address)
	return true
}

//比较服务是否冲突
func (this *GRPCService) Equals(other IService) bool {
	otherService, ok := other.(*GRPCService)
	if !ok {
		return false
	}
	return this.Name == otherService.Name && this.Address == otherService.Address && this.Port == this.Port
}

//服务是否本进程启动的
func (this *GRPCService) IsLocal() bool {
	return this.handler != nil
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
func (this *GRPCService) Request(request *base.Any) (*base.Any, error) {
	//服务为本机，直接处理
	if this.requestClient == nil {
		return nil, errors.New("service is not initial")
	}
	//TODO 后续可以考虑如果服务是当前进程启动，可以直接调用服务句柄
	if this.IsLocal() {
		//return this.handler
	}

	client, err := this.requestClient.Request(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return client.Recv()
}

//让服务接收消息，不需要响应
func (this *GRPCService) Send(request *base.Any) error {
	//服务为本机，直接处理
	if this.receiveClient == nil {
		return errors.New("service is not initial")
	}
	return this.receiveClient.Send(request)
}