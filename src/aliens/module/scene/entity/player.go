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

import "aliens/mmorpg/entity"

//玩家实体
type PlayerEntity struct {

}


//可视距离
func (this * PlayerEntity) GetDist() float32 {
	return 100
}


func (this * PlayerEntity) OnEntityEnter(entity *entity.Entity) {

}


func (this * PlayerEntity) OnEntityLeave(entity *entity.Entity) {

}


func (this * PlayerEntity) OnEntityMove(entity *entity.Entity) {


}