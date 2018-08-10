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


//ignore d s
func handleSpaceMove(authID int64, request *protocol.SpaceMove, response *protocol.SpaceMoveRet) {
	mmorpg.SpaceManager.MoveEntity(request.GetSpaceID(), request.GetEntityID(), util.TransVector(request.GetPosition()), util.TransVector(request.GetDirection()))
}
