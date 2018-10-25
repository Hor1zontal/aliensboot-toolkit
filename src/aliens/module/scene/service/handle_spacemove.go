/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
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
	"aliens/mmorpg/core"
)


//ignore d s
func handleSpaceMove(authID int64, gateID string, request *protocol.SpaceMove, response *protocol.SpaceMoveRet) {
	core.SpaceManager.MoveEntity(core.EntityID(authID), request.GetPosition(), request.GetDirection())
}
