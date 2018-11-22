/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/6/8
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package manager

//var GameManager = &manager{games: make(map[string]*game.Game), players: make(map[int64]string)}
//
//type manager struct {
//
//	games map[string]*game.Game //运行的游戏  游戏id - 游戏对象
//
//	players map[int64]string //玩家id - 房间id
//
//	//authUDPAgent map[uint32]*network.UDPAgent //验权通过的udp agent
//}
//
//func (this *manager) GetGame(playerID int64) *game.Game {
//	gameID := this.players[playerID]
//	if gameID == "" {
//		exception.GameException(protocol.Code_gameNotFound)
//	}
//	return this.EnsureGame(gameID)
//}
//
////初始化游戏
//func (this *manager) AllocGame(appID string, players []*protocol.Player) *game.Game {
//	game := game.NewGame(appID, players)
//
//	this.games[game.GetID()] = game
//	for _, player := range players {
//		this.players[player.GetPlayerid()] = game.GetID()
//	}
//	return game
//}
//
////初始化游戏
//func (this *manager) RemoveGame(gameID string) {
//	game := this.EnsureGame(gameID)
//	players := game.GetAllPlayer()
//	for _, player := range players {
//		delete(this.players, player.GetPlayerid())
//	}
//	game.Reset()
//}
//
//
//func (this *manager) EnsureGame(gameID string) *game.Game{
//	game := this.games[gameID]
//	if game == nil {
//		exception.GameException(protocol.Code_gameNotFound)
//	}
//	return game
//}

//func (this *manager) CloseRoom(roomID string) {
//	room := this.rooms[roomID]
//	if room == nil {
//		return
//	}
//	room.close()
//	delete(this.rooms, roomID)
//}

//接收房间消息
//func (this *manager) AcceptRoomMessage(roomID string, request interface{}, response interface{}, agent network.Agent) {
//	this.Lock()
//	room := this.games[roomID]
//	this.Unlock()
//	if room != nil {
//		room.AcceptMessage(agent, request, response)
//	}
//}
