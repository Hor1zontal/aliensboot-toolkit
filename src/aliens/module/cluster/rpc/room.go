/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/6/15
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package rpc

import (
	"aliens/protocol/room"
	"aliens/module/cluster/dispatch"
	"aliens/protocol"
	"aliens/exception"
	"aliens/log"
)

var Proxy_room = &roomRPCHandle{"room"}


type roomRPCHandle struct {
	name string
}


func (this *roomRPCHandle) request(request *room.Request) *room.Response {
	rpcRet, err := dispatch.RPC.SyncRequest(this.name, request)
	if err != nil {
		log.Error(err)
		exception.GameException(exception.INVALID_SERVICE)
	}
	any, ok := rpcRet.(*protocol.Any)
	if !ok {
		log.Error("invalid rpc ret data")
		exception.GameException(exception.INVALID_SERVICE)
	}
	messageRet := &room.Response{}
	messageRet.Unmarshal(any.GetValue())
	return  messageRet
}


func (this *roomRPCHandle) AllocFreeRoomSeat(request *room.AllocFreeRoomSeat) *room.AllocFreeRoomSeatRet {
	message := &room.Request{
		Request:&room.Request_AllocFreeRoomSeat{
			AllocFreeRoomSeat:request,
		},
	}
	messageRet := this.request(message)
	return messageRet.GetAllocFreeRoomSeatRet()
}

func (this *roomRPCHandle) JoinRoom(request *room.JoinRoom) *room.JoinRoomRet {
	message := &room.Request{
		Request:&room.Request_JoinRoom{
			JoinRoom:request,
		},
	}
	messageRet := this.request(message)
	return messageRet.GetJoinRoomRet()
}

func (this *roomRPCHandle) LeaveRoom(request *room.LeaveRoom) *room.LeaveRoomRet {
	message := &room.Request{
		Request:&room.Request_LeaveRoom{
			LeaveRoom:request,
		},
	}
	messageRet := this.request(message)
	return messageRet.GetLeaveRoomRet()
}
