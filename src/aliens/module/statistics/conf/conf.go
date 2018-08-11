/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/5/10
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package conf

import (
	"aliens/config"
	"aliens/module/base"
)

const (
	configPath = base.BaseConfPath + "statistics/config.json"
	AnalysisFlag = false //是否开启性能分析
	Game = "aliens" //日志索引信息
)

var Config struct {
	Enable			bool
	ES 				ESConfig
}

type ESConfig struct {
	Url string
	Host string
	Username string
	Password string
}

func init() {
	config.LoadConfig(&Config, configPath) //加载服务器配置
}
