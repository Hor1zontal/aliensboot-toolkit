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

import (
	"aliens/protocol/room"
	"aliens/module/room/core"
)


//
func handleCreateRoom(request *room.CreateRoom, response *room.CreateRoomRet) {
	roomID := core.Manager.CreateRoom(uint8(request.MaxPlayer))
	response.RoomID = roomID
	response.Result = room.RoomResult_success
}
