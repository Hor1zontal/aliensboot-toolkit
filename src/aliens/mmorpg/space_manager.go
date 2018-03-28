/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/3/23
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package mmorpg

import (
	"aliens/log"
	"aliens/mmorpg/entity"
)

var (
	SpaceManager = newSpaceManager()
)


type _SpaceManager struct {
	spaces map[int32]*entity.Space
}

func newSpaceManager() *_SpaceManager {
	return &_SpaceManager{
		spaces: make(map[int32]*entity.Space),
	}
}

//create new space
func (spmgr *_SpaceManager) CreateSpace(spaceProxy entity.ISpace, minX, maxX, minY, maxY float32, towerRange float32) {
	space := &entity.Space{}
	space.Init(spaceProxy, minX, maxX, minY, maxY, towerRange)
	spmgr.spaces[space.GetID()] = space
}

//release exist space
func (spmgr *_SpaceManager) ReleaseSpace(id int32) {
	delete(spmgr.spaces, id)
}

//get space
func (spmgr *_SpaceManager) getSpace(id int32) *entity.Space {
	return spmgr.spaces[id]
}

func (spmgr *_SpaceManager) CreateEntity(spaceID int32, proxy entity.IEntity, position entity.Vector3, direction entity.Vector3) *entity.Entity {
	space := spmgr.getSpace(spaceID)
	if space == nil {
		log.Warn("space is not found %v", spaceID)
		return nil
	}
	entity := &entity.Entity{}
	entity.Init(space, proxy, position, direction)
	space.EntityEnter(entity, position)
	return entity
}

func (spmgr *_SpaceManager) LeaveEntity(spaceID int32, entityID int32)  {
	space := spmgr.getSpace(spaceID)
	if space == nil {
		log.Warn("space is not found %v", spaceID)
		return
	}
	space.EntityLeave(entityID)
}

func (spmgr *_SpaceManager) MoveEntity(spaceID int32, entityID int32, position entity.Vector3, direction entity.Vector3)  {
	space := spmgr.getSpace(spaceID)
	if space == nil {
		log.Warn("space is not found %v", spaceID)
		return
	}
	space.EntityMove(entityID, position, direction)
}

func (spmgr *_SpaceManager) GetEntityState(spaceID int32, entityID int32) entity.EntitySet {
	space := spmgr.getSpace(spaceID)
	if space == nil {
		log.Warn("space is not found %v", spaceID)
		return nil
	}
	return space.GetNeighbors(entityID)
}