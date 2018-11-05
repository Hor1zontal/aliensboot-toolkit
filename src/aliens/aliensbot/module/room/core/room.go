/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/5/25
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core

import (
	"aliens/aliensbot/common/util"
	"aliens/testserver/module/room/games"
	"aliens/aliensbot/network"
)

type Room struct {
	id string //房间id

	game games.Game //绑定游戏

	channel chan games.GameMessage //房间消息处理管道

	//tokens map[string]int64  //token-分配的时间戳
	//allocSeat uint8 //当前分配到的座位编号
}

func (this *Room) init(gameID uint32) {

	this.id = util.GenUUID()

	this.game = games.NewGame(gameID)

	this.channel = make(chan games.GameMessage)

	//this.allocSeat = 0

	//this.maxSeat = game.MaxPlayer()

	this.game.Init(this.id)

}

func (this *Room) close() {
	this.game.Stop()
}

func (this *Room) startGame() {
	this.game.Start()
	timerGame, ok := this.game.(games.TimeGame)
	if ok {
		go this.gameTimerLogic(timerGame)
	} else {
		go this.gameLogic()
	}
}

func (this *Room) IsFull() bool {
	return this.game.IsMaxPlayer()
}

func (this *Room) gameTimerLogic(game games.TimeGame) {
	timer := game.GetTimer()
	for {
		select {
		case <-timer.C:
			game.HandleTimer()
		case message := <-this.channel:
			game.HandleMessage(message)
		}
	}
}

func (this *Room) gameLogic() {
	for {
		message := <-this.channel
		this.game.HandleMessage(message)
	}
}

func (this *Room) AcceptMessage(agent network.Agent, request interface{}, response interface{}) {
	//player, _ := agent.UserData().(*games.Player)
	if this.channel != nil {
		this.channel <- games.NewGameMessage(agent, request, response)
	}
}

//验证udp网络连接权限
//func (this *Room) auth(seatID uint8, token string, agent *network.UDPAgent) {
//	player := this.players[seatID]
//	if player == nil {
//		log.Warnf("invalid auth room : %v  token : %v", this.id, token)
//		return
//	}
//	if player.auth(token, agent) {
//		agent.SetUserData(player)
//		agent.WriteData(authResponseData)
//	}
//}
