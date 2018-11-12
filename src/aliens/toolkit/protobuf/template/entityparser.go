/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/9/1
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package template

import (
	//"aliens/log"
	"aliens/toolkit/protobuf/proto"
	"fmt"
	"os"
)

func ParseEntityProto(protoPath string) {
	//message = &ServiceMessage{modules:make(map[string]*Module)}
	//"/Users/hejialin/git/server/kylin/src/aliens/protocol/scene/protocol.proto"
	reader, _ := os.Open(protoPath)
	defer reader.Close()

	parser := proto.NewParser(reader)
	definition, _ := parser.Parse()

	for _, element := range definition.Elements {
		switch element.(type) {
		case *proto.Package:
			message.PackageName = element.(*proto.Package).Name
			//log.Println()
			break
		case *proto.Message:
			messageData := element.(*proto.Message)
			fmt.Sprintf("%v", messageData)
			//proto.Walk(, proto.WithOneof(&messageWalk{tag:tag}.handleMessage))

			//if tag == REQUEST_TAG {
			//	//message.RequestName = element.(*proto.Message).Name
			//} else if tag == RESPONSE_TAG  {
			//	//message.ResponseName = element.(*proto.Message).Name
			//} else if tag == PUSH_TAG{
			//	//message.PushName = element.(*proto.Message).Name
			//}
		}
	}

}
