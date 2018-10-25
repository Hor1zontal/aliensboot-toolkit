/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/8/20
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package service

import (
	"aliens/chanrpc"
	"aliens/protocol/base"
	"time"
	"google.golang.org/grpc"
	"aliens/log"
	"net"
	"strconv"
)

const (
	suspendedTimeOut = time.Millisecond * 500
	commandRequest = "request"
	commandReceive = "receive"
)

type handler func(request *base.Any) *base.Any

func NewRpcHandler(chanRpc *chanrpc.Server, handler handler) *rpcServer {
	if chanRpc == nil {
		log.Fatalf("chanRpc can not be nil")
	}
	service := &rpcServer{}
	service.chanRpc = chanRpc
	service.handler = handler
	service.chanRpc.Register(commandRequest, service.request)
	service.chanRpc.Register(commandReceive, service.receive)
	return service
}

type rpcServer struct {
	chanRpc   *chanrpc.Server
	handler   handler
	suspended bool
	//启动服务参数
	server  *grpc.Server      //
}

func (this *rpcServer) start(name string, port int) bool {
	server := grpc.NewServer()
	base.RegisterRPCServiceServer(server, this)
	address := ":" + strconv.Itoa(port)
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Errorf("failed to listen: %v", err)
		return false
	}
	go func() {
		server.Serve(lis)
		log.Infof("rpc service %v stop", name)
	}()
	this.server = server
	return true
}

func (this *rpcServer) close() {
	if this.server != nil {
		this.server.Stop()
	}
}

func (this *rpcServer) request(args []interface{}) {
	request := args[0].(*base.Any)
	server := args[1].(base.RPCService_RequestServer)
	response := this.handler(request)
	if response != nil {
		server.Send(response)
	}
}

func (this *rpcServer) receive(args []interface{}) {
	request := args[0].(*base.Any)
	this.handler(request)
}

//func (this *rpcServer) LocalRequest(request *base.Any) (*base.Any, error) {
//	this.chanRpc.Call0(commandRequest, request, server)
//}
//
//func (this *rpcServer) LocalReceive(request *base.Any) error {
//
//}

func (this *rpcServer) Request(request *base.Any, server base.RPCService_RequestServer) error {
	return this.chanRpc.Call0(commandRequest, request, server)
}

func (this *rpcServer) Receive(server base.RPCService_ReceiveServer) error {
	for {
		if this.suspended {
			time.Sleep(suspendedTimeOut)
			continue
		}
		request, err := server.Recv()
		if err != nil {
			//log.Debugf("accept async message error : %v", err)
			return err
		}
		this.chanRpc.Go(commandReceive, request)
	}
	return nil
}

