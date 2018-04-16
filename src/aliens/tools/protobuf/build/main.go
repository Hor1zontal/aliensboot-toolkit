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
)


var (
	proto    string
	templatePath string
	output    string
	prefix      string
	overwrite    bool
)

func main() {
	flag.StringVar(&proto, "proto", "", "protobuf file path")
	flag.StringVar(&templatePath,"template", "", "output template path")
	flag.StringVar(&output,"output", "", "output path")
	flag.StringVar(&prefix,"prefix", "", "output file prefix")
	flag.BoolVar(&overwrite,"overwrite", false, "is overwrite ?")


	flag.Parse()

	//fmt.Printf("proto: %v template: %v output: %v prefix: %v overwrite: %v", proto, templatePath, output, prefix, overwrite)
	if proto == "" || templatePath == "" || output == "" {
		fmt.Printf("Please input correct params => codegen -h \n")
		return
	}
	template.Convert(proto, templatePath, output, prefix, overwrite)
	//genSceneProto()
	//genPassportProto()
}

func genSceneProto() {
	protoPath := "/Users/hejialin/git/server/kylin/src/aliens/protocol/scene/protocol.proto"
	templatePath := "/Users/hejialin/git/server/kylin/src/aliens/protocol/common.template"
	outputPath := "/Users/hejialin/git/server/kylin/src/aliens/module/scene/service/handle.go"
	template.Convert(protoPath, templatePath, outputPath, "", true)

	protoPath1 := "/Users/hejialin/git/server/kylin/src/aliens/protocol/scene/protocol.proto"
	templatePath1 := "/Users/hejialin/git/server/kylin/src/aliens/protocol/common_json.template"
	outputPath1 := "/Users/hejialin/git/server/kylin/src/aliens/module/scene/service/handlejson.go"
	template.Convert(protoPath1, templatePath1, outputPath1, "", true)

	handleTemplatePath := "/Users/hejialin/git/server/kylin/src/aliens/protocol/common_handle.template"
	outputDir := "/Users/hejialin/git/server/kylin/src/aliens/module/scene/service/"
	filePrefix := "handle_${}.go"
	template.Convert(protoPath, handleTemplatePath, outputDir, filePrefix, false)
}

func genPassportProto() {
	protoPath := "/Users/hejialin/git/server/kylin/src/aliens/protocol/passport/protocol.proto"
	templatePath := "/Users/hejialin/git/server/kylin/src/aliens/protocol/common.template"
	outputPath := "/Users/hejialin/git/server/kylin/src/aliens/module/passport/service/handle.go"
	template.Convert(protoPath, templatePath, outputPath, "", true)


	protoPath1 := "/Users/hejialin/git/server/kylin/src/aliens/protocol/passport/protocol.proto"
	templatePath1 := "/Users/hejialin/git/server/kylin/src/aliens/protocol/common_json.template"
	outputPath1 := "/Users/hejialin/git/server/kylin/src/aliens/module/passport/service/handlejson.go"
	template.Convert(protoPath1, templatePath1, outputPath1, "", true)

	handleTemplatePath := "/Users/hejialin/git/server/kylin/src/aliens/protocol/common_handle.template"
	outputDir := "/Users/hejialin/git/server/kylin/src/aliens/module/passport/service/"
	filePrefix := "handle_${}.go"
	template.Convert(protoPath, handleTemplatePath, outputDir, filePrefix, false)
}



