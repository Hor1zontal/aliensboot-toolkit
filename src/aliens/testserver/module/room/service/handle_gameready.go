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
	"aliens/testserver/module/room/manager"
	"aliens/testserver/protocol"
)




//
func handleGameReady(authID int64, gateID string, request *protocol.GameReady) {
	room := manager.RoomManager.GetRoomByPlayerID(authID)
	room.PlayerReady(authID)
}
