package core

import (
	"aliens/aliensbot/log"
	"aliens/aliensbot/mmorpg/aoi"
	"aliens/aliensbot/mmorpg/config"
	"aliens/aliensbot/protocol"
	"fmt"
	"github.com/gogo/protobuf/proto"
)

type ISpace interface {
	IEntity
	//GetConfig() config.SpaceConfig //

	GetMeta() proto.Message

	// Called when any entity enters space
	OnEntityEnter(entity *Entity)

	// Called when any entity leaves space
	OnEntityLeave(entity *Entity)

	// Called when any entity leaves space
	OnEntityMove(entity *Entity)
}

type Space struct {
	id EntityID

	entities EntityMap //entities in current space

	aoiMgr aoi.Manager

	proxy ISpace
}

func (space *Space) GetID() EntityID {
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

//修改视野范围
//func (space *Space) EntityChangeViewRadius(entity *Entity, radius float32) {
//	space.aoiMgr.ChangeViewRadius(entity.aoi, radius)
//}

//进入场景
func (space *Space) enter(entity *Entity, pos *protocol.Vector) {
	if entity.space != nil {
		log.Panicf("%s.enter(%s): current space is not nil, but %s", space, entity, entity.space)
	}

	entity.space = space
	space.entities.Add(entity)
	space.aoiMgr.Enter(entity.aoi, pos.X, pos.Y)
	space.proxy.OnEntityEnter(entity)
}

//离开场景
func (space *Space) leave(entity *Entity) {
	if entity.space != space {
		log.Panicf("%s.leave(%s): entity is not in this Space", space, entity)
	}

	entity.space = nil
	space.entities.Del(entity.GetID())
	space.aoiMgr.Leave(entity.aoi)
	space.proxy.OnEntityLeave(entity)
}

//场景中移动
func (space *Space) move(entity *Entity, newPos *protocol.Vector) {
	if entity.space != space {
		log.Panicf("%s.leave(%s): entity is not in this Space", space, entity)
	}

	entity.Position = newPos
	space.aoiMgr.Moved(entity.aoi, newPos.X, newPos.Y)
	//for neighbor, _ := range entity.interestedIn {
	//	neighbor.proxy.OnEntityMove(entity)
	//}
	//space.proxy.OnEntityMove(entity)
}

// CreateEntity creates a new local entity in this space
func (space *Space) CreateEntity(typeName string, pos *protocol.Vector) {
	EntityManager.CreateLocalEntity(typeName, space, pos)
}

func (space *Space) GetNeighbors(entityID EntityID) EntitySet {
	entity := space.entities.Get(entityID)
	if entity == nil {
		return nil
	}
	return entity.interestedIn
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
	return space.entities.Get(entityID)
}
