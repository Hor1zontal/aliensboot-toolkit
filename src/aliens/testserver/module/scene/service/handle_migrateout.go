// Code generated by aliensbot. DO NOT EDIT.
// source: scene_interface.proto
package service

import (
	"aliens/aliensbot/log"
	"aliens/aliensbot/mmo"
	"aliens/testserver/protocol"
)




//
func handleMigrateOut(authID int64, gateID string, request *protocol.MigrateOut) {
	err := mmo.MigrateOut(mmo.EntityID(request.GetToSpaceID()), mmo.EntityID(request.GetEntityID()))
	if err != nil {
		log.Errorf("handle migrateOut error : %v", err)
	}
}