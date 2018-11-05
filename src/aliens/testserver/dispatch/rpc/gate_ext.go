/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/8/25
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package rpc

import (
	"aliens/aliensbot/cluster/center/service"
	"aliens/testserver/dispatch"
	"aliens/testserver/protocol"
	"errors"
	"github.com/gogo/protobuf/proto"
)

func (this *gateRPCHandle) BindService1(authID int64, node string, service service.IService) error {
	if service == nil {
		return errors.New("service can not be nil")
	}
	service.GetName()
	service.GetID()

	request := &protocol.BindService{
		AuthID: authID,
		//Binds:center.ClusterCenter.GetNodeID()
	}
	return this.BindService(node, request)
}

//推送玩家消息
func (this *gateRPCHandle) Push(authID int64, node string, response *protocol.Response) error {
	data, err := proto.Marshal(response)
	if err != nil {
		return err
	}
	pushMessage := &protocol.PushMessage{
		AuthID: authID,
		Data:   data,
	}
	this.PushMessage(node, pushMessage)
	return nil
}

func (this *gateRPCHandle) BroadcastAll(node string, response *protocol.Response) {
	data, _ := proto.Marshal(response)
	message := &protocol.Request{
		Gate: &protocol.Request_PushMessage{
			PushMessage: &protocol.PushMessage{
				AuthID: -1,
				Data:   data,
			},
		},
	}
	dispatch.Broadcast(this.name, message)
}
