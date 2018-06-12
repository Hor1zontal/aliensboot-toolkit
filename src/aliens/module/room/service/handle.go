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
	"aliens/protocol/room"
	"github.com/gogo/protobuf/proto"
	"golang.org/x/net/context"
	"github.com/pkg/errors"
	"aliens/protocol"
	"aliens/exception"
	"aliens/common/util"
)

type roomService struct {
}

func (this *roomService) Request(ctx context.Context,request *protocol.Any) (response *protocol.Any,err error) {
	isJSONRequest := request.TypeUrl != ""
	if isJSONRequest {
		data, error := handleJsonRequest(request.TypeUrl, request.Value)
		if error != nil {
			return nil, error
		}
		return &protocol.Any{TypeUrl:"", Value:data}, nil
	}

	requestProxy := &room.Request{}
	error := proto.Unmarshal(request.Value, requestProxy)
	if error != nil {
		return nil, error
	}
	responseProxy := &room.Response{Session:requestProxy.GetSession()}

    defer func() {
    	//处理消息异常
    	if err := recover(); err != nil {
    		switch err.(type) {
    		    case exception.GameCode:
    				responseProxy.Response = &room.Response_Exception{Exception:uint32(err.(exception.GameCode))}
    				break
    			default:
    				util.PrintStackDetail()
    				//未知异常不需要回数据
                    response = nil
                    return
    			}
    	}
    	data, _ := proto.Marshal(responseProxy)
        response = &protocol.Any{TypeUrl:"", Value:data}
    }()
	err = handleRequest(requestProxy, responseProxy)
    return
}

func handleRequest(request *room.Request, response *room.Response) error {
	
	if request.GetCreateRoom() != nil {
		messageRet := &room.CreateRoomRet{}
		handleCreateRoom(request.GetCreateRoom(), messageRet)
		response.Response = &room.Response_CreateRoomRet{messageRet}
		return nil
	}
	
	if request.GetJoinRoom() != nil {
		messageRet := &room.JoinRoomRet{}
		handleJoinRoom(request.GetJoinRoom(), messageRet)
		response.Response = &room.Response_JoinRoomRet{messageRet}
		return nil
	}
	
	if request.GetLeaveRoom() != nil {
		messageRet := &room.LeaveRoomRet{}
		handleLeaveRoom(request.GetLeaveRoom(), messageRet)
		response.Response = &room.Response_LeaveRoomRet{messageRet}
		return nil
	}
	
	return errors.New("unexpect request")

}

