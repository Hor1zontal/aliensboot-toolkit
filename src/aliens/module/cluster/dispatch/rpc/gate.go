/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/6/15
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package rpc

import (
	"aliens/protocol"
	"aliens/exception"
	"aliens/log"
	"aliens/protocol/base"
	"github.com/gogo/protobuf/proto"
	"aliens/module/cluster/dispatch"
)

var Gate = &gateRPCHandle{"gate"}


type gateRPCHandle struct {
	name string
}


func (this *gateRPCHandle) RequestNode(node string, request *protocol.Request) *protocol.Response {
	rpcRet, err := dispatch.RequestNodeMessage(this.name, node, request)
	if err != nil {
    	log.Error(err)
    	exception.GameException(protocol.Code_InvalidService)
    }
    return rpcRet
}

//func (this *sceneRPCHandle) Request(request *protocol.Request) *protocol.Response {
//	rpcRet, err := dispatch.RPC.SyncRequest(this.name, request)
//	if err != nil {
//        log.Error(err)
//        exception.GameException(protocol.Code_InvalidService)
//    }
//    return rpcRet
//}

//推送玩家消息
func (this *gateRPCHandle) Push(uid int64, gateID string, response *protocol.Response) error {
	data, err := proto.Marshal(response)
	if err != nil {
		return err
	}
	pushMessage := &base.Any{AuthId: uid, Value: data}
	return dispatch.SendNode(this.name, gateID, pushMessage)
}
