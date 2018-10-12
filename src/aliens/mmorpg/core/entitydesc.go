/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/8/31
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core

import (
	"reflect"
	"strings"
	"aliens/common/data_structures/set"
)

var _VALID_ATTR_DEFS = set.StringSet{} // all valid attribute defs



func init() {
	_VALID_ATTR_DEFS.Add(AttrTagFeatureSelf) //
	_VALID_ATTR_DEFS.Add(AttrTagFeatureAll)
	_VALID_ATTR_DEFS.Add(AttrTagFeaturePersist)
}

const (
	rfServer      = 1 << iota
	rfOwnClient   = 1 << iota
	rfOtherClient = 1 << iota

	AttrTagFeature = "feature"

	AttrTagFeatureSelf    = "self"
	AttrTagFeatureAll     = "all"
	AttrTagFeaturePersist = "persist"


)

type methodDesc struct {
	Func       reflect.Value
	Flags      uint
	MethodType reflect.Type
	NumArgs    int
}

type methodDescMap map[string]*methodDesc

func (rdm methodDescMap) visit(method reflect.Method) {
	methodName := method.Name
	var flag uint
	var rpcName string
	if strings.HasSuffix(methodName, "_Client") {
		flag |= rfServer + rfOwnClient
		rpcName = methodName[:len(methodName)-7]
	} else if strings.HasSuffix(methodName, "_AllClients") {
		flag |= rfServer + rfOwnClient + rfOtherClient
		rpcName = methodName[:len(methodName)-11]
	} else {
		// server method
		flag |= rfServer
		rpcName = methodName
	}
	methodType := method.Type
	rdm[rpcName] = &methodDesc{
		Func:       method.Func,
		Flags:      flag,
		MethodType: methodType,
		NumArgs:    methodType.NumIn() - 1, // do not count the receiver
	}
}


// EntityTypeDesc is the entity type description for registering entity types
type EntityDesc struct {
	name string //entity type name

	client bool //is client exist this entity

	useAOI          bool
	aoiDistance     float32

	entityType      reflect.Type

	methodDescs     methodDescMap

	selfAttrs set.StringSet

	allAttrs set.StringSet

	persistAttrs set.StringSet
}


func (desc *EntityDesc) IsPersistent() bool {
	return !desc.persistAttrs.IsEmpty()
}