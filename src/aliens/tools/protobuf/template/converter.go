/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/3/30
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package template

import (
	"strings"
	"io/ioutil"
	"os"
	"fmt"
)

const (
	SPLIT_STR = "<message>"
)


func Convert(config *Config) {
	message := ParseProto(config.ProtoPath)

	//fmt.Printf("proto data %v", message.modules["passport"].Handlers[6])

	for _, moduleConfig := range config.Modules {
		module := message.modules[moduleConfig.Name]
		if module == nil {
			fmt.Printf("module %v is nou found in proto file %v \n", moduleConfig.Name, config.ProtoPath)
			continue
		}

		convertModule(moduleConfig, module)
	}


}


func convertModule(moduleConfig *ModuleConfig, module *Module) {
	for _, outputConfig := range moduleConfig.Outputs {
		b, err := ioutil.ReadFile(outputConfig.Template)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
		templateContent := string(b)

		results := strings.Split(templateContent, SPLIT_STR)

		header := ""
		content := ""
		tailf := ""

		if len(results) == 3 {
			header = replaceMessage(results[0], module)
			tailf = replaceMessage(results[2], module)


			for _, handler := range module.Handlers {
				handleStr := replaceMessage(results[1], module)
				if !handler.IsValid() {
					continue
				}
				handleStr = replaceHandle(handleStr, handler)
				if outputConfig.Prefix != "" {
					filePath := outputConfig.Output + "/" + strings.Replace(outputConfig.Prefix, "${}", strings.ToLower(handler.ORequest), -1)
					//单独写文件
					writeFile(filePath, header + handleStr + tailf, outputConfig.Overwrite)
				} else {
					content += handleStr
				}
			}
		} else {
			header = replaceMessage(templateContent, module)
		}
		if  outputConfig.Prefix == "" {
			//写一个文件
			writeFile( outputConfig.Output, header + content + tailf,  outputConfig.Overwrite)
		}
	}



}

func writeFile(filePath string, content string, overwrite bool) {
	if !overwrite {
		//文件存在不允许覆盖
		_, err := os.Stat(filePath)
		if err == nil {
			fmt.Printf("file " + filePath + " alread exist, skip it! \n")
			return
		}
	}
	f, err := os.Create(filePath) //创建文件
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer f.Close()
	_, err1 := f.Write([]byte(content)) //写入文件(字节数组)
	if err1 != nil {
		fmt.Printf(err1.Error())
		return
	}
	fmt.Printf("gen code file " + filePath + " success! \n")
}


func replaceMessage(content string, module *Module) string {
	content = strings.Replace(content, "${package}", message.PackageName, -1)
	content = strings.Replace(content, "${module}", module.Name, -1)
	content = strings.Replace(content, "${Module}", module.UName, -1)
	content = strings.Replace(content, "${request}", module.Name, -1)
	content = strings.Replace(content, "${response}", module.Name, -1)
	return content
}

func replaceHandle(content string, handler *ProtoHandler) string {
	content = strings.Replace(content, "${o_desc}", handler.Desc, -1)
	content = strings.Replace(content, "${o_request}", handler.ORequest, -1)
	content = strings.Replace(content, "${o_response}", handler.OResponse, -1)
	return content
}