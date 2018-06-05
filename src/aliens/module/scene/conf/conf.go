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
	"aliens/config"
	"aliens/cluster/center/service"
)


var Config struct {
	Service  service.ServiceConfig
}

func init() {
	config.LoadConfig(&Config, "conf/aliens/scene/server.json")
}
