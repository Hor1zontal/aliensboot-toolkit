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
	"aliens/protocol/passport"
	"github.com/gogo/protobuf/proto"
	"golang.org/x/net/context"
	"github.com/pkg/errors"
	"aliens/protocol"
	"aliens/exception"
	"aliens/common/util"
)

type passportService struct {
}

func (this *passportService) Request(ctx context.Context,request *protocol.Any) (response *protocol.Any,err error) {
	isJSONRequest := request.TypeUrl != ""
	if isJSONRequest {
		data, error := handleJsonRequest(request.TypeUrl, request.Value)
		if error != nil {
			return nil, error
		}
		return &protocol.Any{TypeUrl:"", Value:data}, nil
	}

	requestProxy := &passport.PassportRequest{}
	error := proto.Unmarshal(request.Value, requestProxy)
	if error != nil {
		return nil, error
	}
	responseProxy := &passport.PassportResponse{Session:requestProxy.GetSession()}

    defer func() {
    	//处理消息异常
    	if err := recover(); err != nil {
    		switch err.(type) {
    		    case exception.GameCode:
    				responseProxy.Response = &passport.PassportResponse_Exception{Exception:uint32(err.(exception.GameCode))}
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

func handleRequest(request *passport.PassportRequest, response *passport.PassportResponse) error {
	
	if request.GetLoginRegister() != nil {
		messageRet := &passport.LoginRegisterRet{}
		handleLoginRegister(request.GetLoginRegister(), messageRet)
		response.Response = &passport.PassportResponse_LoginRegisterRet{messageRet}
		return nil
	}
	
	if request.GetLoginLogin() != nil {
		messageRet := &passport.LoginLoginRet{}
		handleLoginLogin(request.GetLoginLogin(), messageRet)
		response.Response = &passport.PassportResponse_LoginLoginRet{messageRet}
		return nil
	}
	
	if request.GetNewInterface() != nil {
		messageRet := &passport.NewInterfaceRet{}
		handleNewInterface(request.GetNewInterface(), messageRet)
		response.Response = &passport.PassportResponse_NewInterfaceRet{messageRet}
		return nil
	}
	
	return errors.New("unexpect request")

}

