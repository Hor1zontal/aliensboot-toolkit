/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/3/23
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core

var (
	spaceManager = newSpaceManager()
)

func GetSpace(id EntityID) *Space {
	return spaceManager.spaces[id]
}

type _SpaceManager struct {
	spaces map[EntityID]*Space
}

func newSpaceManager() *_SpaceManager {
	return &_SpaceManager{
		spaces: map[EntityID]*Space{},
	}
}

// CreateSpaceLocally creates a space in the local game server
func CreateSpaceLocally(typeName string) *Space {
	e := EntityManager.createEntity(typeName, nil, nil, "", nil)
	return e.AsSpace()
}

func (spmgr *_SpaceManager) putSpace(space *Space) {
	spmgr.spaces[space.id] = space
}

func (spmgr *_SpaceManager) delSpace(id EntityID) {
	delete(spmgr.spaces, id)
}

func (spmgr *_SpaceManager) getSpace(id EntityID) *Space {
	return spmgr.spaces[id]
}

//release exist space
func (spmgr *_SpaceManager) ReleaseSpace(id EntityID) {
	delete(spmgr.spaces, id)
}
