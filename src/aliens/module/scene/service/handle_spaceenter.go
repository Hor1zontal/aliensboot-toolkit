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
	"aliens/module/scene/entity"
	"aliens/module/scene/util"
)

func handleSpaceEnter(request *scene.SpaceEnter, response *scene.SpaceEnterRet) {
	entity := mmorpg.SpaceManager.CreateEntity(request.GetSpaceID(), &entity.PlayerEntity{}, util.TransVector(request.GetPosition()), util.TransVector(request.GetDirection()))
	response.EntityID = entity.GetID()
}