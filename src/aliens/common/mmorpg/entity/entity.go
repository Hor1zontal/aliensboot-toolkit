package entity

import (
	"aliens/mmorpg/aoi"
	"fmt"
)

var currID int32 = 0

func GenEntityID () int32 {
	currID++
	return currID
}

type IEntity interface {
	GetDist() float32
	OnEntityEnter(entity *Entity) //call it when other entity EntityEnter
	OnEntityLeave(entity *Entity) //call it when other entity EntityLeave
	OnEntityMove(entity *Entity)  //call it when other entity EntityMove
}

type Entity struct {
	id int32
	//client Client
	space *Space   //

	aoi   *aoi.AOI //one entity , one aoi!
	position  Vector3 //entity position
	direction Vector3   //entity direction

	neighbors EntitySet //aoi entity

	proxy IEntity
}

func (e *Entity) GetID() int32 {
	return e.id
}

func (e *Entity) GetPosition() Vector3 {
	return e.position
}

func (e *Entity) GetDirection() Vector3 {
	return e.direction
}

func (e *Entity) Init(space *Space, proxy IEntity, position Vector3, direction Vector3) {
	e.id = GenEntityID()
	e.aoi = aoi.NewAOI(e, proxy.GetDist())
	e.proxy = proxy
	e.neighbors = make(EntitySet)
	e.position = position
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
func (e *Entity) DistanceTo(other *Entity) float32 {
	return e.position.DistanceTo(other.position)
}

func (e *Entity) String() string {
	return fmt.Sprintf("%s<%s>", e, e.GetID())
}
