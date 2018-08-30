package core

import (
	"aliens/mmorpg/aoi"
	"fmt"
	"aliens/protocol"
)

type EntityID int64
//var currID int32 = 0
//
//func GenEntityID () int32 {
//	currID++
//	return currID
//}

type IEntity interface {
	GetDist() float32 //可视距离
	OnEntityEnter(entity *Entity) //call it when other entity EntityEnter
	OnEntityLeave(entity *Entity) //call it when other entity EntityLeave
	OnEntityMove(entity *Entity)  //call it when other entity EntityMove
}

type Entity struct {
	//client Client
	//space *Space   //

	id    EntityID

	spaceId SpaceID  //实体所在的空间id

	aoi   *aoi.AOI //one entity , one aoi!

	position  *protocol.Vector   //entity position

	direction *protocol.Vector   //entity direction

	layer int8	//玩家当前的层级

	topSpeed float32 //实体XY轴最大速度

	topSpeedZ float32 //实体Z轴最大的速度

	neighbors EntitySet //aoi entity

	proxy IEntity
}

func (e *Entity) GetID() EntityID {
	return e.id
}

func (e *Entity) GetPosition() *protocol.Vector {
	return e.position
}

func (e *Entity) GetDirection() *protocol.Vector {
	return e.direction
}

func (e *Entity) Init(space *Space, proxy IEntity, position *protocol.Vector, direction *protocol.Vector) {
	e.aoi = aoi.NewAOI(e, proxy.GetDist())
	e.proxy = proxy
	e.neighbors = make(EntitySet)
	e.position = position
	//TODO 初始化位置如果越界 需要修正
	e.direction = direction
}

func (e *Entity) OnEnterAOI(otherAoi *aoi.AOI) {
	otherEntity := otherAoi.Data.(*Entity)
	e.neighbors.Add(otherEntity)
	e.proxy.OnEntityEnter(otherEntity)
	//e.client.sendCreateEntity(other, false)
}

func (e *Entity) OnLeaveAOI(otherAoi *aoi.AOI) {
	otherEntity := otherAoi.Data.(*Entity)
	e.neighbors.Del(otherEntity)
	e.proxy.OnEntityLeave(otherEntity)
}

// IsNeighbor checks if other entity is a neighbor
func (e *Entity) IsNeighbor(other *Entity) bool {
	return e.neighbors.Contains(other)
}

// DistanceTo calculates the distance between two entities
//func (e *Entity) DistanceTo(other *Entity) float32 {
//	return entity.DistanceTo(e.position, other.position)
//}

func (e *Entity) String() string {
	return fmt.Sprintf("%s<%s>", e, e.GetID())
}
