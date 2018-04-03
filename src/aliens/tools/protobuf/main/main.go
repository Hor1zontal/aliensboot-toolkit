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
)

func main() {
	genSceneProto()
	//genPassportProto()
}

func genSceneProto() {
	protoPath := "/Users/hejialin/git/server/kylin/src/aliens/protocol/scene/protocol.proto"
	templatePath := "/Users/hejialin/git/server/kylin/src/aliens/protocol/scene/scene.template"
	outputPath := "/Users/hejialin/git/server/kylin/src/aliens/module/scene/service/handle.go"
	template.Convert(protoPath, templatePath, outputPath, "", true)

	protoPath1 := "/Users/hejialin/git/server/kylin/src/aliens/protocol/scene/protocol.proto"
	templatePath1 := "/Users/hejialin/git/server/kylin/src/aliens/protocol/scene/scene_json.template"
	outputPath1 := "/Users/hejialin/git/server/kylin/src/aliens/module/scene/service/handlejson.go"
	template.Convert(protoPath1, templatePath1, outputPath1, "", true)

	handleTemplatePath := "/Users/hejialin/git/server/kylin/src/aliens/protocol/scene/scene_handle.template"
	outputDir := "/Users/hejialin/git/server/kylin/src/aliens/module/scene/service/"
	filePrefix := "handle_${}.go"
	template.Convert(protoPath, handleTemplatePath, outputDir, filePrefix, false)
}

func genPassportProto() {
	protoPath := "/Users/hejialin/git/server/kylin/src/aliens/protocol/passport/protocol.proto"
	templatePath := "/Users/hejialin/git/server/kylin/src/aliens/protocol/passport/passport.template"
	outputPath := "/Users/hejialin/git/server/kylin/src/aliens/module/passport/service/handle.go"
	template.Convert(protoPath, templatePath, outputPath, "", true)

	handleTemplatePath := "/Users/hejialin/git/server/kylin/src/aliens/protocol/passport/passport_handle.template"
	outputDir := "/Users/hejialin/git/server/kylin/src/aliens/module/passport/service/"
	filePrefix := "handle_${}.go"
	template.Convert(protoPath, handleTemplatePath, outputDir, filePrefix, false)
}



