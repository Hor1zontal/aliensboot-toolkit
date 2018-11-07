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
	"encoding/json"
	"aliens/aliensbot/exception"
	"aliens/testserver/protocol"
	"aliens/testserver/data"
	"aliens/aliensbot/cluster/center"
)


func Init() {
	center.ClusterCenter.SubscribeConfig("TestData", UpdateBelieverUpgrade)
}

func Close() {

}

var (
	believerUpgradeData map[int32]*data.BelieverUpgrade
)

func UpdateBelieverUpgrade(content []byte) {
	var dataArray []*data.BelieverUpgrade
	json.Unmarshal(content, &dataArray)
	results := make(map[int32]*data.BelieverUpgrade)
	for _, data := range dataArray {
		results[data.Id] = data
	}
	believerUpgradeData = results
}

func GetBelieverUpgrade(id int32) *data.BelieverUpgrade {
	if believerUpgradeData == nil {
		exception.GameException(protocol.Code_ConfigException)
	}
	return believerUpgradeData[id]
}

