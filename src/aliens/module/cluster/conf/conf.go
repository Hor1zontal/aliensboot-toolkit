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
)

var Config struct {
	//Node        string	//当前集群节点的标识，信息需要注册到中心服务器
	ZKServers []string //集群中心服务器地址
	ZKName    string   //集群名称，不用业务使用不同的集群
	LBS       string   //负载均衡策略  polling 轮询
}

func init() {
	config.LoadConfig(&Config, "conf/aliens/cluster.json")
}
