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
	"aliens/module/scene/entity"
	"aliens/module/scene/util"
)

func handleSpaceEnter(authID int64, request *protocol.SpaceEnter, response *protocol.SpaceEnterRet) {
	entity := mmorpg.SpaceManager.CreateEntity(request.GetSpaceID(), &entity.PlayerEntity{}, util.TransVector(request.GetPosition()), util.TransVector(request.GetDirection()))
	response.EntityID = entity.GetID()

	//GatePush(serviceType string, clientID string, message proto.Message) error {
	//dispatch.MQ.GatePush(constant.SERVICE_SCENE, "1_1", &scene.SceneResponse{
	//	SpacePush:
	//	Response:&scene.SceneResponse_ScenePush{
	//		ScenePush:"测试",
	//	},
	//})
	//a := &scene.SpacePush{{
	//	SpacePush:&scene.SpacePush{
	//
	//	},
	//}
}
