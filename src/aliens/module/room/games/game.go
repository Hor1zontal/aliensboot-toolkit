/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/6/13
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package games

import (
	"aliens/network"
	"time"
)

var GameRegistry = make(map[uint32]func() Game)

const (
	GameNoobFight uint32 = 1
	GameSimple    uint32 = 2
)

func init() {
	GameRegistry[GameNoobFight] = NewNoobFightGame
}

//创建游戏
func NewGame(id uint32) Game {
	factory := GameRegistry[id]
	if factory != nil {
		return factory()
	}
	return nil
}

func NewGameMessage(agent network.Agent, request interface{}, response interface{}) GameMessage {
	return GameMessage{agent, request, response}
}

type GameMessage struct {
	agent    network.Agent
	request  interface{}
	response interface{}
}

type Game interface {
	Init(roomID string) //初始化游戏

	Start() //开始游戏

	Stop() //停止游戏

	IsMaxPlayer() bool

	HandleMessage(message GameMessage) bool //处理游戏消息 seatID 消息的发送者

	BroadcastWithout(response []byte, playerID uint32)

	BroadcastAll(response []byte)
}

type TimeGame interface {
	HandleMessage(message GameMessage) bool //处理游戏消息 seatID 消息的发送者

	GetTimer() *time.Timer

	HandleTimer()
}
