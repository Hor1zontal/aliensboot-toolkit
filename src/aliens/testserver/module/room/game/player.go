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
	"aliens/testserver/dispatch/rpc"
	"aliens/testserver/protocol"
)

type Player struct {

	*protocol.Player

	gateID  string //是否绑定了网关、玩家是否连接到服务器

	ready 	bool  //玩家是否准备完毕

}

func (player *Player) SendMsg(data []byte) {
	pushMessage := &protocol.PushMessage{
		AuthID: player.GetPlayerid(),
		Data:   data,
		Service: "room",
	}
	rpc.Gate.PushMessage(player.gateID, pushMessage)
}

func (player *Player) Ready() {
	player.ready = true
}

func (player *Player) IsReady() bool {
	return player.ready
}