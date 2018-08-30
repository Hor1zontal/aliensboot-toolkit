/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/3/23
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core

import (
	"aliens/protocol"
	"errors"
	"fmt"
)

var (
	SpaceManager = newSpaceManager()
)

type spaceManager struct {

	spaces map[SpaceID]*Space //

	entities map[EntityID]SpaceID //实体所在的场景

}

func newSpaceManager() *spaceManager {
	return &spaceManager{
		spaces: make(map[SpaceID]*Space),
	}
}

//create new space
func (spmgr *spaceManager) CreateSpace(spaceProxy ISpace) {
	space := &Space{}
	space.Init(spaceProxy, spaceProxy.GetConfig())
	spmgr.spaces[space.GetID()] = space
}

//release exist space
func (spmgr *spaceManager) ReleaseSpace(id SpaceID) {
	delete(spmgr.spaces, id)
}

func (spmgr *spaceManager) getEntitySpace(id EntityID) (*Space, error) {
	spaceID, ok := spmgr.entities[id]
	if !ok {
		return nil, errors.New(fmt.Sprintf("entity %v is not found in any space", id))
	}
	return spmgr.getSpace(spaceID)
}

//get space
func (spmgr *spaceManager) getSpace(id SpaceID) (*Space, error) {
	space, ok := spmgr.spaces[id]
	if !ok {
		return space, errors.New(fmt.Sprintf("space %v is not found ", id))
	}
	return  space, nil
}

func (spmgr *spaceManager) CreateEntity(spaceID SpaceID, proxy IEntity, position *protocol.Vector, direction *protocol.Vector) (*Entity, error) {
	space, err := spmgr.getSpace(spaceID)
	if err != nil {
		return nil, err
	}
	entity := &Entity{}
	entity.Init(space, proxy, position, direction)
	space.EntityEnter(entity, position)
	return entity, nil
}

func (spmgr *spaceManager) LeaveEntity(entityID EntityID) error {
	space, err := spmgr.getEntitySpace(entityID)
	if err != nil {
		return err
	}
	space.EntityLeave(entityID)
	return nil
}

func (spmgr *spaceManager) MoveEntity(entityID EntityID, position *protocol.Vector, direction *protocol.Vector) error {
	space, err := spmgr.getEntitySpace(entityID)
	if err != nil {
		return err
	}

	entity := space.EntityMove(entityID, position, direction)
	if entity == nil {
		return errors.New(fmt.Sprintf("entity %v in not found in space %v", entityID, space.GetID()))
	}
	return nil
}

func (spmgr *spaceManager) GetEntityState(entityID EntityID) (EntitySet, error) {
	space, err := spmgr.getEntitySpace(entityID)
	if err != nil {
		return nil, err
	}
	return space.GetNeighbors(entityID), nil
}