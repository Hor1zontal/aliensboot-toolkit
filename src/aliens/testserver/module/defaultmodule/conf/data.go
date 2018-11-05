/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/10/31
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package conf

import (
	"aliens/aliensbot/config"
	"aliens/aliensbot/cluster/center"
)

func Init() {

}

func Close() {

}

func RegisterData(config config.DataConfig) {
	center.ClusterCenter.SubscribeConfig(config.GetPath(), config.DataChange)
}