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
	"aliens/log"
	"path/filepath"
	"github.com/go-yaml/yaml"
	"fmt"
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

type BaseConfig struct {
	ClusterName  string    `yaml:"cluster.name"`
	ClusterCenter []string `yaml:"cluster.center"`
	ClusterTTL int64 	   `yaml:"cluster.ttl"`

	NodeName  string       `yaml:"node.name"`

	PathLog string    	   `yaml:"path.log"`



}

func Init(configPath string) *BaseConfig {
	dir, _ := filepath.Abs(filepath.Base(configPath))
	//log.Debugf("configuration path is %v", dir)
	ModuleConfigRoot = dir + string(filepath.Separator) + "modules" + string(filepath.Separator)

	config := &BaseConfig{}
	baseConfigPath := dir + string(filepath.Separator) + "aliensbot.yml"
	LoadConfigData(baseConfigPath, config)
	fmt.Println("%v", config)
	return config
}


//func LoadConfigData(name string, config interface{}) {
//	if config == nil {
//		return
//	}
//	path := ModuleConfigRoot + name + ".yml"
//	data, err := ioutil.ReadFile(path)
//	if err != nil {
//		log.Fatalf("module %v config file is not found %v", name, path)
//	}
//	err = json.Unmarshal(data, config)
//	if err != nil {
//		log.Fatalf("load config %v err %v", path, err)
//	}
//}

func LoadConfigData(path string, config interface{}) {
	if config == nil {
		return
	}
	ymlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("config file is not found %v", path)
	}
	err = yaml.Unmarshal(ymlFile, config)
	if err != nil {
		log.Fatalf("load config %v err %v", path, err)
	}
}

func LoadModuleConfigData(name string, config interface{}) {
	if config == nil {
		return
	}
	path := ModuleConfigRoot + name + ".yml"

	LoadConfigData(path, config)
}
