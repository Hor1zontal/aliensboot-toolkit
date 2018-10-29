/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2017/9/13
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package main

import (
	"aliens/tools/zookeeper"
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	from := *flag.String("from", "", "zookeeper copy from address")
	to := *flag.String("to", "", "zookeeper copy to address")
	fromPath := *flag.String("fromPath", "", "fromPath")
	toPath := *flag.String("toPath", "", "node name")
	flag.Parse()

	zookeeper.Transfer(from, to, fromPath, toPath)

	//zookeeper.Transfer("120.77.213.111:2182", "39.108.220.37:2181", "/gok")

	//result, _ := ip.GetLocalPublicIpUseDnspod()
	//log.Debug(result)
}

func uploadWithConfig() {
	address := flag.String("address", "", "zookeeper address")
	config := flag.String("config", "", "zookeeper file")
	node := flag.String("node", "", "node name")
	flag.Parse()

	f, err := os.Open(*config)
	if err != nil {
		return
	}
	defer f.Close()

	zkConn, err := zookeeper.Connect(*address)
	if err != nil {
		return
	}
	nodePath := "/" + *node
	zookeeper.Create(zkConn, nodePath)
	zookeeper.Create(zkConn, nodePath+"/config")
	zookeeper.Create(zkConn, nodePath+"/service")

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}

		var strLine, name, value string
		strLine = string(line)
		strArr := strings.Split(strLine, "    ")
		if len(strArr) >= 2 {
			name = strArr[0]                     //第一个字段是表名
			value = strings.Join(strArr[1:], "") //第二个以上的字段有可能被拆分多个,就需要连接到一起
		} else {
			fmt.Println("##error##", strArr[:1], "表格解析失败，请检查")
		}

		//fmt.Println("name => " + name)
		//fmt.Println("value => " + value)

		path := nodePath + "/config/" + name
		zookeeper.UpdateByPath(zkConn, path, []byte(value))
	}
}
