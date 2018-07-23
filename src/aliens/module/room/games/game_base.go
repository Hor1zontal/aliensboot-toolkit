/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/6/13
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package games

import (
	"time"
	"aliens/protocol/room"
	"aliens/network"
)

//只需要处理消息转发即可
type SimpleGame struct {

	roomID string

	players map[uint32]*Player  //加入游戏的玩家

	maxPlayer int //最大人数

	startTimestamp time.Time

	createTimeStamp time.Time

	allocPlayerID uint32

	//token string //
}

func (this *SimpleGame) Init(roomID string) {
	this.roomID = roomID
	this.createTimeStamp = time.Now()
	//this.token = util.GenUUID()
	this.players = make(map[uint32]*Player)
}

func (this *SimpleGame) Start() {
	this.startTimestamp = time.Now()
}

func (this *SimpleGame) IsStarted() bool {
	return this.startTimestamp.Unix() > 0
}

func (this *SimpleGame) Stop() {

}

//玩家数量是否达到最大
func (this *SimpleGame) IsMaxPlayer() bool {
	return this.maxPlayer == len(this.players)
}

func (this *SimpleGame) BroadcastWithout(response []byte, playerID uint32) {
	for _, player := range this.players {
		if player.id != playerID {
			player.sendData(response)
			//dispatch.MQ.GatePush(conf.Config.Service.Name, player.id, response)
		}
	}
}

//广播所有消息
func (this *SimpleGame) BroadcastAll(response []byte) {
	for _, player := range this.players {
		player.sendData(response)
	}
}

//func (this *SimpleGame) BroadcastMessageAll(response *room.Response) {
//	for _, player := range this.players {
//		player.sendData(response)
//	}
//}

//func (this *SimpleGame) Send(response []byte, playerID uint32) {
//	player := this.players[playerID]
//	if player != nil {
//		player.sendData(response)
//	}
//}

func (this *SimpleGame) HandleMessage(message GameMessage) bool {
	request := message.request
	response := message.response
	switch request.(type) {
	case *room.LeaveRoom:
		this.handleLeaveRoom(request.(*room.LeaveRoom), response.(*room.LeaveRoomRet))
		return true
	case *room.JoinRoom:
		this.handleJoinRoom(request.(*room.JoinRoom), response.(*room.JoinRoomRet), message.agent)
		return true
	case *room.AllocFreeRoomSeat:
		this.handleAllocFreeSeat(request.(*room.AllocFreeRoomSeat), response.(*room.AllocFreeRoomSeatRet))
		return true
	}
	return false
}

func (this *SimpleGame) handleJoinRoom(request *room.JoinRoom, response *room.JoinRoomRet, agent network.Agent) {
	player := this.players[request.GetPlayerID()]
	if player == nil {
		response.Result = room.RoomResult_invalidToken
		return
	}
	player.agent = agent
	player.data = request.GetData()
	response.Result = room.RoomResult_success
}

func (this *SimpleGame) handleLeaveRoom(request *room.LeaveRoom, response *room.LeaveRoomRet) {
	_, ok := this.players[request.GetPlayerID()]
	if ok {
		delete(this.players, request.GetPlayerID())
	}

	//this.BroadcastAll()
	//this.BroadcastAll()
	response.Result = room.RoomResult_success
}

func (this *SimpleGame) handleAllocFreeSeat(request *room.AllocFreeRoomSeat, response *room.AllocFreeRoomSeatRet) {
	if this.IsMaxPlayer() {
		response.Result = room.RoomResult_maxPlayers
		return
	}
	this.allocPlayerID ++
	player := &Player{id:this.allocPlayerID}
	this.players[player.id] = player

	response.Result = room.RoomResult_success
	response.PlayerID = player.id
}


func buildPlayerLeavePush(playerID uint32) *room.Response {
	return &room.Response{
		PlayerLeavePush:&room.PlayerLeavePush{

		},
	}
}

func buildPlayerJoinPush(playerID uint32) {

}