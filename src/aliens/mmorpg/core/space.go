package core

import (
	"aliens/mmorpg/aoi"
	"fmt"
	"aliens/protocol"
	"aliens/mmorpg/config"
)

type SpaceID uint32

// ISpace is the space delegate interface
type ISpace interface {
	// space Operations
	//GetID() SpaceID

	GetConfig() config.SpaceConfig //

	// Called when any entity enters space
	OnEntityEnter(entity *Entity)

	// Called when any entity leaves space
	OnEntityLeave(entity *Entity)
}

type Space struct {

	id SpaceID

	entities EntityMap  //entities in current space

	aoiMgr aoi.Manager


	proxy ISpace
}

func (space *Space) GetID() SpaceID {
	return space.id
}

func (space *Space) String() string {
	return fmt.Sprintf("space<%d>", space.GetID())
}

// Init initialize space entity
func (space *Space) Init(proxy ISpace, spaceConfig config.SpaceConfig) {
	//TODO 通过地形数据生成对象
	space.entities = EntityMap{}
	space.proxy = proxy
	space.aoiMgr = aoi.NewTowerAOIManager(spaceConfig.MinX, spaceConfig.MaxX, spaceConfig.MinY, spaceConfig.MaxY, spaceConfig.TowerRange)
}

func (space *Space) EntityEnter(entity *Entity, pos *protocol.Vector) {
	entity.spaceId = space.id
	space.entities.Add(entity)
	space.aoiMgr.Enter(entity.aoi, pos.X, pos.Y)
	space.proxy.OnEntityEnter(entity)
}

func (space *Space) EntityLeave(entityID EntityID) *Entity {
	entity := space.entities.Get(entityID)
	if entity == nil {
		return nil
	}
	// remove from space entities
	entity.spaceId = 0
	space.entities.Del(entity.GetID())
	space.aoiMgr.Leave(entity.aoi)
	space.proxy.OnEntityLeave(entity)
	return entity
}

func (space *Space) EntityMove(entityID EntityID, newPos *protocol.Vector, direction *protocol.Vector) *Entity {
	entity := space.entities.Get(entityID)
	if entity == nil {
		return nil
	}
	entity.position = newPos
	entity.direction = direction
	space.aoiMgr.Moved(entity.aoi, newPos.X, newPos.Y)

	for neighbor, _ := range entity.neighbors {
		neighbor.proxy.OnEntityMove(entity)
	}
	return entity
}


func (space *Space) GetNeighbors(entityID EntityID) EntitySet {
	entity := space.entities.Get(entityID)
	if entity == nil {
		return nil
	}
	return entity.neighbors
}


// GetEntityCount returns the total count of entities in space
func (space *Space) GetEntityCount() int {
	return len(space.entities)
}

// ForEachEntity visits all entities in space and call function f with each entity
func (space *Space) ForEachEntity(f func(e *Entity)) {
	for _, entity := range space.entities {
		f(entity)
	}
}

// GetEntity returns the entity in space with specified id, nil otherwise
func (space *Space) GetEntity(entityID EntityID) *Entity {
	return  space.entities.Get(entityID)
}

