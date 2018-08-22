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
	"path/filepath"
)


var (
	Version = "1.0.0"

	LenStackBuf = 4096

	//// log
	//LogLevel string
	//LogPath  string
	//LogFlag  int
	ModuleConfigRoot = ""

	// console
	ConsolePort   int
	ConsolePrompt string = "AliensBot# "
	ProfilePath   string
)

func Init(configPath string) {
	dir, _ := filepath.Abs(filepath.Base(configPath))
	//log.Debugf("configuration path is %v", dir)
	ModuleConfigRoot = dir + string(filepath.Separator) + "modules" + string(filepath.Separator)
}


func LoadConfigData(name string, config interface{}) {
	if config == nil {
		return
	}
	path := ModuleConfigRoot + name + ".json"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("module %v config file is not found %v", name, path)
	}
	err = json.Unmarshal(data, config)
	if err != nil {
		log.Fatalf("load config %v err %v", path, err)
	}
}
