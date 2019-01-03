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
	"e.coding.net/aliens/aliensboot_testserver/module/room/core"
	"e.coding.net/aliens/aliensboot_testserver/protocol"
)

//
func handleGameReady(authID int64, gateID string, request *protocol.GameReady) {
	room := core.RoomManager.GetRoomByPlayerID(authID)
	room.PlayerReady(authID)
}
