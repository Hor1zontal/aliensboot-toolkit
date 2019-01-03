/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/3/30
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package main

import (
	"github.com/KylinHe/aliensboot-toolkit/model"
	"github.com/KylinHe/aliensboot-toolkit/protobuf/template"
	"flag"
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

var (
	//proto        string
	//templatePath string
	//output       string
	//prefix       string
	//overwrite    bool

	configPath string
)

func loadConfig(config interface{}, path string) {
	//data, err := ioutil.ReadFile(path)
	//if err != nil {
	//	fmt.Printf("module %v config file is not found %v", path)
	//}
	//err = json.Unmarshal(data, config)
	//if err != nil {
	//	fmt.Printf("load config %v err %v", path, err)
	//}

	ymlFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("config file is not found %v", path)
	}
	err = yaml.Unmarshal(ymlFile, config)
	if err != nil {
		fmt.Printf("load config %v err %v", path, err)
	}
}

func main() {
	//template.ParseEntityProto("/Users/hejialin/git/server/aliensbot/src/aliens/protocol/entity/player.proto")

	//flag.StringVar(&proto, "proto", "", "protobuf file path")
	//flag.StringVar(&templatePath, "template", "", "output template path")
	//flag.StringVar(&output, "output", "", "output path")
	//flag.StringVar(&prefix, "prefix", "", "output file prefix")
	//flag.BoolVar(&overwrite, "overwrite", false, "is overwrite ?")

	flag.StringVar(&configPath, "configPath", "", "config file path")
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

	configObject := &model.CodeGenConfig{}
	loadConfig(configObject, configPath) //加载服务器配置

	template.Convert(configObject)

	//template.Convert(proto, templatePath, output, prefix, overwrite)
	//genSceneProto()
	//genPassportProto()
}

func genSceneProto() {
	//protoPath := "/Users/hejialin/git/server/kylin/src/aliens/protocol/scene/protocol.proto"
	//templatePath := "/Users/hejialin/git/server/kylin/src/aliens/protocol/common.template"
	//outputPath := "/Users/hejialin/git/server/kylin/src/aliens/module/scene/service/handle.go"
	//template.Convert(protoPath, templatePath, outputPath, "", true)
	//
	//protoPath1 := "/Users/hejialin/git/server/kylin/src/aliens/protocol/scene/protocol.proto"
	//templatePath1 := "/Users/hejialin/git/server/kylin/src/aliens/protocol/common_json.template"
	//outputPath1 := "/Users/hejialin/git/server/kylin/src/aliens/module/scene/service/handlejson.go"
	//template.Convert(protoPath1, templatePath1, outputPath1, "", true)
	//
	//handleTemplatePath := "/Users/hejialin/git/server/kylin/src/aliens/protocol/common_handle.template"
	//outputDir := "/Users/hejialin/git/server/kylin/src/aliens/module/scene/service/"
	//filePrefix := "handle_${}.go"
	//template.Convert(protoPath, handleTemplatePath, outputDir, filePrefix, false)
}

func genPassportProto() {
	//protoPath := "/Users/hejialin/git/server/kylin/src/aliens/protocol/passport/protocol.proto"
	//templatePath := "/Users/hejialin/git/server/kylin/src/aliens/protocol/common.template"
	//outputPath := "/Users/hejialin/git/server/kylin/src/aliens/module/passport/service/handle.go"
	//template.Convert(protoPath, templatePath, outputPath, "", true)
	//
	//
	//protoPath1 := "/Users/hejialin/git/server/kylin/src/aliens/protocol/passport/protocol.proto"
	//templatePath1 := "/Users/hejialin/git/server/kylin/src/aliens/protocol/common_json.template"
	//outputPath1 := "/Users/hejialin/git/server/kylin/src/aliens/module/passport/service/handlejson.go"
	//template.Convert(protoPath1, templatePath1, outputPath1, "", true)
	//
	//handleTemplatePath := "/Users/hejialin/git/server/kylin/src/aliens/protocol/common_handle.template"
	//outputDir := "/Users/hejialin/git/server/kylin/src/aliens/module/passport/service/"
	//filePrefix := "handle_${}.go"
	//template.Convert(protoPath, handleTemplatePath, outputDir, filePrefix, false)

	//protoPath := "/Users/hejialin/git/server/kylin/src/aliens/protocol/passport/protocol.proto"
	//templatePath := "/Users/hejialin/git/server/kylin/src/aliens/protocol/service.template"
	//outputPath := "/Users/hejialin/git/server/kylin/src/aliens/module/passport/service/service1.go"
	//template.Convert(protoPath, templatePath, outputPath, "", true)
}
