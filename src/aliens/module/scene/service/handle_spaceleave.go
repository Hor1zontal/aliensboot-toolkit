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
	"aliens/protocol/scene"
	"aliens/mmorpg"
)


//
func handleSpaceLeave(request *scene.SpaceLeave, response *scene.SpaceLeaveRet) {
	mmorpg.SpaceManager.LeaveEntity(request.GetSpaceID(), request.GetEntityID())
}
