/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
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
	"aliens/log"
)

const (
	suspendedTimeOut = time.Millisecond * 500
	commandRequest = "request"
	commandReceive = "receive"
)

type handler func(request *base.Any) *base.Any

func NewRpcHandler(chanRpc *chanrpc.Server, handler handler) *rpcHandler {
	service := &rpcHandler{}
	service.chanRpc = chanRpc
	service.handler = handler
	service.chanRpc.Register(commandRequest, service.request)
	service.chanRpc.Register(commandReceive, service.receive)
	return service
}

type rpcHandler struct {
	chanRpc *chanrpc.Server
	handler handler
	suspended bool
}

func (this *rpcHandler) request(args []interface{}) {
	request := args[0].(*base.Any)
	server := args[1].(base.RPCService_RequestServer)
	response := this.handler(request)
	server.Send(response)
}

func (this *rpcHandler) receive(args []interface{}) {
	request := args[0].(*base.Any)
	this.handler(request)
}

//func (this *rpcHandler) LocalRequest(request *base.Any) (*base.Any, error) {
//	this.chanRpc.Call0(commandRequest, request, server)
//}
//
//func (this *rpcHandler) LocalReceive(request *base.Any) error {
//
//}

func (this *rpcHandler) Request(request *base.Any, server base.RPCService_RequestServer) error {
	if this.chanRpc != nil {
		this.chanRpc.Call0(commandRequest, request, server)
		return nil
	}
	return nil
}

func (this *rpcHandler) Receive(server base.RPCService_ReceiveServer) error {
	for {
		if this.suspended {
			time.Sleep(suspendedTimeOut)
			continue
		}
		request, err := server.Recv()
		if err != nil {
			log.Debug("accept async message error : %v", request)
			continue
		}
		if this.chanRpc != nil {
			this.chanRpc.Go(commandReceive, request)
		}
	}
	return nil
}


