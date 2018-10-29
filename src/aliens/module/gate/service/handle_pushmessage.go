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
	"aliens/module/gate/network"
	"aliens/protocol"
	"aliens/protocol/base"
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
