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
	center.ClusterCenter.SubscribeConfig("army", UpdateArmyData)
}

func Close() {

}

var (
	armyData map[int32]*data.Army
)

func UpdateArmyData(content []byte) {
	var dataArray []*data.Army
	json.Unmarshal(content, &dataArray)
	results := make(map[int32]*data.Army)
	for _, data := range dataArray {
		results[data.Tid] = data
	}
	armyData = results
}

func GetArmyData(id int32) *data.Army {
	if armyData == nil {
		exception.GameException(protocol.Code_ConfigException)
	}
	return armyData[id]
}

