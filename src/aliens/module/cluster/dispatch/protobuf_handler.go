/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/4/23
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package dispatch

import (
	"aliens/protocol"
	"github.com/gogo/protobuf/proto"
)


func NewProtobufHandler(proxy func(message *protocol.Any) error) *protobufHandler {
	return &protobufHandler{proxy}
}

type protobufHandler struct {
	proxy func(message *protocol.Any) error
}

func (this *protobufHandler) HandleMessage(data []byte) error {
	requestProxy := &protocol.Any{}
	error := proto.Unmarshal(data, requestProxy)
	if error != nil {
		return error
	}
	return this.proxy(requestProxy)
}