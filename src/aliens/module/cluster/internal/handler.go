/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/8/23
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package internal

import (
	"aliens/module/cluster/dispatch"
	"aliens/protocol/base"
)

const (
	commandRequest = "request"
	commandRequestNode = "requestNode"
	commandSend = "send"
	commandSendNode = "sendNode"
	commandPublicService = ""
)

func init() {
	Skeleton.RegisterChanRPC(commandRequest, handleRequest)
	Skeleton.RegisterChanRPC(commandRequestNode, handleRequest)
	Skeleton.RegisterChanRPC(commandSend, handleRequest)
	Skeleton.RegisterChanRPC(commandSendNode, handleRequest)
	Skeleton.RegisterChanRPC(commandPublicService, handleRequest)
	//dispatch.MQ.RegisterConsumer(constant.SERVICE_GATE, HandlePush)
}


func handleRequest(args []interface{}) []interface{} {
	serviceName := args[0].(string)
	request := args[1].(*base.Any)
	hashKey := args[2].(string)
	response, err := dispatch.Request(serviceName, request, hashKey)
	return []interface{} {response, err}
}