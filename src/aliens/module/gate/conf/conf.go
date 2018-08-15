/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2017/8/4
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package conf

import (
	"time"
	"aliens/config"
	"aliens/module/gate/route"
	"aliens/cluster/center/service"
	"aliens/module"
)

var (

	configPath =  module.BaseConfPath + "gate/server.json"
	// skeleton conf
	GoLen              = 10000
	TimerDispatcherLen = 10000
	AsynCallLen        = 10000
	ChanRPCLen         = 10000

	// aliens conf
	PendingWriteNum        = 2000
	MaxMsgLen       uint32 = 4096
	HTTPTimeout            = 10 * time.Second
	LenMsgLen              = 2
	LittleEndian           = true
)

var Config struct {
	//Enable              bool   //网络模块是否开启
	Service 			service.CenterService

	MaxConnNum          int
	WSAddr              string //
	TCPAddr             string //
	HTTPAddr			string //
	SecretKey           string
	MessageChannelLimit int
	AuthTimeout         float64
	HeartbeatTimeout    float64
	Route	[]route.Route   //路由配置
}

func Init() {
	config.LoadConfig(&Config, configPath)
	route.LoadRoute(Config.Route)
}
