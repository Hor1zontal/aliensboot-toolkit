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
	"aliens/testserver/protocol"
)

//新建游戏
func NewGame(room *Room) *Game {
	return &Game{room}
}

type Game struct {
	*Room  //游戏所处的房间
}

//开始游戏
func (game *Game) Start() {
	//通知所有玩家游戏开始
	push := &protocol.Response{Room: &protocol.Response_GameStartRet{GameStartRet: &protocol.GameStartRet{}}}
	game.BroadcastOtherPlayer(-1, push)
}

//结束游戏
func (game *Game) Stop() {
	push := &protocol.Response{Room: &protocol.Response_GameResetRet{GameResetRet: &protocol.GameResetRet{}}}
	game.BroadcastOtherPlayer(-1, push)
}



//接收玩家数据，同步给其他玩家
func (game *Game) AcceptPlayerData(playerID int64, data string) {
	push := &protocol.Response{Room: &protocol.Response_GameDataRet{
		GameDataRet: &protocol.GameDataRet{
			Data: data,
		},
	},
	}
	game.BroadcastOtherPlayer(playerID, push)
}
