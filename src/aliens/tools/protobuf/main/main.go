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

	protoPath := "/Users/hejialin/git/server/kylin/src/aliens/protocol/scene/protocol.proto"
	templatePath := "/Users/hejialin/git/server/kylin/src/aliens/protocol/scene/scene.template"
	outputPath := "/Users/hejialin/git/server/kylin/src/aliens/module/scene/service/handle.go"
	template.Convert(protoPath, templatePath, outputPath, "", true)

	handleTemplatePath := "/Users/hejialin/git/server/kylin/src/aliens/protocol/scene/scene_handle.template"
	outputDir := "/Users/hejialin/git/server/kylin/src/aliens/module/scene/service/"
	filePrefix := "handle_${}.go"
	template.Convert(protoPath, handleTemplatePath, outputDir, filePrefix, false)
}


