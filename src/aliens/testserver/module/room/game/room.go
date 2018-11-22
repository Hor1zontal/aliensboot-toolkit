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
	"aliens/aliensbot/common/util"
	"aliens/aliensbot/exception"
	"aliens/testserver/module/room/config"
	"aliens/testserver/protocol"
	"github.com/gogo/protobuf/proto"
	"github.com/xiaonanln/goworld/engine/uuid"
)

//新建游戏
func NewRoom(appID string) *Room {
	//TODO 加载游戏配置
	var config *config.RoomConfig = &config.RoomConfig{AppID:appID, MaxSeat:2}

	result := &Room{
		id:uuid.GenUUID(),
		appID:appID,
		RoomConfig: config,
		players: make(map[int64]*Player),
	}
	return result
}

type Room struct {

	id string //房间id

	appID string //房间所属的游戏id

	*config.RoomConfig  //房间配置

	game *Game //房间内进行的游戏对象

	players map[int64]*Player //加入的玩家
}

func (room *Room) GetID() string {
	return room.id
}


//新增玩家
//func (room *Room) AddPlayer(player *Player) {
//	room.players[player.GetPlayerid()] = player
//	if room.IsMaxPlayer() {
//		//通知所有玩家初始化成功
//		push := &protocol.Push{Room: &protocol.Push_RoomCreatedRet{RoomCreatedRet: &protocol.RoomCreatedRet{
//			Players:room.GetAllPlayerData(),
//		}}}
//
//		room.BroadcastOtherPlayer(-1, push)
//	}
//}


//关闭房间
func (room *Room) Close() {
	if room.game != nil {
		room.game.Stop()
	}
}

//一次性添加房间人员
func (room *Room) InitPlayers(players []*protocol.Player) {
	if players == nil {
		return
	}
	//初始化玩家
	for index, player := range players {
		//座位号递增
		player.Seat = int32(index + 1)
		player.GroupId = util.Int32ToString(player.Seat)
		room.players[player.GetPlayerid()] = &Player{Player:player}
	}
}

func (room *Room) GetAllPlayer() map[int64]*Player {
	return room.players
}

//获取所有玩家数据
func (room *Room) GetAllPlayerData() []*protocol.Player {
	results := make([]*protocol.Player, len(room.players))
	index := 0
	for _, player := range room.players {
		results[index] = player.Player
		index++
	}
	return results
}

func (room *Room) GetPlayerData(playerID int64) *protocol.Player {
	player := room.EnsurePlayer(playerID)
	return player.Player
}

func (room *Room) EnsurePlayer(playerid int64) *Player {
	player := room.players[playerid]
	if player == nil {
		exception.GameException(protocol.Code_playerNotFound)
	}
	return player
}

//玩家是否全部加入
func (room *Room) IsMaxPlayer() bool {
	return len(room.players) == room.MaxSeat
}

//玩家准备
func (room *Room) PlayerReady(playerID int64) {
	if room.game != nil {
		exception.GameException(protocol.Code_gameAlreadyStart)
	}
	player := room.EnsurePlayer(playerID)
	player.Ready()
	//所有玩家准备完毕、即可开始游戏
	if room.IsAllPlayerReady() {
		//启动新游戏
		room.game = NewGame(room)
		room.game.Start()
	}
}

//是否所有玩家准备完毕
func (room *Room) IsAllPlayerReady() bool {
	for _, player := range room.players {
		if !player.IsReady() {
			return false
		}
	}
	return true
}


//玩家上报游戏结果
//玩家上报结果
func (room *Room) UploadResult(playerID int64, reports []*protocol.PlayerResult) {
	//TODO 处理玩家的上报结果


}

//游戏是否开始
func (room *Room) IsGameStart() bool {
	return room.game != nil
}

func (room *Room) EnsureGame() *Game {
	if room.game == nil {
		exception.GameException(protocol.Code_gameNotFound)
	}
	return room.game
}

//广播其他玩家
func (room *Room) BroadcastOtherPlayer(playerID int64, message proto.Message) {
	sendData, _ := proto.Marshal(message)
	for _, player := range room.players {
		if player.GetPlayerid() != playerID {
			player.SendMsg(sendData)
		}
	}
}

