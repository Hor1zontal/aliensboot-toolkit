/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/3/30
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package service

import (
	"aliens/aliensbot/protocol/base"
	"aliens/mmoserver/module/gate/network"
	"aliens/mmoserver/protocol"
)

//
func handlePushMessage(request *protocol.PushMessage) {
	msg := &base.Any{Id: 1000, Value: request.GetData()}
	authID := request.GetAuthID()
	if authID == -1 {
		network.Manager.Broadcast(msg)
	} else if authID > 0 {
		network.Manager.Push(request.GetAuthID(), msg)
	}
}
