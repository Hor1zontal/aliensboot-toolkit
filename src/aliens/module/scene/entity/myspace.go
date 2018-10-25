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
	"aliens/mmorpg/core"
	"aliens/mmorpg/config"
)

type MySpace struct {
	*core.Space
}

func (this *MySpace) GetConfig() config.SpaceConfig {
	return config.SpaceConfig{"map1",-400, 400, -300, 300, 200}
}

func (this *MySpace) OnEntityEnter(entity *core.Entity) {

}

func (this *MySpace) OnEntityLeave(entity *core.Entity) {

}

func (this *MySpace) OnEntityMove(entity *core.Entity) {
	//dispatch.gate
}



