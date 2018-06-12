/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/6/12
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core

import (
	"aliens/protocol/room"
	"aliens/log"
)

func (this *Room) handle(request interface{}, response interface{}) {
	log.Debugf("%v - %v", request, response)
	switch request.(type) {
		case *room.JoinRoom:
			this.handleJoinRoom(request.(*room.JoinRoom), response.(*room.JoinRoomRet))
			break
		case *room.LeaveRoom:
			this.handleLeaveRoom(request.(*room.LeaveRoom), response.(*room.LeaveRoomRet))
			break
	}

}

func (this *Room) handleJoinRoom(request *room.JoinRoom, response *room.JoinRoomRet) {
	if this.token != request.GetToken() {
		response.Result = room.RoomResult_invalidToken
		return
	}
	if this.allocSeat >= this.maxSeat {
		response.Result = room.RoomResult_maxPlayers
		return
	}

	this.allocSeat ++
	player := &Player{seat:this.allocSeat, room:this}
	this.players[player.seat] = player

	response.Result = room.RoomResult_success
	response.SeatID = uint32(player.seat)
}

func (this *Room) handleLeaveRoom(request *room.LeaveRoom, response *room.LeaveRoomRet) {

}


