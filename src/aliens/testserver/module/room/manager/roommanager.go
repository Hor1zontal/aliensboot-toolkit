/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/6/8
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package manager

import (
	"aliens/aliensbot/exception"
	"aliens/testserver/module/room/game"
	"aliens/testserver/protocol"
)

var RoomManager = &roomManager{rooms: make(map[string]*core.Room), players: make(map[int64]string)}

type roomManager struct {

	rooms map[string]*core.Room //运行的游戏  游戏id - 房间对象

	players map[int64]string //所有玩家的对应信息 玩家id - 房间id

}

//获取玩家在哪个房间
func (this *roomManager) GetRoomByPlayerID(playerID int64) *core.Room {
	roomID := this.players[playerID]
	if roomID == "" {
		exception.GameException(protocol.Code_roomNotFound)
	}
	return this.EnsureRoom(roomID)
}

//获取房间
func (this *roomManager) EnsureRoom(roomID string) *core.Room{
	game := this.rooms[roomID]
	if game == nil {
		exception.GameException(protocol.Code_roomNotFound)
	}
	return game
}

//分配新房间
func (this *roomManager) AllocRoom(appID string, players []*protocol.Player) *core.Room {
	room := core.NewRoom(appID)
	room.InitPlayers(players)

	this.rooms[room.GetID()] = room

	for _, player := range players {
		this.players[player.GetPlayerid()] = room.GetID()
	}
	return room
}

//关闭房间
func (this *roomManager) RemoveRoom(roomID string) {
	room := this.EnsureRoom(roomID)
	players := room.GetAllPlayer()
	for _, player := range players {
		delete(this.players, player.GetPlayerid())
	}
	room.Close()
	delete(this.rooms, roomID)
}

