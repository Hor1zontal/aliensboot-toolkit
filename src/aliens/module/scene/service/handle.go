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
	"aliens/protocol/scene"
	"github.com/gogo/protobuf/proto"
	"golang.org/x/net/context"
	"github.com/pkg/errors"
	"aliens/protocol"
	"aliens/exception"
	"aliens/common/util"
)

type sceneService struct {
}

func (this *sceneService) Request(ctx context.Context,request *protocol.Any) (response *protocol.Any,err error) {
	isJSONRequest := request.TypeUrl != ""
	if isJSONRequest {
		data, error := handleJsonRequest(request.TypeUrl, request.Value)
		if error != nil {
			return nil, error
		}
		return &protocol.Any{TypeUrl:"", Value:data}, nil
	}

	requestProxy := &scene.SceneRequest{}
	error := proto.Unmarshal(request.Value, requestProxy)
	if error != nil {
		return nil, error
	}
	responseProxy := &scene.SceneResponse{Session:requestProxy.GetSession()}

     defer func() {
    		//处理消息异常
    		if err := recover(); err != nil {
    			switch err.(type) {
    			case exception.GameCode:
    				responseProxy.Response = &scene.SceneResponse_Exception{Exception:uint32(err.(exception.GameCode))}
    				break
    			default:
    				util.PrintStackDetail()
    				//未知异常不需要回数据
                    response = nil
                    return
    			}
    		}

    		data, _ := proto.Marshal(response)
            response = &protocol.Any{TypeUrl:"", Value:data}
    	}()
	err = handleRequest(requestProxy, responseProxy)
    return
}

func handleRequest(request *scene.SceneRequest, response *scene.SceneResponse) error {
	
	if request.GetSpaceEnter() != nil {
		messageRet := &scene.SpaceEnterRet{}
		handleSpaceEnter(request.GetSpaceEnter(), messageRet)
		response.Response = &scene.SceneResponse_SpaceEnterRet{messageRet}
		return nil
	}
	
	if request.GetSpaceLeave() != nil {
		messageRet := &scene.SpaceLeaveRet{}
		handleSpaceLeave(request.GetSpaceLeave(), messageRet)
		response.Response = &scene.SceneResponse_SpaceLeaveRet{messageRet}
		return nil
	}
	
	if request.GetGetState() != nil {
		messageRet := &scene.GetStateRet{}
		handleGetState(request.GetGetState(), messageRet)
		response.Response = &scene.SceneResponse_GetStateRet{messageRet}
		return nil
	}
	
	if request.GetSpaceMove() != nil {
		messageRet := &scene.SpaceMoveRet{}
		handleSpaceMove(request.GetSpaceMove(), messageRet)
		response.Response = &scene.SceneResponse_SpaceMoveRet{messageRet}
		return nil
	}
	
	return errors.New("unexpect request")

}

