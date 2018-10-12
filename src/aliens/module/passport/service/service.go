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
	"github.com/gogo/protobuf/proto"
    "aliens/chanrpc"
    "aliens/exception"
    "aliens/protocol/base"
    "aliens/protocol"
    "aliens/cluster/center/service"
    "aliens/module/passport/conf"
    "aliens/cluster/center"
    "aliens/log"
)

var instance service.IService = nil

func Init(chanRpc *chanrpc.Server) {
	instance = center.PublicService(conf.Config.Service, service.NewRpcHandler(chanRpc, handle))
}

func Close() {
	center.ReleaseService(instance)
}

func handle(request *base.Any) *base.Any {
	requestProxy := &protocol.Request{}
	responseProxy := &protocol.Response{}
	response := &base.Any{}
	authID := request.GetAuthId()
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
		data, _ := proto.Marshal(responseProxy)
		responseProxy.Session = requestProxy.GetSession()
        response.AuthId = authID
		response.Value = data
	}()
	error := proto.Unmarshal(request.Value, requestProxy)
	if error != nil {
	    log.Debug(error)
		exception.GameException(protocol.Code_InvalidRequest)
	}
	authID = handleRequest(requestProxy, responseProxy)
	return response
}

func handleRequest(request *protocol.Request, response *protocol.Response) int64 {
	
	if request.GetC2S_UserRegister() != nil {
		messageRet := &protocol.S2C_UserRegister{}
		result := handleUserRegister(request.GetC2S_UserRegister(), messageRet)
		response.Passport = &protocol.Response_S2C_UserRegister{messageRet}
		return result
	}
	
	if request.GetC2S_UserLogin() != nil {
		messageRet := &protocol.S2C_UserLogin{}
		result := handleUserLogin(request.GetC2S_UserLogin(), messageRet)
		response.Passport = &protocol.Response_S2C_UserLogin{messageRet}
		return result
	}
	
	if request.GetC2S_TokenLogin() != nil {
		messageRet := &protocol.S2C_TokenLogin{}
		result := handleTokenLogin(request.GetC2S_TokenLogin(), messageRet)
		response.Passport = &protocol.Response_S2C_TokenLogin{messageRet}
		return result
	}
	
	response.Code = protocol.Code_InvalidRequest
	return 0
}

