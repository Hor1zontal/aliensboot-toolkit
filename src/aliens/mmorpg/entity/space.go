package entity

import (
	"aliens/mmorpg/aoi"
	"fmt"
)

// ISpace is the space delegate interface
type ISpace interface {
	// space Operations
	GetID() int32

	OnEntityEnter(entity *Entity) // Called when any entity enters space
	OnEntityLeave(entity *Entity) // Called when any entity leaves space
}

type Space struct {
	entities EntityMap  //entities in current space
	aoiMgr aoi.AOIManager
	proxy ISpace
}

func (space *Space) GetID() int32 {
	return space.proxy.GetID()
}

func (space *Space) String() string {
	return fmt.Sprintf("space<%d>", space.GetID())
}

// Init initialize space entity
func (space *Space) Init(proxy ISpace, minX, maxX, minY, maxY float32, towerRange float32) {
	space.entities = EntityMap{}
	space.proxy = proxy
	space.aoiMgr = aoi.NewTowerAOIManager(minX, maxX, minY, maxY, towerRange)
}

func (space *Space) EntityEnter(entity *Entity, pos Vector3) {
	entity.space = space
	space.entities.Add(entity)
	space.proxy.OnEntityEnter(entity)
	space.aoiMgr.Enter(entity.aoi, pos.X, pos.Y)
}

func (space *Space) EntityLeave(entityID int32) *Entity {
	entity := space.entities.Get(entityID)
	if entity == nil {
		return nil
	}
	// remove from space entities
	entity.space = nil
	space.entities.Del(entity.GetID())
	space.proxy.OnEntityLeave(entity)
	space.aoiMgr.Leave(entity.aoi)
	return entity
}

func (space *Space) EntityMove(entityID int32, newPos Vector3, direction Vector3) *Entity {
	entity := space.entities.Get(entityID)
	if entity == nil {
		return nil
	}
	entity.position = newPos
	entity.direction = direction
	space.aoiMgr.Moved(entity.aoi, newPos.X, newPos.Y)
	return entity
}


func (space *Space) GetNeighbors(entityID int32) EntitySet {
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
func (space *Space) GetEntity(entityID int32) *Entity {
	return  space.entities.Get(entityID)
}

