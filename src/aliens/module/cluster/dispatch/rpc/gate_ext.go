/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/8/25
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package rpc

import (
	"aliens/protocol"
	"aliens/protocol/base"
	"aliens/module/cluster/dispatch"
	"github.com/gogo/protobuf/proto"
)

//推送玩家消息
func (this *gateRPCHandle) Push(uid int64, gateID string, response *protocol.Response) error {
	data, err := proto.Marshal(response)
	if err != nil {
		return err
	}
	pushMessage := &base.Any{AuthId: uid, Value: data}
	return dispatch.SendNode(this.name, gateID, pushMessage)
}

