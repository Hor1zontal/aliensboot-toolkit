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


type spaceManager struct {
	spaces map[int32]*entity.Space
}

func newSpaceManager() *spaceManager {
	return &spaceManager{
		spaces: make(map[int32]*entity.Space),
	}
}

//create new space
func (spmgr *spaceManager) CreateSpace(spaceProxy entity.ISpace, minX, maxX, minY, maxY float32, towerRange float32) {
	space := &entity.Space{}
	space.Init(spaceProxy, minX, maxX, minY, maxY, towerRange)
	spmgr.spaces[space.GetID()] = space
}

//release exist space
func (spmgr *spaceManager) ReleaseSpace(id int32) {
	delete(spmgr.spaces, id)
}

//get space
func (spmgr *spaceManager) getSpace(id int32) *entity.Space {
	return spmgr.spaces[id]
}

func (spmgr *spaceManager) CreateEntity(spaceID int32, proxy entity.IEntity, position entity.Vector3, direction entity.Vector3) *entity.Entity {
	space := spmgr.getSpace(spaceID)
	if space == nil {
		log.Warnf("space is not found %v", spaceID)
		return nil
	}
	entity := &entity.Entity{}
	entity.Init(space, proxy, position, direction)
	space.EntityEnter(entity, position)
	return entity
}

func (spmgr *spaceManager) LeaveEntity(spaceID int32, entityID int32)  {
	space := spmgr.getSpace(spaceID)
	if space == nil {
		log.Warnf("space is not found %v", spaceID)
		return
	}
	space.EntityLeave(entityID)
}

func (spmgr *spaceManager) MoveEntity(spaceID int32, entityID int32, position entity.Vector3, direction entity.Vector3)  {
	space := spmgr.getSpace(spaceID)
	if space == nil {
		log.Warnf("space is not found %v", spaceID)
		return
	}
	space.EntityMove(entityID, position, direction)
}

func (spmgr *spaceManager) GetEntityState(spaceID int32, entityID int32) entity.EntitySet {
	space := spmgr.getSpace(spaceID)
	if space == nil {
		log.Warnf("space is not found %v", spaceID)
		return nil
	}
	return space.GetNeighbors(entityID)
}