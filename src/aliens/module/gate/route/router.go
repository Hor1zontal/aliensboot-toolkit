/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/3/30
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package route

import (
	"aliens/cluster/message"
	"github.com/gogo/protobuf/types"
	"errors"
	"strings"
)

//requestID - service
var RequestServiceMapping = make(map[uint16]*message.RemoteService)

//url - service
var URLServiceMapping = make(map[string]*message.RemoteService)

//service - responseID
var ServiceResponseMapping = make(map[string]uint16)

type Route struct
{
	Service string `json:"service"`
	RequestID uint16 `json:"requestID"`
	ResponseID uint16 `json:"responseID"`
}

func LoadRoute(routes []Route) {
	for _, route := range routes {
		if route.Service == "" {
			continue
		}
		service := RequestServiceMapping[route.RequestID]

		serviceHandler := URLServiceMapping[route.Service]
		if serviceHandler == nil {
			serviceHandler = message.NewRemoteService(route.Service)
			URLServiceMapping[route.Service] = serviceHandler
		}
		//服务
		if service == nil || service.GetType() != route.Service {
			RequestServiceMapping[route.RequestID] = serviceHandler
		}
		ServiceResponseMapping[route.Service] = route.ResponseID
	}
}

func HandleUrlMessage(requestURL string, requestData []byte) ([]byte, error) {
	params := strings.Split(requestURL, "/")
	if len(params) < 3 {
		return nil, errors.New("invalid param")
	}

	messageService := URLServiceMapping[params[1]]
	if messageService == nil {
		return nil, errors.New("unexpect request url " + params[1])
	}
	request := &types.Any{TypeUrl:params[2], Value:requestData}
	response, error := messageService.HandleMessage(request)
	if error != nil {
		return nil, error
	}
	responseProxy, ok := response.(*types.Any)
	if !ok {
		return nil, errors.New("unexpect response type")
	}
	return responseProxy.Value, nil
}


func HandleMessage(request interface{}) (interface{}, error) {
	any, _ := request.(*types.Any)
	messageService := RequestServiceMapping[any.ID]
	if messageService == nil {
		return nil, errors.New("unexpect requestID")
	}
	response, error := messageService.HandleMessage(request)
	if error != nil {
		return nil, error
	}
	responseProxy, ok := response.(*types.Any)
	if !ok {
		return nil, errors.New("unexpect response type")
	}
	responseProxy.ID = ServiceResponseMapping[messageService.GetType()]
	return responseProxy, nil
}