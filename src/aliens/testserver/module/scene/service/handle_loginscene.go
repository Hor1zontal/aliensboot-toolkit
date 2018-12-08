// Code generated by aliensbot. DO NOT EDIT.
// source: scene_interface.proto
package service

import (
	"aliens/aliensbot/cluster/center"
	"aliens/aliensbot/exception"
	"aliens/aliensbot/log"
	"aliens/aliensbot/mmo"
	"aliens/aliensbot/mmo/unit"
	"aliens/testserver/dispatch/rpc"
	"aliens/testserver/module/scene/cache"
	"aliens/testserver/module/scene/entity"
	"aliens/testserver/protocol"
)


//
func handleLoginScene(authID int64, gateID string, request *protocol.LoginScene) {
	//获取空间所在的服务器节点
	node, err := cache.Manager.GetSpaceNode(request.GetSpaceID())
	if err != nil {
		exception.GameException1(protocol.Code_DBExcetpion, err)
	}
	if node == "" {
		exception.GameException(protocol.Code_entityNotFound)
	}

	authID = request.GetAuthID()
	gateID = request.GetGateID()

	//是当前服务器节点
	if node == center.ClusterCenter.GetNodeID() {
		playerID := entity.GetPlayerID(authID)

		entity, err := mmo.EnterSpace(
			mmo.EntityID(request.GetSpaceID()),
			entity.TypePlayer,
			playerID,
			unit.Vector{X:0,Y:0,Z:0})
		if err != nil {
			log.Errorf("enter space error : %v", err)
			return
		}
		entity.OnCallFromLocal("Login", []interface{}{authID, gateID})
	} else {
		//登录请求转发到对应节点
		rpc.Scene.LoginScene(node, request)
	}

}


