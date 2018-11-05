/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/5/25
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package games

import (
	"aliens/aliensbot/network"
)

type Player struct {
	id uint32 //玩家id

	agent network.Agent //连接代理 是否连接

	game Game

	//token string //权限token

	data string //玩家数据
}

func (this *Player) getID(data []byte) {
	if this.agent != nil {
		this.agent.WriteMsg(data)
	}
}

//发动数据
func (this *Player) sendData(data []byte) {
	if this.agent != nil {
		this.agent.WriteMsg(data)
	}
}

//func (this *Player) auth(token string, agent gate.Agent) bool {
//	if this.token != token {
//		return false
//	}
//	this.agent = agent
//	return true
//}
