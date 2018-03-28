/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/3/24
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package entity

import (
	"aliens/mmorpg/entity"
	"aliens/mmorpg"
)

type MySpace struct {
	id int32
}


func (this *MySpace) GetID() int32 {
	return this.id
}

func (this *MySpace) OnEntityEnter(entity *entity.Entity) {

}

func (this *MySpace) OnEntityLeave(entity *entity.Entity) {

}

func (this *MySpace) OnEntityMove(entity *entity.Entity) {

}

func Init() {
	mmorpg.SpaceManager.CreateSpace(&MySpace{id:1}, -400, 400, -300, 300, 200)
}

