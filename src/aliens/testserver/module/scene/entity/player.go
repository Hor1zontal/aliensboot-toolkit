/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/12/4
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package entity

import (
	"aliens/aliensbot/common/util"
	"aliens/aliensbot/log"
	"aliens/aliensbot/mmo"
	"aliens/aliensbot/mmo/core"
	"aliens/aliensbot/mmo/unit"
	"aliens/testserver/dispatch/rpc"
	"aliens/testserver/module/scene/conf"
	"aliens/testserver/module/scene/utils"
	"aliens/testserver/protocol"
	"time"
)

const (
	TypePlayer mmo.EntityType = "Player"

	playerAttrUid = "uid"
	playerAttrGateid = "gid"

)

func GetPlayerID(authID int64) mmo.EntityID {
	return mmo.EntityID("P_" + util.Int64ToString(authID))
}

//
type Player struct {
	mmo.Entity   // Entity type should always inherit entity.Entity

	syncTimerID mmo.EntityTimerID

}

func (player *Player) DescribeEntityType(desc *core.EntityDesc) {
	//视野范围
	desc.SetUseAOI(true, 500)
	desc.DefineAttr(playerAttrUid, core.AttrAllClient| core.AttrPersist) //用户id
	desc.DefineAttr(playerAttrGateid, core.AttrClient)	//网关id

	desc.DefineAttr("lv", core.AttrAllClient| core.AttrPersist)
	desc.DefineAttr("hp", core.AttrAllClient| core.AttrPersist)
	desc.DefineAttr("maxHp", core.AttrAllClient| core.AttrPersist)
	desc.DefineAttr("action", core.AttrAllClient| core.AttrPersist)

}


func (player *Player) Login(authID int64, gateID string) {
	player.Set(playerAttrUid, authID)
	player.Set(playerAttrGateid, gateID)

	syncMessage := &protocol.Response{
		Scene:&protocol.Response_LoginSceneRet{
			LoginSceneRet:&protocol.LoginSceneRet{
				Entity:utils.BuildEntity(player.Entity, true),
			},
		},
	}

	//玩家的消息绑定到当前服务器节点
	rpc.Gate.BindService1(gateID, authID, conf.GetServiceName())

	rpc.Gate.Push(conf.GetServiceName(), authID, gateID, syncMessage)

	//玩家每100ms同步一次数据
	player.syncTimerID = player.AddTimer(200 * time.Millisecond, "SyncData")
}


func (player *Player) Logout() {
	player.Destroy()
}


func (player *Player) Move_Client(x string, y string) {
	log.Debugf("%v move %v - %v", player.GetID(), x, y)


	player.SetPosition(unit.Vector{X:unit.Coord(util.StringToFloat32(x)), Y:unit.Coord(util.StringToFloat32(y)), Z:0})
}


////sync self 发送自己的玩家数据
func (player *Player) SyncData() {
	if !player.IsLogin() {
		//log.Warnf("player %v is not login", player.GetID())
		return
	}

	interest := player.GetInterest()
	var entities = make([]*protocol.Entity, len(interest))

	index := 0
	for entity, _ := range interest {
		entities[index] = utils.BuildEntity(*entity, entity.GetID() == player.GetID())
		index ++
	}

	syncMessage := &protocol.Response{
		Scene:&protocol.Response_EntityPush{
			EntityPush:&protocol.EntityPush{
				Neighbors:entities,
			},
		},
	}

	rpc.Gate.Push(conf.GetServiceName(), player.GetInt64(playerAttrUid), player.GetString(playerAttrGateid), syncMessage)
}



func (player *Player) IsLogin() bool {
	return player.GetInt64(playerAttrUid) > 0
}