/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/12/5
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package mmo

import (
	"aliens/aliensbot/mmo/core"
	"aliens/aliensbot/mmo/unit"
	"errors"
	"fmt"
	"github.com/vmihailenco/msgpack"
)

type RemoteSender interface {

	CallRemote(id EntityID, method string, args [][]byte) error

}

type PlayerClient interface {

	//CallRemote(id EntityID, method string, args []string)

}

var sender RemoteSender = nil

func RegisterRemoteSender(newSender RemoteSender) {
	sender = newSender
}

type Space = core.Space

type Entity = core.Entity

type EntityID = core.EntityID

type EntityTimerID = core.EntityTimerID

type EntityType = core.EntityType

func RegisterSpace(spacePtr core.ISpace) {
	core.EntityManager.RegisterEntity(spacePtr)
}

func RegisterEntity(entity core.IEntity) {
	core.EntityManager.RegisterEntity(entity)
}

func CreateSpace(eType EntityType, id EntityID) (*Space, error) {
	e, err := core.EntityManager.CreateEntity(eType, nil, unit.EmptyVector, id)
	return e.AsSpace(), err
}

func CreateEntity(eType EntityType, space *Space, pos unit.Vector) (*Entity, error) {
	return core.EntityManager.CreateEntity(eType, space, pos, "")
}


// GetSpace gets the space by ID
func GetSpace(id EntityID) *Space {
	return core.SpaceManager.GetSpace(id)
}


//实体登录到场景
func EnterSpace(spaceID EntityID, eType EntityType, entityID EntityID, pos unit.Vector) (*Entity, error) {
	space := GetSpace(spaceID)
	if space == nil {
		return nil, errors.New(fmt.Sprintf("space %v not found ", spaceID))
	}
	entity := core.EntityManager.GetEntity(entityID)
	//实体已经存在
	if entity != nil {
		return entity, nil
	}
	return space.CreateEntity(eType, pos, entityID)
}


//handle
func HandlerRemoteEntityCall(caller EntityID, id EntityID, method string, args[][]byte) (*Entity, error) {
	return core.EntityManager.HandleRemoteEntityCall(caller, id, method, args)
}

//call entity method
func Call(id EntityID, method string, args ...interface{}) error {
	entity, err := core.EntityManager.HandleLocalEntityCall(id, method, args)
	//本地不存在、调用远程对象
	if entity == nil {
		return callRemote(id, method, args)
	}
	return err
}

func callRemote(id EntityID, method string, args []interface{}) error {
	argsData := make([][]byte, len(args))
	for i, arg := range args {
		data, err := msgpack.Marshal(arg)
		if err != nil {
			return err
		}
		argsData[i] = data
	}
	return sender.CallRemote(id, method, argsData)
}





