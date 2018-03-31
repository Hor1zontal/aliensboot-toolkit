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
	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
)

type sceneService struct {
}

func (this *sceneService) Request(ctx context.Context,request *types.Any) (*types.Any, error) {
	requestProxy := &scene.SceneRequest{}
	error := proto.Unmarshal(request.Value, requestProxy)
	if error != nil {
		return nil, error
	}
	//response, error := this.HandleRequest(ctx, requestProxy)
	response, error := handleRequest(requestProxy)


	if response == nil {
		return nil, error
	}
	data, _ := proto.Marshal(response)
	responseProxy := &types.Any{TypeUrl:"", Value:data}
	return responseProxy, error
}

func handleRequest(request *scene.SceneRequest) (*scene.SceneResponse, error) {
	response := &scene.SceneResponse{Session:request.GetSession()}
	
	if request.GetSpaceMove() != nil {
		messageRet := &scene.SpaceMoveRet{}
		handleSpaceMove(request.GetSpaceMove(), messageRet)
		response.Response = &scene.SceneResponse_SpaceMoveRet{messageRet}
		return response, nil
	}
	
	if request.GetSpaceEnter() != nil {
		messageRet := &scene.SpaceEnterRet{}
		handleSpaceEnter(request.GetSpaceEnter(), messageRet)
		response.Response = &scene.SceneResponse_SpaceEnterRet{messageRet}
		return response, nil
	}
	
	if request.GetSpaceLeave() != nil {
		messageRet := &scene.SpaceLeaveRet{}
		handleSpaceLeave(request.GetSpaceLeave(), messageRet)
		response.Response = &scene.SceneResponse_SpaceLeaveRet{messageRet}
		return response, nil
	}
	
	if request.GetGetState() != nil {
		messageRet := &scene.GetStateRet{}
		handleGetState(request.GetGetState(), messageRet)
		response.Response = &scene.SceneResponse_GetStateRet{messageRet}
		return response, nil
	}
	
	return nil, errors.New("unexpect request")

}

