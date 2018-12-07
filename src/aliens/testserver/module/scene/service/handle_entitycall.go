// Code generated by aliensbot. DO NOT EDIT.
// source: scene_interface.proto
package service

import (
	"aliens/aliensbot/exception"
	"aliens/aliensbot/log"
	"aliens/aliensbot/mmo"
	"aliens/aliensbot/mmo/core"
	entity2 "aliens/testserver/module/scene/entity"
	"aliens/testserver/protocol"
)

//
func handleEntityCall(authID int64, gateID string, request *protocol.EntityCall) {
	entity, err := mmo.HandlerRemoteEntityCall(entity2.GetPlayerID(authID), core.EntityID(request.GetEntityID()), request.GetMethod(), request.GetArgs())
	if entity == nil {
		exception.GameException(protocol.Code_entityNotFound)
	}
	if err != nil {
		log.Error(err)
		exception.GameException(protocol.Code_invalidEntityCall)
	}
}