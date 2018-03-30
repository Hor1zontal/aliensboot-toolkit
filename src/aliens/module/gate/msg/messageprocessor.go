/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/3/30
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package msg

import (
	"github.com/name5566/leaf/chanrpc"
	"errors"
	"github.com/gogo/protobuf/types"
	"encoding/binary"
	"reflect"
)

type MessageProcessor struct {
	littleEndian bool
	msgRouter     *chanrpc.Server

	urlIDMapping map[string]uint16
	idURLMapping map[uint16]string

}

func NewMsgProcessor() *MessageProcessor {
	return &MessageProcessor{
		urlIDMapping:make(map[string]uint16),
		idURLMapping:make(map[uint16]string),
	}
}

// It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *MessageProcessor) SetByteOrder(littleEndian bool) {
	p.littleEndian = littleEndian
}

func (this *MessageProcessor) Register(id uint16, url string) uint16 {
	this.urlIDMapping[url] = id
	this.idURLMapping[id] = url
	return id
}

func (this *MessageProcessor) Route(msg interface{}, userData interface{}) error {
	this.msgRouter.Go(reflect.TypeOf(&types.Any{}), msg, userData)
	return nil
}

// must goroutine safe
func (this *MessageProcessor) Unmarshal(data []byte) (interface{}, error) {
	if len(data) < 2 {
		return nil, errors.New("data too short")
	}

	var id uint16 = 0
	if this.littleEndian {
		id = binary.LittleEndian.Uint16(data)
	} else {
		id = binary.BigEndian.Uint16(data)
	}
	url, ok := this.idURLMapping[id]
	if !ok {

	}
	return &types.Any{TypeUrl:url, Value:data[2:]}, nil
}

// must goroutine safe
func (this *MessageProcessor) Marshal(msg interface{}) ([][]byte, error) {
	any, ok := msg.(*types.Any)
	if !ok {
		return nil, errors.New("invalid any type")
	}
	msgID := this.urlIDMapping[any.TypeUrl]
	id := make([]byte, 2)
	if this.littleEndian {
		binary.LittleEndian.PutUint16(id, msgID)
	} else {
		binary.BigEndian.PutUint16(id, msgID)
	}
	return [][]byte{id, any.Value}, nil
}
