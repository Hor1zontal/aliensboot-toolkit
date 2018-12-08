// Code generated by aliensbot. DO NOT EDIT.
// source: module gate
package rpc

import (
	"aliens/testserver/protocol"
)

var Gate = &gateRPCHandler{&rpcHandler{name:"gate"}}


type gateRPCHandler struct {
	*rpcHandler
}




func (this *gateRPCHandler) KickOut(node string, request *protocol.KickOut) error {
	message := &protocol.Request{
		Gate:&protocol.Request_KickOut{
			KickOut:request,
		},
	}
	return this.Send(node, message)
}

func (this *gateRPCHandler) PushMessage(node string, request *protocol.PushMessage) error {
	message := &protocol.Request{
		Gate:&protocol.Request_PushMessage{
			PushMessage:request,
		},
	}
	return this.Send(node, message)
}

func (this *gateRPCHandler) BindService(node string, request *protocol.BindService) error {
	message := &protocol.Request{
		Gate:&protocol.Request_BindService{
			BindService:request,
		},
	}
	return this.Send(node, message)
}
