/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/6/8
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core

import (
	"aliens/protocol/framesync"
	"sync"
	"aliens/network"
)

var Manager = &manager{rooms:make(map[string]*Room)}

var authResponseData []byte = nil

func init () {
	response := &framesync.Response{AuthRet:true}
	data, _ := response.Marshal()
	authResponseData = data
}

type manager struct {
	sync.RWMutex
	rooms map[string]*Room //运行的游戏 游戏类型 - 游戏id - 游戏对象

	authUDPAgent map[string]*network.UDPAgent  //验权通过的udp agent




}

func (this *manager) CreateRoom(maxSeat uint8) string {
	room := &Room{}
	room.init(maxSeat)
	this.Lock()
	defer this.Unlock()
	this.rooms[room.id] = room
	return room.id
}


//接收房间消息
func (this *manager) AcceptRoomMessage(roomID string, request interface{}, response interface{}) {
	this.Lock()
	room := this.rooms[roomID]
	this.Unlock()
	if room != nil {
		room.acceptMessage(request, response)
	}
}


func (this *manager) CloseRoom(roomID string) {
	room := this.rooms[roomID]
	if room == nil {
		return
	}
	room.close()
	delete(this.rooms, roomID)
}

func (this *manager) AcceptFrameMessage(message *framesync.Request, agent *network.UDPAgent) {
	authMessage := message.GetAuth()
	if authMessage != nil {
		room := this.rooms[authMessage.GetRoomID()]
		if room != nil {
			room.acceptFrameMessage(message, agent)
		}
	} else {
		player, ok := agent.UserData().(*Player)
		if ok {
			if player.room != nil {
				player.room.acceptFrameMessage(message, agent)
			}
		}
	}
}
