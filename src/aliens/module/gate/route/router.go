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
	"fmt"
	"aliens/protocol/base"
	"aliens/cluster/center/dispatch"
	"aliens/module/gate/conf"
)

//requestID - service
var seqServiceMapping = make(map[uint16]string)

//service/alias - responseID
var serviceSeqMapping = make(map[string]uint16)



func Init() {
	routes := conf.Config.Route
	for _, route := range routes {
		if route.Service == "" {
			continue
		}
		seqServiceMapping[route.Seq] = route.Service
		serviceSeqMapping[route.Service] = route.Seq
	}
}

//func HandleUrlMessage(requestURL string, requestData []byte) ([]byte, error) {
//	params := strings.Split(requestURL, "/")
//	if len(params) < 3 {
//		return nil, errors.New("invalid param")
//	}
//
//	serviceID := params[1]
//	request := &base.Any{TypeUrl:params[2], Value:requestData}
//	response, error := dispatch.RPC.Request(serviceID, request, "")
//	if error != nil {
//		return nil, error
//	}
//	responseProxy, ok := response.(*base.Any)
//	if !ok {
//		return nil, errors.New("unexpect response type")
//	}
//	return responseProxy.Value, nil
//}
//
//func GetPushID(service string) uint16 {
//	return serviceSeqMapping[service]
//}


func HandleMessage(request *base.Any, hashKey string) (*base.Any, error) {
	serviceName, ok := seqServiceMapping[request.Id]
	if !ok {
		return nil, errors.New(fmt.Sprintf("un expect request id %v", request.Id))
	}
	response, error := dispatch.Request(serviceName, request, hashKey)
	if error != nil {
		return nil, error
	}
	response.Id = request.Id
	return response, nil
}