/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/3/30
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package route

import (
	"aliens/cluster/center/service"
	"aliens/module/dispatch"
	"aliens/module/gate/conf"
	"aliens/protocol/base"
	"errors"
	"fmt"
)

//requestID - service
var seqServiceMapping = make(map[uint16]string)

//service/alias - responseID
var serviceSeqMapping = make(map[string]uint16)

//goroutine pool, deal async request and callback

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

func GetServiceSeq(serviceName string) uint16 {
	return serviceSeqMapping[serviceName]
}

func GetServiceByeSeq(seq uint16) string {
	return seqServiceMapping[seq]
}

func HandleUrlMessage(serviceName string, requestData []byte) ([]byte, error) {
	seq := GetServiceSeq(serviceName)
	if seq <= 0 {
		return nil, errors.New(fmt.Sprintf("service %v seq not found", serviceName))
	}
	request := &base.Any{Id: seq, Value: requestData}
	response, error := dispatch.Request(serviceName, request, "")
	if error != nil {
		return nil, error
	}
	return response.Value, nil
}

//func GetPushID(service string) uint16 {
//	return serviceSeqMapping[service]
//}

func AsyncHandleMessage(hashKey string, asyncCall *service.AsyncCall) error {
	serviceName, ok := seqServiceMapping[asyncCall.ReqID()]
	if !ok {
		return errors.New(fmt.Sprintf("un expect request id %v", asyncCall.ReqID()))
	}
	return dispatch.AsyncRequest(serviceName, hashKey, asyncCall)
}

//发送到指定节点
func AsyncHandleNodeMessage(serviceID string, asyncCall *service.AsyncCall) error {
	serviceName, ok := seqServiceMapping[asyncCall.ReqID()]
	if !ok {
		return errors.New(fmt.Sprintf("un expect request id %v", asyncCall.ReqID()))
	}
	return dispatch.AsyncRequestNode(serviceName, serviceID, asyncCall)
}

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

//发送到指定节点
func HandleNodeMessage(request *base.Any, node string) (*base.Any, error) {
	serviceName, ok := seqServiceMapping[request.Id]
	if !ok {
		return nil, errors.New(fmt.Sprintf("un expect request id %v", request.Id))
	}
	response, error := dispatch.RequestNode(serviceName, node, request)
	if error != nil {
		return nil, error
	}
	response.Id = request.Id
	return response, nil
}
