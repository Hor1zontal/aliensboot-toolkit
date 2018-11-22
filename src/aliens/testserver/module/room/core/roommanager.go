/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/6/8
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core

import (
	"aliens/aliensbot/common/util"
	"aliens/aliensbot/exception"
	"aliens/testserver/module/room/conf"
	"aliens/testserver/module/room/config"
	"aliens/testserver/module/room/game"
	"aliens/testserver/protocol"
	"reflect"
)

var RoomManager = &roomManager{
	gameFactories: make(map[reflect.Type]game.Factory),
	rooms: make(map[string]*Room),
	players: make(map[int64]string),
}

func init() {
	//
	RoomManager.RegisterGameFactory(&game.CommonGameFactory{})
}

type roomManager struct {
	gameFactories map[reflect.Type]game.Factory //游戏工厂类

	rooms map[string]*Room //运行的游戏  游戏id - 房间对象

	players map[int64]string //所有玩家的对应信息 玩家id - 房间id

}

func (this *roomManager) RegisterGameFactory(factory game.Factory) {
	this.gameFactories[reflect.TypeOf(factory)] = factory
}

//获取玩家在哪个房间
func (this *roomManager) GetRoomByPlayerID(playerID int64) *Room {
	roomID := this.players[playerID]
	if roomID == "" {
		exception.GameException(protocol.Code_roomNotFound)
	}
	return this.EnsureRoom(roomID)
}

//获取房间
func (this *roomManager) EnsureRoom(roomID string) *Room {
	game := this.rooms[roomID]
	if game == nil {
		exception.GameException(protocol.Code_roomNotFound)
	}
	return game
}

//新建房间
func (this *roomManager) newRoom(config *config.RoomConfig) *Room {
	result := &Room{
		id:     util.GenUUID(),
		config: config,
		Seats:  NewSeats(config.MaxSeat),
	}

	for _, factory := range this.gameFactories {
		if factory.Match(config.AppID) {
			result.game = factory.NewGame(result)
			break
		}
	}

	if result.game == nil {
		exception.GameException(protocol.Code_gameNotFound)
	}
	return result
}

//玩家加入房间
func (this *roomManager) JoinRoom(appID string, roomID string, playerID int64) {
	room := this.EnsureRoom(roomID)
	room.AddPlayer(&protocol.Player{
		Playerid:playerID,
		Nickname:"蛇皮" + util.Int64ToString(playerID),
	})

}


//分配新房间
func (this *roomManager) AllocRoom(appID string, players []*protocol.Player) *Room {
	config := conf.GetRoomConfig(appID)
	room := this.newRoom(config)
	room.InitPlayers(players)
	this.rooms[room.GetID()] = room
	if players != nil {
		for _, player := range players {
			this.players[player.GetPlayerid()] = room.GetID()
		}
	}

	return room
}

//关闭房间
func (this *roomManager) RemoveRoom(roomID string) {
	room := this.EnsureRoom(roomID)
	room.Close()
	delete(this.rooms, roomID)
}
