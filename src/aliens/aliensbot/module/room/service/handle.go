/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/3/30
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package service

import (
	"aliens/aliensbot/common/util"
	"aliens/aliensbot/exception"
	"aliens/aliensbot/network"
	"aliens/aliensbot/protocol"
	"aliens/aliensbot/protocol/room"
	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type roomService struct {
}

func (this *roomService) Request(ctx context.Context, request *scene_model_proto.Any) (response *scene_model_proto.Any, err error) {
	isJSONRequest := request.TypeUrl != ""
	if isJSONRequest {
		data, error := handleJsonRequest(request.TypeUrl, request.Value)
		if error != nil {
			return nil, error
		}
		return &scene_model_proto.Any{TypeUrl: "", Value: data}, nil
	}

	requestProxy := &room.Request{}
	error := proto.Unmarshal(request.Value, requestProxy)
	if error != nil {
		return nil, error
	}
	responseProxy := &room.Response{Session: requestProxy.GetSession()}

	defer func() {
		//处理消息异常
		if err := recover(); err != nil {
			switch err.(type) {
			case exception.GameCode:
				responseProxy.Response = &room.Response_Exception{Exception: uint32(err.(exception.GameCode))}
				break
			default:
				util.PrintStackDetail()
				//未知异常不需要回数据
				response = nil
				return
			}
		}
		data, _ := proto.Marshal(responseProxy)
		response = &scene_model_proto.Any{TypeUrl: "", Value: data}
	}()
	err = handleRequest(requestProxy, responseProxy, request.Agent)
	return
}

func handleRequest(request *room.Request, response *room.Response, agent network.Agent) error {

	if request.GetLeaveRoom() != nil {
		messageRet := &room.LeaveRoomRet{}
		handleLeaveRoom(request.GetLeaveRoom(), messageRet, agent)
		response.Response = &room.Response_LeaveRoomRet{messageRet}
		return nil
	}

	if request.GetAllocFreeRoomSeat() != nil {
		messageRet := &room.AllocFreeRoomSeatRet{}
		handleAllocFreeRoomSeat(request.GetAllocFreeRoomSeat(), messageRet, agent)
		response.Response = &room.Response_AllocFreeRoomSeatRet{messageRet}
		return nil
	}

	if request.GetJoinRoom() != nil {
		messageRet := &room.JoinRoomRet{}
		handleJoinRoom(request.GetJoinRoom(), messageRet, agent)
		response.Response = &room.Response_JoinRoomRet{messageRet}
		return nil
	}

	return errors.New("unexpect request")

}
