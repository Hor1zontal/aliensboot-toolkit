/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/8/24
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package service

import (
	"aliens/protocol/base"
	"github.com/AsynkronIT/protoactor-go/actor"
	"context"
)

type call struct {
	request *base.Any
	callback Callback
}

type rpcClient struct {
	client base.RPCServiceClient
}

func (this *rpcClient) Receive(actorContext actor.Context) {
	switch msg := actorContext.Message().(type) {
	//case *actor.Stopped:
	//	//fmt.Println("")
	case *call:
		ctx, _ := context.WithTimeout(context.Background(), requestTimeout)
		client, err := this.client.Request(ctx, msg.request)
		if err != nil {
			msg.callback(nil, err)
			break
		}
		response, err := client.Recv()
		if response != nil {
			response.Id = msg.request.Id
		}
		msg.callback(response, err)
	}
}
