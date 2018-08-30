/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/3/30
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package main

import (
	"aliens/tools/protobuf/template"
	"flag"
	"fmt"
	"io/ioutil"
	"encoding/json"
)


var (
	proto    string
	templatePath string
	output    string
	prefix      string
	overwrite    bool

	configPath string
)

func loadConfig(config interface{}, path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("module %v config file is not found %v", path)
	}
	err = json.Unmarshal(data, config)
	if err != nil {
		fmt.Printf("load config %v err %v", path, err)
	}
}

func main() {
	//flag.StringVar(&proto, "proto", "", "protobuf file path")
	//flag.StringVar(&templatePath,"template", "", "output template path")
	//flag.StringVar(&output,"output", "", "output path")
	//flag.StringVar(&prefix,"prefix", "", "output file prefix")
	//flag.BoolVar(&overwrite,"overwrite", false, "is overwrite ?")

	flag.StringVar(&configPath,"configPath", "", "config file path")
	flag.Parse()


	if configPath == "" {
		fmt.Printf("Please input correct params => codegen -h \n")
		return
	}

	//fmt.Printf("proto: %v template: %v output: %v prefix: %v overwrite: %v", proto, templatePath, output, prefix, overwrite)
	//if proto == "" || templatePath == "" || output == "" {
	//	fmt.Printf("Please input correct params => codegen -h \n")
	//	return
	//}

	configObject := &template.Config{}
	loadConfig(configObject, configPath) //加载服务器配置

	template.Convert(configObject)

	//template.Convert(proto, templatePath, output, prefix, overwrite)
	//genSceneProto()
	//genPassportProto()
}




