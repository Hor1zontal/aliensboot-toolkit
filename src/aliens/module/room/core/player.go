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
	"aliens/protocol/framesync"
	"aliens/network"
)

type Player struct {
	seat uint8  //座位编号

	agent *network.UDPAgent //连接代理

	//seat agent.ID
	room *Room //当前加入的房间

	token string //权限token
}

//发动数据
func (this *Player) sendData(data []byte) {
	if this.agent != nil {
		this.agent.WriteData(data)
	}
}

func (this *Player) lostFrame(lostFrames []*framesync.Frame) {
	message := &framesync.Response{
		RequestLostFrameRet:&framesync.RequestLostFrameRet{
			Frame:lostFrames,
		},
	}
	data, _ := message.Marshal()
	this.agent.WriteData(data)
}

func (this *Player) auth(token string, agent *network.UDPAgent) bool {
	if this.token != token {
		return false
	}
	this.agent = agent
	return true
}
