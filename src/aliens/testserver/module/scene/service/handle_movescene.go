// Code generated by aliensbot. DO NOT EDIT.
// source: scene_interface.proto
package service

import (
	"aliens/aliensbot/exception"
	"aliens/aliensbot/mmo"
	"aliens/testserver/module/scene/entity"
	"aliens/testserver/protocol"
)




//
func handleMoveScene(authID int64, gateID string, request *protocol.MoveScene) {
	//获取空间所在的服务器节点
	err := mmo.MigrateTo(mmo.EntityID(request.GetSpaceID()), entity.GetPlayerID(authID))
	if err != nil {
		exception.GameException1(protocol.Code_InternalException, err)
	}
}
