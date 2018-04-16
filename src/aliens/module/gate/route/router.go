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
	"errors"
	"strings"
	"aliens/protocol"
	"aliens/module/cluster/dispatch"
)

//requestID - service
var requestServiceMapping = make(map[uint16]string)

//service/alias - pushID
var servicePushMapping = make(map[string]uint16)

//requestID - responseID
var responseMapping = make(map[uint16]uint16)

type Route struct
{
	RequestID uint16 `json:"requestID"`
	ResponseID uint16 `json:"responseID"`
	PushID uint16 `json:"pushID"`
	Service string `json:"service"`
	Alias string `json:"alias"`
}

func LoadRoute(routes []Route) {
	for _, route := range routes {
		if route.Service == "" {
			continue
		}
		requestServiceMapping[route.RequestID] = route.Service
		responseMapping[route.RequestID] = route.ResponseID
		servicePushMapping[route.Alias] = route.PushID
	}
}

func HandleUrlMessage(requestURL string, requestData []byte) ([]byte, error) {
	params := strings.Split(requestURL, "/")
	if len(params) < 3 {
		return nil, errors.New("invalid param")
	}

	serviceID := params[1]
	request := &protocol.Any{TypeUrl:params[2], Value:requestData}
	response, error := dispatch.Request(serviceID, request)
	if error != nil {
		return nil, error
	}
	responseProxy, ok := response.(*protocol.Any)
	if !ok {
		return nil, errors.New("unexpect response type")
	}
	return responseProxy.Value, nil
}

func GetPushID(service string) uint16 {
	return servicePushMapping[service]
}


func HandleMessage(request interface{}, sessionID string) (interface{}, error) {
	any, _ := request.(*protocol.Any)
	serviceID, ok := requestServiceMapping[any.Id]
	if !ok {
		return nil, errors.New("unexpect requestID")
	}
	if sessionID != "" {
		any.SessionId = sessionID
	}
	response, error := dispatch.Request(serviceID, request)
	if error != nil {
		return nil, error
	}
	responseProxy, ok := response.(*protocol.Any)
	if !ok {
		return nil, errors.New("unexpect response type")
	}
	responseProxy.Id = responseMapping[any.Id]
	return responseProxy, nil
}