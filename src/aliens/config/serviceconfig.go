/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/3/29
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package config

import (
	"io/ioutil"
	"encoding/json"
	"aliens/log"
)

func LoadConfig(config interface{}, path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, config)
	if err != nil {
		log.Fatal("%v", err)
	}
	//log.Debug("json init success %v", config)
}
