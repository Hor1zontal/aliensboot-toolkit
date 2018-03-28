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
	"io/ioutil"
	"encoding/json"
	"aliens/log"
)


var Config struct {
	Enable              bool   //场景模块是否开启
	RPCPort				int    //rpc端口
}

func init() {
	data, err := ioutil.ReadFile("conf/aliens/scene/server.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &Config)
	if err != nil {
		log.Critical("%v", err)
	}
	log.Debug("json init %v", Config)
}
