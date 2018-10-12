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
	"aliens/cluster/center/service"
)


type Route struct
{
	Service string `json:"service"`
	Seq uint16 `json:"seq"`
	Auth bool `json:"auth"`
}

var (

	// aliens conf
	PendingWriteNum        = 2000
	MaxMsgLen       uint32 = 4096
	HTTPTimeout            = 10 * time.Second
	LenMsgLen              = 2
	LittleEndian           = true
)

var Config struct {
	//Enable              bool   //网络模块是否开启
	Service 			service.Config

	MaxConnNum          int
	WSAddr              string //
	TCPAddr             string //
	HTTPAddr			string //
	SecretKey           string //
	AuthTimeout         float64
	HeartbeatTimeout    float64
	Route	[]Route   //路由配置
}

func Init(name string) {
	Config.Service.Name = name
}
