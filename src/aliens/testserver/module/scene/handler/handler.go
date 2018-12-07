/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/12/4
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package handler

import (
	"aliens/aliensbot/mmo"
	"aliens/testserver/module/scene/cache"
	"aliens/testserver/module/scene/entity"
)



//初始化space
func Init() {
	mmo.RegisterSpace(&entity.GameSpace{})
	mmo.RegisterEntity(&entity.Monster{})
	mmo.RegisterEntity(&entity.Player{})

	space, _ := mmo.CreateSpace(entity.TypeGameSpace, "space1")
	cache.Manager.SetSpaceNode(string(space.GetID()))

	//entity := mmo.CreateEntity("Monster", space, unit.Vector{0,0,0})
	//log.Debugf("entity %v", entity.GetID())

	//mmo.Call(entity.GetID(), "TestCall", "1", 1, []string{"3"})

}
