/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/11/13
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core

import (
	"aliens/aliensbot/exception"
	"aliens/testserver/module/room/config"
	"aliens/testserver/module/room/game"
	"aliens/testserver/protocol"
	"github.com/gogo/protobuf/proto"
)

type Seat int32

type Room struct {

	id string //房间id

	Seats //桌子,数组下标为座位编号

	config *config.RoomConfig //房间配置

	game game.Game //房间内进行的游戏对象

}

func (room *Room) GetID() string {
	return room.id
}

//新增玩家
func (room *Room) AddPlayer(player *protocol.Player) {
	ok := room.Add(&Player{Player:player})
	if !ok {
		exception.GameException(protocol.Code_roomMaxPlayer)
	}
	if room.IsFull() {
		//通知所有玩家初始化成功
		push := &protocol.Response{Room: &protocol.Response_PlayerJoinRet{PlayerJoinRet: &protocol.PlayerJoinRet{
			RoomID:room.GetID(),
			Player:player,
		}}}
		room.BroadcastOtherPlayer(player.GetPlayerid(), push)
	}
}

//一次性添加房间人员
func (room *Room) InitPlayers(players []*protocol.Player) {
	if players == nil {
		return
	}
	//初始化玩家
	for _, player := range players {
		room.AddPlayer(player)
	}
}

//关闭房间
func (room *Room) Close() {
	room.Clean()
	//for _, player := range players {
	//	delete(this.players, player.GetPlayerid())
	//}
	if room.game != nil {
		room.game.Stop()
	}
}

//获取所有玩家数据
func (room *Room) GetAllPlayerData() []*protocol.Player {
	results := []*protocol.Player{}
	room.Foreach(func (player *Player) {
		results = append(results, player.Player)
	})
	return results
}

func (room *Room) GetPlayerData(playerID int64) *protocol.Player {
	player := room.EnsurePlayer(playerID)
	return player.Player
}

func (room *Room) EnsurePlayer(playerID int64) *Player {
	player := room.Get(playerID)
	if player == nil {
		exception.GameException(protocol.Code_playerNotFound)
	}
	return player
}

//玩家准备
func (room *Room) PlayerReady(playerID int64) {
	if room.IsGameStart() {
		exception.GameException(protocol.Code_gameAlreadyStart)
	}
	player := room.EnsurePlayer(playerID)
	player.Ready()

	//所有玩家准备完毕、即可开始游戏
	if room.IsAllReady() {
		//启动新游戏
		room.game.Start()
	}
}

//玩家上报游戏结果
//玩家上报结果
func (room *Room) UploadResult(playerID int64, reports []*protocol.PlayerResult) {
	//TODO 处理玩家的上报结果
	//game := room.EnsureGame()

}

//游戏是否开始
func (room *Room) IsGameStart() bool {
	return room.game != nil && room.game.IsStart()
}

func (room *Room) EnsureGame() game.Game {
	if room.game == nil {
		exception.GameException(protocol.Code_gameNotFound)
	}
	return room.game
}

//广播其他玩家
func (room *Room) BroadcastOtherPlayer(playerID int64, message proto.Message) {
	sendData, _ := proto.Marshal(message)
	room.Foreach(func(player *Player) {
		if player.GetPlayerid() != playerID {
			player.SendMsg(sendData)
		}
	})
}

//接收玩家数据，同步给其他玩家
func (room *Room) AcceptPlayerData(playerID int64, data string) {
	room.game.AcceptPlayerData(playerID, data)
}
