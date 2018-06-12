/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/5/25
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core

import (
	"aliens/module/room/frame"
	"aliens/protocol/framesync"
	"aliens/log"
	"aliens/network"
	"aliens/common/util"
)

type Room struct {
	id string

	token string //令牌

	frameChannel chan *RoomMessage //接受到帧消息管道

	channel chan []interface{} //接收rpc消息

	players map[uint8]*Player  //加入游戏的玩家

	frameManager *frame.Manager //帧管理容器

	maxSeat uint8 //最大人数

	allocSeat uint8 //当前分配到的座位编号

}

func (this *Room) init(maxSeat uint8) {
	this.frameChannel = make(chan *RoomMessage, 5)
	this.channel = make(chan []interface{})

	this.players = make(map[uint8]*Player)
	this.frameManager = frame.NewFrameManager(20)
	this.id = util.GenUUID()
	this.allocSeat = 0
}

func (this *Room) close() {
	close(this.frameChannel)
	this.frameChannel = nil
}

func (this *Room) start() {
	timer := this.frameManager.Start()
	go func() {
		for {
			select {
			case message := <-this.frameChannel:
				//处理udp消息
				request := message.request
				agent := message.agent
				player, ok := agent.UserData().(*Player)

				if ok {
					if request.GetCommand() != nil {
						this.frameManager.AcceptCommand(request.GetCommand())
					} else if request.GetRequestLostFrame() != nil {
						lostFrames := this.frameManager.GetFrames(request.GetRequestLostFrame().GetSeq())
						player.lostFrame(lostFrames)
					}
				} else {
					if request.GetAuth() != nil {
						authMessage := request.GetAuth()
						this.auth(uint8(authMessage.GetSeatID()), authMessage.GetToken(), agent)
					}
				}

				case <-timer.C:
					//处理定时消息
					frame := this.frameManager.NextFrame()
					if frame != nil {
						this.syncFrame(frame)
					}

				case rw := <-this.channel:
					//处理rpc调用消息
					this.handle(rw[0], rw[1])
			}


		}

	}()
}

func (this *Room) syncFrame(frame *framesync.Frame) {
	data, _ := frame.Marshal()
	for _, player := range this.players {
		player.sendData(data)
	}
}

func (this *Room) acceptFrameMessage(message *framesync.Request, agent *network.UDPAgent) {
	if this.frameChannel != nil {
		this.frameChannel <- &RoomMessage{agent,message,}
	}
}

func (this *Room) acceptMessage(request interface{}, response interface{}) {
	if this.channel != nil {
		this.channel <- []interface{}{request, response}
	}
}

//验证udp网络连接权限
func (this *Room) auth(seatID uint8, token string, agent *network.UDPAgent) {
	player := this.players[seatID]
	if player == nil {
		log.Warnf("invalid auth room : %v  token : %v", this.id, token)
		return
	}
	if player.auth(token, agent) {
		agent.SetUserData(player)
		agent.WriteData(authResponseData)
	}
}


func (this *Room) LeaveRoom(seatID uint8) {
	delete(this.players, seatID)
}


