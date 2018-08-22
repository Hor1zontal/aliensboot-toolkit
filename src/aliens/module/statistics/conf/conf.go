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

const (
	AnalysisFlag = false //是否开启性能分析
	Game = "aliens" //日志索引信息
)

var Config struct {
	ES 				ESConfig
}

type ESConfig struct {
	Url string
	Host string
	Username string
	Password string
}