// Code generated by aliensbot. DO NOT EDIT.
// source: module room
package service

import (

	"github.com/gogo/protobuf/proto"
    "aliens/aliensbot/chanrpc"
    "aliens/aliensbot/exception"
    "aliens/aliensbot/cluster/center/service"
    "aliens/aliensbot/cluster/center"
    "aliens/aliensbot/protocol/base"
    "aliens/testserver/protocol"
    "aliens/testserver/module/room/conf"

)

var instance service.IService = nil

func Init(chanRpc *chanrpc.Server) {
	instance = center.PublicService(conf.Config.Service, service.NewRpcHandler(chanRpc, handle))
}

func Close() {
	center.ReleaseService(instance)
}


func handle(request *base.Any) (response *base.Any) {
	requestProxy := &protocol.Request{}
	responseProxy := &protocol.Response{}
	response = &base.Any{}
	isResponse := false
	defer func() {
		//处理消息异常
		if err := recover(); err != nil {
			switch err.(type) {
			case protocol.Code:
				responseProxy.Code = err.(protocol.Code)
				break
			default:
				exception.PrintStackDetail(err)
				responseProxy.Code = protocol.Code_ServerException
			}
		}
		if !isResponse {
            return
        }
		data, _ := proto.Marshal(responseProxy)
		responseProxy.Session = requestProxy.GetSession()
		response.Value = data
	}()
	error := proto.Unmarshal(request.Value, requestProxy)
	if error != nil {
		exception.GameException(protocol.Code_InvalidRequest)
	}
	isResponse = handleRequest(request.GetAuthId(), request.GetGateId(), requestProxy, responseProxy)
	return
}

func handleRequest(authID int64, gateID string, request *protocol.Request, response *protocol.Response) bool {
	
	if request.GetJoinRoom() != nil {
		messageRet := &protocol.JoinRoomRet{}
		handleJoinRoom(authID, gateID, request.GetJoinRoom(), messageRet)
		response.Room = &protocol.Response_JoinRoomRet{messageRet}
		return true
	}
	
	if request.GetShowUser() != nil {
		messageRet := &protocol.ShowUserRet{}
		handleShowUser(authID, gateID, request.GetShowUser(), messageRet)
		response.Room = &protocol.Response_ShowUserRet{messageRet}
		return true
	}
	
	if request.GetRoomCreate() != nil {
		messageRet := &protocol.RoomCreateRet{}
		handleRoomCreate(authID, gateID, request.GetRoomCreate(), messageRet)
		response.Room = &protocol.Response_RoomCreateRet{messageRet}
		return true
	}
	
	if request.GetGetBigoData() != nil {
		messageRet := &protocol.GetBigoDataRet{}
		handleGetBigoData(authID, gateID, request.GetGetBigoData(), messageRet)
		response.Room = &protocol.Response_GetBigoDataRet{messageRet}
		return true
	}
	
	if request.GetOnGameStateChange() != nil {
		messageRet := &protocol.OnGameStateChangeRet{}
		handleOnGameStateChange(authID, gateID, request.GetOnGameStateChange(), messageRet)
		response.Room = &protocol.Response_OnGameStateChangeRet{messageRet}
		return true
	}
	
	if request.GetOnPlayerStateChange() != nil {
		messageRet := &protocol.OnPlayerStateChangeRet{}
		handleOnPlayerStateChange(authID, gateID, request.GetOnPlayerStateChange(), messageRet)
		response.Room = &protocol.Response_OnPlayerStateChangeRet{messageRet}
		return true
	}
	
	if request.GetGetRoomInfo() != nil {
		messageRet := &protocol.GetRoomInfoRet{}
		handleGetRoomInfo(authID, gateID, request.GetGetRoomInfo(), messageRet)
		response.Room = &protocol.Response_GetRoomInfoRet{messageRet}
		return true
	}
	
	
    if request.GetFrameData() != nil {
    	handleFrameData(authID, gateID, request.GetFrameData())
    	return false
    }
    
    if request.GetUpdateBigoData() != nil {
    	handleUpdateBigoData(authID, gateID, request.GetUpdateBigoData())
    	return false
    }
    
    if request.GetBroadcastViewer() != nil {
    	handleBroadcastViewer(authID, gateID, request.GetBroadcastViewer())
    	return false
    }
    
    if request.GetContinueJoinGame() != nil {
    	handleContinueJoinGame(authID, gateID, request.GetContinueJoinGame())
    	return false
    }
    
    if request.GetRequestJoinGame() != nil {
    	handleRequestJoinGame(authID, gateID, request.GetRequestJoinGame())
    	return false
    }
    
    if request.GetGameData() != nil {
    	handleGameData(authID, gateID, request.GetGameData())
    	return false
    }
    
    if request.GetGameReady() != nil {
    	handleGameReady(authID, gateID, request.GetGameReady())
    	return false
    }
    
    if request.GetUploadGameResult() != nil {
    	handleUploadGameResult(authID, gateID, request.GetUploadGameResult())
    	return false
    }
    
    if request.GetPreJoinGame() != nil {
    	handlePreJoinGame(authID, gateID, request.GetPreJoinGame())
    	return false
    }
    
    if request.GetRespondJoinGame() != nil {
    	handleRespondJoinGame(authID, gateID, request.GetRespondJoinGame())
    	return false
    }
    
	response.Code = protocol.Code_InvalidRequest
	return true
}

