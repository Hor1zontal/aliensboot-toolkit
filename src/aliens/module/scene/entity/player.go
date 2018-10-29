/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/3/24
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package entity

import (
	"aliens/log"
	"aliens/mmorpg/core"
)

//func NewPlayerEntity(authID int64) *PlayerEntity {
//	player := &PlayerEntity{
//		authID:authID,
//		entityID:core.EntityID(authID),
//	}
//	return player
//}

//玩家实体
type PlayerEntity struct {
	*core.Entity
}

//可视距离
func (this *PlayerEntity) GetDist() float32 {
	return 100
}

func (this *PlayerEntity) OnEntityEnter(entity *core.Entity) {
	log.Debugf("entity enter %v", entity.GetID())

	//dispatch.GatePush(this.clientID, )
}

func (this *PlayerEntity) OnEntityLeave(entity *core.Entity) {
	log.Debugf("entity leave %v", entity.GetID())

	//dispatch.GatePush(this.clientID, )
}

func (this *PlayerEntity) OnEntityMove(entity *core.Entity) {
	log.Debugf("entity move %v", entity.GetID())

}
