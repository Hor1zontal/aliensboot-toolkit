// Code generated by aliensbot. DO NOT EDIT.
// source: room_interface.proto
package service

import (
	"aliens/testserver/module/room/core"
	"aliens/testserver/protocol"
)


//
func handleRespondJoinGame(authID int64, gateID string, request *protocol.RespondJoinGame) {
	room := core.RoomManager.GetRoomByPlayerID(authID)
	room.JoinGame(authID)
}
