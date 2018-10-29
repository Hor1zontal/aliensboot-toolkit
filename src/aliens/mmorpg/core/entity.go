package core

import (
	"aliens/log"
	"aliens/mmorpg/aoi"
	"aliens/protocol"
	"fmt"
	"reflect"
	"unsafe"
)

//type IEntity interface {
//	OnInit(self *Entity)
//	OnEntityEnter(entity *Entity) //call it when other entity EntityEnter
//	OnEntityLeave(entity *Entity) //call it when other entity EntityLeave
//	OnEntityMove(entity *Entity)  //call it when other entity EntityMove
//	OnReceive(entity *Entity, event interface{}) //call it when listener event happen
//}

type Entity struct {
	//client Client
	//space *Space   //

	*protocol.Entity //sync data

	clientID *ClientID

	I IEntity

	V reflect.Value

	space *Space //实体所在的空间id

	aoi *aoi.AOI //one entity , one aoi!

	attrs Attr //属性列表

	interestedIn EntitySet //当前实体视野范围内的实体

	interestedBy EntitySet //视野内存在当前实体的对象

	destroyed bool
}

func (e *Entity) GetID() EntityID {
	return EntityID(e.Id)
}

// SetPosition sets the entity position
func (e *Entity) SetPosition(pos *protocol.Vector) {
	e.setPositionYaw(pos, e.Yaw)
}

func (e *Entity) setPositionYaw(pos *protocol.Vector, yaw float32) {
	space := e.space
	if space == nil {
		log.Warnf("%s.SetPosition(%s): space is nil", e, pos)
		return
	}
	space.move(e, pos)
	//e.yaw = yaw
}

func (e *Entity) init(typeName string, entityID EntityID, entityInstance reflect.Value, attrs Attr) {
	e.Id = string(entityID)
	e.V = entityInstance
	e.I = entityInstance.Interface().(IEntity)
	e.TypeName = typeName
	//e.typeDesc = registeredEntityTypes[typeName]
	if attrs != nil {
		e.attrs = attrs
	} else {
		e.attrs = make(Attr)
	}
	e.interestedIn = make(EntitySet)
	e.interestedBy = make(EntitySet)
	e.aoi = aoi.NewAOI(e, 500)
	//e.I.OnAttrsReady()
	//e.I.OnCreated()
	e.I.OnInit()
}

func (e *Entity) OnEnterAOI(otherAoi *aoi.AOI) {
	e.interest(otherAoi.Callback.(*Entity))
}

func (e *Entity) OnLeaveAOI(otherAoi *aoi.AOI) {
	e.uninterest(otherAoi.Callback.(*Entity))
}

// Destroy destroys the entity
func (e *Entity) Destroy() {
	if e.destroyed {
		return
	}
	log.Debugf("%s.Destroy ...", e)
	e.destroyEntity(false)
	//TODO
	//dispatchercluster.SendNotifyDestroyEntity(e.id)
}

func (e *Entity) destroyEntity(isMigrate bool) {
	e.space.leave(e)
	if !isMigrate {
		e.I.OnDestroy()
	} else {
		e.I.OnMigrateOut()
	}
	if !isMigrate {
		//e.SetClient(nil) // always set Client to nil before destroy
		//e.Save()
	} else {
		//e.assignClient(nil)
	}
	e.destroyed = true
	EntityManager.del(e)
}

// IsInterestedIn checks if other entity is interested by this entity
func (e *Entity) IsInterestedIn(other *Entity) bool {
	return e.interestedIn.Contains(other)
}

// DistanceTo calculates the distance between two entities
//func (e *Entity) DistanceTo(other *Entity) Coord {
//	return e.Position.DistanceTo(other.Position)
//}

func (e *Entity) String() string {
	return fmt.Sprintf("%s<%s>", e, e.GetID())
}

//// IsSpaceEntity returns if the entity is actually a space
//func (e *Entity) IsSpaceEntity() bool {
//	return e.typeName == _SPACE_ENTITY_TYPE
//}

// AsSpace converts entity to space (only works for space entity)
func (e *Entity) AsSpace() *Space {
	//if !e.IsSpaceEntity() {
	//	gwlog.Panicf("%s is not a space", e)
	//}

	return (*Space)(unsafe.Pointer(e))
}

// Interests and Uninterest among entities
func (e *Entity) interest(other *Entity) {
	e.interestedIn.Add(other)
	other.interestedBy.Add(e)
	//e.proxy.OnEntityEnter(other)
	//e.client.sendCreateEntity(other, false)
}

func (e *Entity) uninterest(other *Entity) {
	e.interestedIn.Del(other)
	other.interestedBy.Del(e)
	//e.proxy.OnEntityLeave(other)
	//e.client.sendDestroyEntity(other)
}
