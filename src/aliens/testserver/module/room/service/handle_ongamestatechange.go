// Code generated by aliensbot. DO NOT EDIT.
// source: room_interface.proto
package service

import (
	"aliens/testserver/module/room/core"
	"aliens/testserver/protocol"
)


//
func handleOnGameStateChange(authID int64, gateID string, request *protocol.OnGameStateChange, response *protocol.OnGameStateChangeRet) {
	response.Code = core.RoomManager.ChangeGameState(authID, request.GetState())


}


