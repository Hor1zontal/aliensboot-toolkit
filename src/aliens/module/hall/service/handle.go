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
	"aliens/protocol/hall"
	"github.com/gogo/protobuf/proto"
	"golang.org/x/net/context"
	"github.com/pkg/errors"
	"aliens/protocol"
	"aliens/exception"
	"aliens/common/util"
	"github.com/name5566/leaf/chanrpc"
)

func newService(chanRpc *chanrpc.Server) *hallService {
	service := &hallService{}
	service.chanRpc = chanRpc
	service.chanRpc.Register("m", service.handle)
	return service
}

type hallService struct {
	chanRpc *chanrpc.Server
}

func (this *hallService) handle(args []interface{}) {
	request := args[0].(*scene_model_proto.Any)
	server := args[1].(scene_model_proto.RPCService_RequestServer)
	response, _ := this.RequestProxy(nil, request)
	server.Send(response)
}

func (this *hallService) Request(request *scene_model_proto.Any, server scene_model_proto.RPCService_RequestServer) error {
	if this.chanRpc != nil {
		this.chanRpc.Call0("m", request, server)
		return nil
	}
	return nil
}

func (this *hallService) RequestProxy(ctx context.Context,request *scene_model_proto.Any) (response *scene_model_proto.Any, err error) {
	isJSONRequest := request.TypeUrl != ""
	if isJSONRequest {
		data, error := handleJsonRequest(request.TypeUrl, request.Value)
		if error != nil {
			return nil, error
		}
		return &scene_model_proto.Any{TypeUrl:"", Value:data}, nil
	}

	requestProxy := &hall.HallRequest{}
	error := proto.Unmarshal(request.Value, requestProxy)
	if error != nil {
		return nil, error
	}
	responseProxy := &hall.HallResponse{Session:requestProxy.GetSession()}

    defer func() {
    	//处理消息异常
    	if err := recover(); err != nil {
    		switch err.(type) {
    		    case exception.GameCode:
    				responseProxy.Response = &hall.HallResponse_Exception{Exception:uint32(err.(exception.GameCode))}
    				break
    			default:
    				util.PrintStackDetail()
    				//未知异常不需要回数据
                    response = nil
                    return
    			}
    	}
    	data, _ := proto.Marshal(responseProxy)
        response = &scene_model_proto.Any{Value:data}
    }()
	err = handleRequest(requestProxy, responseProxy)
    return
}

func handleRequest(request *hall.HallRequest, response *hall.HallResponse) error {
	
	if request.GetQuickMatch() != nil {
		messageRet := &hall.QuickMatchRet{}
		handleQuickMatch(request.GetQuickMatch(), messageRet)
		response.Response = &hall.HallResponse_QuickMatchRet{messageRet}
		return nil
	}
	
	return errors.New("unexpect request")

}

