/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/6/8
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package internal

import (
	"aliens/aliensbot/log"
	"aliens/aliensbot/network"
	"aliens/aliensbot/protocol/framesync"
	"aliens/testserver/module/room/conf"
)

var server *network.UDPServer = &network.UDPServer{}

func init() {
	server.Start(conf.Config.UDPService, addUDPMessage)
	skeleton.RegisterChanRPC("msg", handleUDPMessage)
}

func addUDPMessage(data []byte, agent *network.UDPAgent) {
	skeleton.ChanRPCServer.Go("msg", data, agent)
}

func handleUDPMessage(args []interface{}) {
	data := args[0].([]byte)
	agent := args[1].(*network.UDPAgent)
	request := &framesync.Request{}
	err := request.Unmarshal(data)
	if err != nil {
		log.Debugf("invalid message format %v", agent)
		return
	}

	//core.Manager.AcceptFrameMessage(request, agent)
}
