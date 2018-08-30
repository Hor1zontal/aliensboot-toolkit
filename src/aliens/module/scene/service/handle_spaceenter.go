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
	"aliens/module/scene/entity"
	"aliens/mmorpg/core"
)

func handleSpaceEnter(authID int64, request *protocol.SpaceEnter, response *protocol.SpaceEnterRet) {
	player := entity.NewPlayerEntity(authID)
	entity, err := core.SpaceManager.CreateEntity(core.SpaceID(request.GetSpaceID()), player, request.GetPosition(), request.GetDirection())
	if entity != nil && err != nil {
		response.EntityID = int32(entity.GetID())
	}
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
