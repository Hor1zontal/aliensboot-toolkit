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
	"aliens/aliensbot/log"
	"aliens/aliensbot/mmo"
	"aliens/aliensbot/mmo/core"
	"time"
)

const (
	TypeMonster mmo.EntityType = "Monster"
)


//
type Monster struct {
	mmo.Entity   // Entity type should always inherit entity.Entity

	movingToTarget  *mmo.Entity
	attackingTarget *mmo.Entity

	lastTickTime    time.Time
	attackCD        time.Duration
	lastAttackTime  time.Time
}

func (monster *Monster) DescribeEntityType(desc *core.EntityDesc) {
	desc.SetUseAOI(true, 100)
	desc.DefineAttr("lv", core.AttrAllClient| core.AttrPersist)
	desc.DefineAttr("hp", core.AttrAllClient| core.AttrPersist)
	desc.DefineAttr("maxHp", core.AttrAllClient| core.AttrPersist)
	desc.DefineAttr("action", core.AttrAllClient| core.AttrPersist)
}


func (monster *Monster) TestCall(param1 string, param2 int32, param3 []string) {
	log.Debugf("test call : %v %v %v", param1, param2, param3)
}