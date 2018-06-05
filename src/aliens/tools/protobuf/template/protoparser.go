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
	"os"
	"aliens/tools/protobuf/proto"
	"aliens/common/util"
	"strings"
)

const (
	REQUEST_TAG = "request"
	RESPONSE_TAG = "response"
	PUSH_TAG = "push"
)

var message = &ProtoMessage{}

func ParseProto(protoPath string) *ProtoMessage {
	message = &ProtoMessage{Handlers:make( map[int]*ProtoHandler)}
	//"/Users/hejialin/git/server/kylin/src/aliens/protocol/scene/protocol.proto"
	reader, _ := os.Open(protoPath)
	defer reader.Close()

	parser := proto.NewParser(reader)
	definition, _ := parser.Parse()

	for _, element := range definition.Elements {
		switch element.(type) {
			case *proto.Package :
				message.PackageName = element.(*proto.Package).Name
				//log.Println()
				break;
			case *proto.Message:
				comment := element.(*proto.Message).Doc()
				if comment == nil {
					break
				}
				tag := strings.TrimSpace(comment.Message())
				if tag == REQUEST_TAG {
					message.RequestName = element.(*proto.Message).Name
				} else if tag == RESPONSE_TAG {
					message.ResponseName = element.(*proto.Message).Name
				} else if tag == PUSH_TAG {
					message.PushName = element.(*proto.Message).Name
				}
		}
	}
	proto.Walk(definition,
		proto.WithOneof(handleMessage))
	return message

}

func handleMessage(m *proto.Oneof) {
	for _, visitee := range m.Elements {
		field, ok := visitee.(*proto.OneOfField)
		if !ok {
			continue
		}
		handler := message.Handlers[field.Sequence]
		if handler == nil {
			handler = &ProtoHandler{}
			message.Handlers[field.Sequence] = handler
		}
		if field.Doc() != nil {
			handler.Desc = field.Doc().Message()
		}
		if handler.ORequest == "" {
			handler.ORequest = util.FirstToUpper(field.Name)
		} else {
			handler.OResponse = util.FirstToUpper(field.Name)
		}
	}

	//m.
	//proto.Walk(m, proto.WithOneof(handleHandle))

}



