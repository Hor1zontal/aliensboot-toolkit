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
	"aliens/module/cluster/dispatch"
	"fmt"
	"aliens/protocol/base"
)

//requestID - service
var seqServiceMapping = make(map[uint16]string)

//service/alias - responseID
var serviceSeqMapping = make(map[string]uint16)


type Route struct
{
	Service string `json:"service"`
	Seq uint16 `json:"seq"`
	Auth bool `json:"auth"`
}

func LoadRoute(routes []Route) {
	for _, route := range routes {
		if route.Service == "" {
			continue
		}
		seqServiceMapping[route.Seq] = route.Service
		serviceSeqMapping[route.Service] = route.Seq
	}
}

func HandleUrlMessage(requestURL string, requestData []byte) ([]byte, error) {
	params := strings.Split(requestURL, "/")
	if len(params) < 3 {
		return nil, errors.New("invalid param")
	}

	serviceID := params[1]
	request := &base.Any{TypeUrl:params[2], Value:requestData}
	response, error := dispatch.RPC.Request(serviceID, request)
	if error != nil {
		return nil, error
	}
	responseProxy, ok := response.(*base.Any)
	if !ok {
		return nil, errors.New("unexpect response type")
	}
	return responseProxy.Value, nil
}

func GetPushID(service string) uint16 {
	return serviceSeqMapping[service]
}

//未授权的消息转发
func HandleMessage(request *base.Any) (*base.Any, error) {
	serviceName, ok := seqServiceMapping[request.Id]
	if !ok {
		return nil, errors.New(fmt.Sprintf("un expect request id %v", request.Id))
	}
	response, error := dispatch.RPC.Request(serviceName, request)
	if error != nil {
		return nil, error
	}
	responseProxy, ok := response.(*base.Any)
	if !ok {
		return nil, errors.New("un expect response type")
	}
	responseProxy.Id = request.Id
	return responseProxy, nil
}