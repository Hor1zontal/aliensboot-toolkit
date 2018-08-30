/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/3/30
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package service

import (
	"aliens/protocol"

	"aliens/module/scene/util"
	"aliens/mmorpg/core"
)


//
func handleGetState(authID int64, request *protocol.GetState, response *protocol.GetStateRet) {
	neighbors, err := core.SpaceManager.GetEntityState(core.EntityID(authID))
	if err != nil {
		//玩家不在任何场景内
	}
	response.Neighbors = util.BuildEntities(neighbors)
}
