// Code generated by aliensbot. DO NOT EDIT.
// source: module scene
package rpc

import (
	"aliens/testserver/protocol"
)

var Scene = &sceneRPCHandler{&rpcHandler{name:"scene"}}


type sceneRPCHandler struct {
	*rpcHandler
}




func (this *sceneRPCHandler) EntityCall(node string, request *protocol.EntityCall) error {
	message := &protocol.Request{
		Scene:&protocol.Request_EntityCall{
			EntityCall:request,
		},
	}
	return this.Send(node, message)
}

func (this *sceneRPCHandler) LoginScene(node string, request *protocol.LoginScene) error {
	message := &protocol.Request{
		Scene:&protocol.Request_LoginScene{
			LoginScene:request,
		},
	}
	return this.Send(node, message)
}
