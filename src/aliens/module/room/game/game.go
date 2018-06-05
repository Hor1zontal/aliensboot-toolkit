/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/5/25
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package game

import "aliens/module/room/framesync"


func NewGame() *Game {
	game := &Game{}
    game.frameManager = NewFrameManager(20, game.SyncFrame)
    return game
}

type Game struct {
	//frameMnewFrameManager()
	frameManager *FrameManager
	players map[uint8]*Player  //加入游戏的玩家

}

func (this *Game) AcceptMessage(frame interface{}, gate Player) {

}

func (this *Game) SyncFrame(frame *framesync.Frame) {
	message := &room.FrameSync{
		Frames:[]*framesync.Frame{frame},
	}
	data, _ := message.Marshal(message)
	for _, player := range this.players {
		if player.haveLostFrame() {

		} else {
			player.sendData(data)
		}
	}
}


func (this *Game) Over() {

}

func (this *Game) AddPlayer(player *Player) {

}
