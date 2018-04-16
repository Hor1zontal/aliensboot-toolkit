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
	"fmt"
	"io/ioutil"
	"os"
)

const (
	SPLIT_STR = "<message>"
)


func Convert(protoPath string, templatePath string, outputPath string, filePrefix string, overwrite bool) {
	message := ParseProto(protoPath)
	b, err := ioutil.ReadFile(templatePath)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	templateContent := string(b)

	results := strings.Split(templateContent, SPLIT_STR)
	if len(results) < 3 {
		fmt.Errorf("invalid template")
		return
	}
	header := replaceMessage(results[0], message)

	tailf := ""
	if len(results) == 3 {
		tailf = replaceMessage(results[2], message)
	}

	content := ""
	for _, handler := range message.Handlers {
		handleStr := replaceMessage(results[1], message)
		if !handler.IsValid() {
			continue
		}
		handleStr = replaceHandle(handleStr, handler)
		if filePrefix != "" {
			filePath := outputPath + "/" + strings.Replace(filePrefix, "${}", strings.ToLower(handler.ORequest), -1)
			//单独写文件
			writeFile(filePath, header + handleStr + tailf, overwrite)
		} else {
			content += handleStr
		}
	}

	if filePrefix == "" {
		//写一个文件
		writeFile(outputPath, header + content + tailf, overwrite)
	}

}

func writeFile(filePath string, content string, overwrite bool) {
	if !overwrite {
		//文件存在不允许覆盖
		_, err := os.Stat(filePath)
		if err == nil {
			fmt.Println("file " + filePath + " alread exist, skip it!")
			return
		}
	}
	f, err := os.Create(filePath) //创建文件
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	defer f.Close()
	_, err1 := f.Write([]byte(content)) //写入文件(字节数组)
	if err1 != nil {
		fmt.Errorf(err1.Error())
		return
	}
	fmt.Println("gen code file " + filePath + " success!")
}


func replaceMessage(content string, message *ProtoMessage) string {
	content = strings.Replace(content, "${package}", message.PackageName, -1)
	content = strings.Replace(content, "${request}", message.RequestName, -1)
	content = strings.Replace(content, "${response}", message.ResponseName, -1)
	return content
}

func replaceHandle(content string, handler *ProtoHandler) string {
	content = strings.Replace(content, "${o_desc}", handler.Desc, -1)
	content = strings.Replace(content, "${o_request}", handler.ORequest, -1)
	content = strings.Replace(content, "${o_response}", handler.OResponse, -1)
	return content
}