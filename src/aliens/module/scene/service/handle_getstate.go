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
	"aliens/mmorpg"
	"aliens/module/scene/util"
)


//
func handleGetState(authID int64, request *protocol.GetState, response *protocol.GetStateRet) {
	neighbors := mmorpg.SpaceManager.GetEntityState(request.GetSpaceID(), request.GetEntityID())
	response.Neighbors = util.BuildEntities(neighbors)
}
