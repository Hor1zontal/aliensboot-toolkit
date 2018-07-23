/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/3/30
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package service

import "aliens/protocol/room"
import (
	"aliens/network"
	"aliens/module/room/core"
)


//
func handleJoinRoom(request *room.JoinRoom, response *room.JoinRoomRet, agent network.Agent) {
	core.Manager.AcceptRoomMessage(request.GetRoomID(), request, response, agent)
}
