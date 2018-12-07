// Code generated by aliensbot. DO NOT EDIT.
// source: module room
package rpc

import (
	"aliens/testserver/protocol"
)

var Room = &roomRPCHandler{&rpcHandler{name:"room"}}


type roomRPCHandler struct {
	*rpcHandler
}


func (this *roomRPCHandler) OnPlayerStateChange(node string, request *protocol.OnPlayerStateChange) *protocol.OnPlayerStateChangeRet {
	message := &protocol.Request{
		Room:&protocol.Request_OnPlayerStateChange{
			OnPlayerStateChange:request,
		},
	}
	messageRet := this.Request(node, message)
	return messageRet.GetOnPlayerStateChangeRet()
}

func (this *roomRPCHandler) JoinRoom(node string, request *protocol.JoinRoom) *protocol.JoinRoomRet {
	message := &protocol.Request{
		Room:&protocol.Request_JoinRoom{
			JoinRoom:request,
		},
	}
	messageRet := this.Request(node, message)
	return messageRet.GetJoinRoomRet()
}

func (this *roomRPCHandler) OnGameStateChange(node string, request *protocol.OnGameStateChange) *protocol.OnGameStateChangeRet {
	message := &protocol.Request{
		Room:&protocol.Request_OnGameStateChange{
			OnGameStateChange:request,
		},
	}
	messageRet := this.Request(node, message)
	return messageRet.GetOnGameStateChangeRet()
}

func (this *roomRPCHandler) ShowUser(node string, request *protocol.ShowUser) *protocol.ShowUserRet {
	message := &protocol.Request{
		Room:&protocol.Request_ShowUser{
			ShowUser:request,
		},
	}
	messageRet := this.Request(node, message)
	return messageRet.GetShowUserRet()
}

func (this *roomRPCHandler) GetRoomInfo(node string, request *protocol.GetRoomInfo) *protocol.GetRoomInfoRet {
	message := &protocol.Request{
		Room:&protocol.Request_GetRoomInfo{
			GetRoomInfo:request,
		},
	}
	messageRet := this.Request(node, message)
	return messageRet.GetGetRoomInfoRet()
}

func (this *roomRPCHandler) RoomCreate(node string, request *protocol.RoomCreate) *protocol.RoomCreateRet {
	message := &protocol.Request{
		Room:&protocol.Request_RoomCreate{
			RoomCreate:request,
		},
	}
	messageRet := this.Request(node, message)
	return messageRet.GetRoomCreateRet()
}

func (this *roomRPCHandler) GetBigoData(node string, request *protocol.GetBigoData) *protocol.GetBigoDataRet {
	message := &protocol.Request{
		Room:&protocol.Request_GetBigoData{
			GetBigoData:request,
		},
	}
	messageRet := this.Request(node, message)
	return messageRet.GetGetBigoDataRet()
}



func (this *roomRPCHandler) UpdateBigoData(node string, request *protocol.UpdateBigoData) error {
	message := &protocol.Request{
		Room:&protocol.Request_UpdateBigoData{
			UpdateBigoData:request,
		},
	}
	return this.Send(node, message)
}

func (this *roomRPCHandler) UploadGameResult(node string, request *protocol.UploadGameResult) error {
	message := &protocol.Request{
		Room:&protocol.Request_UploadGameResult{
			UploadGameResult:request,
		},
	}
	return this.Send(node, message)
}

func (this *roomRPCHandler) RequestJoinGame(node string, request *protocol.RequestJoinGame) error {
	message := &protocol.Request{
		Room:&protocol.Request_RequestJoinGame{
			RequestJoinGame:request,
		},
	}
	return this.Send(node, message)
}

func (this *roomRPCHandler) FrameData(node string, request *protocol.FrameData) error {
	message := &protocol.Request{
		Room:&protocol.Request_FrameData{
			FrameData:request,
		},
	}
	return this.Send(node, message)
}

func (this *roomRPCHandler) PreJoinGame(node string, request *protocol.PreJoinGame) error {
	message := &protocol.Request{
		Room:&protocol.Request_PreJoinGame{
			PreJoinGame:request,
		},
	}
	return this.Send(node, message)
}

func (this *roomRPCHandler) GameData(node string, request *protocol.GameData) error {
	message := &protocol.Request{
		Room:&protocol.Request_GameData{
			GameData:request,
		},
	}
	return this.Send(node, message)
}

func (this *roomRPCHandler) GameReady(node string, request *protocol.GameReady) error {
	message := &protocol.Request{
		Room:&protocol.Request_GameReady{
			GameReady:request,
		},
	}
	return this.Send(node, message)
}

func (this *roomRPCHandler) ContinueJoinGame(node string, request *protocol.ContinueJoinGame) error {
	message := &protocol.Request{
		Room:&protocol.Request_ContinueJoinGame{
			ContinueJoinGame:request,
		},
	}
	return this.Send(node, message)
}

func (this *roomRPCHandler) RespondJoinGame(node string, request *protocol.RespondJoinGame) error {
	message := &protocol.Request{
		Room:&protocol.Request_RespondJoinGame{
			RespondJoinGame:request,
		},
	}
	return this.Send(node, message)
}

func (this *roomRPCHandler) BroadcastViewer(node string, request *protocol.BroadcastViewer) error {
	message := &protocol.Request{
		Room:&protocol.Request_BroadcastViewer{
			BroadcastViewer:request,
		},
	}
	return this.Send(node, message)
}
