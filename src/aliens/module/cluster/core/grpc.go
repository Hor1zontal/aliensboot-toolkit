/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/8/22
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core


import (
	"aliens/protocol"
	"aliens/protocol/base"
	"github.com/gogo/protobuf/proto"
	"aliens/cluster/center/dispatch"
	"aliens/cluster/center"
	"aliens/module/cluster/conf"
)

func Init() {
	center.ClusterCenter.ConnectCluster(conf.Config.Cluster)
}

func Close() {
	center.ClusterCenter.Close()
}


//阻塞请求消息 - 根据负载均衡动态分配一个节点处理
func Request(serviceName string, message  *protocol.Request, hashKey string) (*protocol.Response, error) {
	data, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}
	request := &base.Any{Value: data}
	response, err := dispatch.Request(serviceName, request, hashKey)
	if err != nil {
		return nil, err
	}
	messageRet := &protocol.Response{}
	messageRet.Unmarshal(response.GetValue())
	return  messageRet, nil
}

//同步阻塞请求
func RequestNode(serviceName string, serviceID string, message *protocol.Request) (*protocol.Response, error) {
	data, _ := message.Marshal()
	request := &base.Any{Value: data}
	response, err := dispatch.RequestNode(serviceName, serviceID, request)
	if err != nil {
		return nil, err
	}
	messageRet := &protocol.Response{}
	messageRet.Unmarshal(response.GetValue())
	return  messageRet, nil
}





