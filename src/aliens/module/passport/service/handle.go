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
    "fmt"
    "aliens/common/util"
    "runtime"
    "aliens/log"
)

type passportService struct {
}

func (this *passportService) Request(ctx context.Context,request *protocol.Any) (any *protocol.Any,err error) {
    defer func() {
		//处理消息异常
		if err := recover(); err != nil {
			switch err.(type) {
			case exception.GameCode:
				any = nil
				err = errors.New(util.Int32ToString(int32(err.(exception.GameCode))))
				break
			default:
				buf := make([]byte, 2048)
				n := runtime.Stack(buf, false)
				log.Error("panic stack info %s", fmt.Sprintf("%s", buf[:n]))
			}
		}
	}()
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
	//response, error := this.HandleRequest(ctx, requestProxy)
	response, error := handleRequest(requestProxy)

	if response == nil {
		return nil, error
	}
	data, _ := proto.Marshal(response)
	responseProxy := &protocol.Any{TypeUrl:"", Value:data}
	return responseProxy, error
}

func handleRequest(request *passport.PassportRequest) (*passport.PassportResponse, error) {
	response := &passport.PassportResponse{Session:request.GetSession()}

	
	if request.GetLoginRegister() != nil {
		messageRet := &passport.LoginRegisterRet{}
		handleLoginRegister(request.GetLoginRegister(), messageRet)
		response.Response = &passport.PassportResponse_LoginRegisterRet{messageRet}
		return response, nil
	}
	
	if request.GetLoginLogin() != nil {
		messageRet := &passport.LoginLoginRet{}
		handleLoginLogin(request.GetLoginLogin(), messageRet)
		response.Response = &passport.PassportResponse_LoginLoginRet{messageRet}
		return response, nil
	}
	
	return nil, errors.New("unexpect request")

}

