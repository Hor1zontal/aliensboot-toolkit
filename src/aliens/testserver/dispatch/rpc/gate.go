/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/6/15
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package rpc

import (
	"aliens/testserver/dispatch"
	"aliens/testserver/protocol"
)

var Gate = &gateRPCHandle{"gate"}

type gateRPCHandle struct {
	name string
}

func (this *gateRPCHandle) KickOut(node string, request *protocol.KickOut) error {
	message := &protocol.Request{
		Gate: &protocol.Request_KickOut{
			KickOut: request,
		},
	}
	return dispatch.SendNodeMessage(this.name, node, message)
}

func (this *gateRPCHandle) BindService(node string, request *protocol.BindService) error {
	message := &protocol.Request{
		Gate: &protocol.Request_BindService{
			BindService: request,
		},
	}
	return dispatch.SendNodeMessage(this.name, node, message)
}

func (this *gateRPCHandle) PushMessage(node string, request *protocol.PushMessage) error {
	message := &protocol.Request{
		Gate: &protocol.Request_PushMessage{
			PushMessage: request,
		},
	}
	return dispatch.SendNodeMessage(this.name, node, message)
}
