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
	
	if request.GetTokenLogin() != nil {
		messageRet := &protocol.TokenLoginRet{}
		result := handleTokenLogin(request.GetTokenLogin(), messageRet)
		response.Passport = &protocol.Response_TokenLoginRet{messageRet}
		return result
	}
	
	if request.GetLoginRegister() != nil {
		messageRet := &protocol.LoginRegisterRet{}
		result := handleLoginRegister(request.GetLoginRegister(), messageRet)
		response.Passport = &protocol.Response_LoginRegisterRet{messageRet}
		return result
	}
	
	if request.GetLoginLogin() != nil {
		messageRet := &protocol.LoginLoginRet{}
		result := handleLoginLogin(request.GetLoginLogin(), messageRet)
		response.Passport = &protocol.Response_LoginLoginRet{messageRet}
		return result
	}
	
	response.Code = protocol.Code_InvalidRequest
	return 0
}

