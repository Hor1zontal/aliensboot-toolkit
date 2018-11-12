/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/8/31
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core

import (
	"aliens/aliensbot/common/data_structures/set"
	"aliens/aliensbot/common/util"
	"aliens/aliensbot/log"
	"aliens/aliensbot/protocol"
	"github.com/xiaonanln/goworld/engine/gwlog"
	"reflect"
	"strings"
)

var EntityManager = newEntityManager()

type _EntityManager struct {
	entities       EntityMap
	entitiesByType map[string]EntityMap
	entitiesDesc   map[string]*EntityDesc
}

func newEntityManager() *_EntityManager {
	return &_EntityManager{
		entities:       EntityMap{},
		entitiesByType: map[string]EntityMap{},
		entitiesDesc:   map[string]*EntityDesc{},
	}
}

func (em *_EntityManager) put(entity *Entity) {
	em.entities.Add(entity)
	etype := entity.TypeName
	eid := entity.GetID()
	if entities, ok := em.entitiesByType[etype]; ok {
		entities.Add(entity)
	} else {
		em.entitiesByType[etype] = EntityMap{eid: entity}
	}
}

func (em *_EntityManager) del(e *Entity) {
	eid := e.GetID()
	em.entities.Del(eid)
	if entities, ok := em.entitiesByType[e.TypeName]; ok {
		entities.Del(eid)
	}
}

func (em *_EntityManager) traverseByType(etype string, cb func(e *Entity)) {
	entities := em.entitiesByType[etype]
	for _, e := range entities {
		cb(e)
	}
}

// GenEntityID generates a new EntityID
func (em *_EntityManager) genEntityID() EntityID {
	return EntityID(util.GenUUID())
}

func (em *_EntityManager) getEntityTypeName(entity IEntity) string {
	entityType := reflect.TypeOf(entity)
	return entityType.Name()
}

func (em *_EntityManager) GetEntity(id EntityID) *Entity {
	return em.entities.Get(id)
}

func (em *_EntityManager) UNRegisterEntity(entity IEntity) {
	typeName := em.getEntityTypeName(entity)
	delete(em.entitiesDesc, typeName)
}

// RegisterEntity registers custom entity type and define entity behaviors
func (em *_EntityManager) RegisterEntity(entity IEntity, meta interface{}) *EntityDesc {
	entityVal := reflect.ValueOf(entity)
	entityType := entityVal.Type()
	typeName := em.getEntityTypeName(entity)

	if desc, ok := em.entitiesDesc[typeName]; ok {
		log.Warnf("RegisterEntity: Entity type %s already registered", typeName)
		return desc
	}

	if entityType.Kind() == reflect.Ptr {
		entityType = entityType.Elem()
	}

	methodDescs := methodDescMap{}
	// register the string of e
	entityTypeDesc := &EntityDesc{
		name:         typeName,
		useAOI:       false,
		entityType:   entityType,
		selfAttrs:    set.StringSet{},
		allAttrs:     set.StringSet{},
		persistAttrs: set.StringSet{},
	}
	em.entitiesDesc[typeName] = entityTypeDesc

	entityPtrType := reflect.PtrTo(entityType)
	numMethods := entityPtrType.NumMethod()

	for i := 0; i < numMethods; i++ {
		method := entityPtrType.Method(i)
		methodDescs.visit(method)
	}

	log.Infof(">>> RegisterEntity %s => %s <<<", typeName, entityType.Name())
	//// define entity Attrs
	entity.DescribeEntityType(entityTypeDesc)

	em.registerAttr(entityTypeDesc, meta)
	return entityTypeDesc
}

func (em *_EntityManager) registerAttr(desc *EntityDesc, meta interface{}) {
	util.VisitTag(meta, AttrTagFeature, func(fieldName string, tagValue string) {
		if strings.Contains(tagValue, AttrTagFeatureSelf) {
			desc.selfAttrs.Add(fieldName)
		} else if strings.Contains(tagValue, AttrTagFeatureAll) {
			desc.allAttrs.Add(fieldName)
		}

		if strings.Contains(tagValue, AttrTagFeaturePersist) {
			desc.persistAttrs.Add(fieldName)
		}
	})
}

//创建一个默认实体
func (em *_EntityManager) CreateLocalEntity(typeName string, space *Space, pos *protocol.Vector) *Entity {
	return em.createEntity(typeName, space, pos, "", nil)
}

//func (em *_EntityManager) NewPersistEntity(typeName string, space *Space, pos *protocol.Vector, entityID EntityID, attr Attr) *Entity {
//
//}

func (em *_EntityManager) createEntity(typeName string, space *Space, pos *protocol.Vector, entityID EntityID, attr Attr) *Entity {
	entityTypeDesc, ok := em.entitiesDesc[typeName]

	if !ok {
		gwlog.Panicf("unknown entity type: %s", typeName)
	}

	if entityID == "" {
		entityID = em.genEntityID()
	}

	var entity *Entity
	var entityInstance reflect.Value

	entityInstance = reflect.New(entityTypeDesc.entityType)
	entity = reflect.Indirect(entityInstance).FieldByName("Entity").Addr().Interface().(*Entity)
	entity.init(typeName, entityID, entityInstance, attr)
	em.put(entity)

	log.Debugf("Entity %s created.", entity)

	if space != nil {
		space.enter(entity, pos)
	}

	return entity
}

//func (em *_EntityManager) onGateDisconnected(gateid uint16) {
//	for _, entity := range em.entities {
//		client := entity.client
//		if client != nil && client.gateid == gateid {
//			entity.notifyClientDisconnected()
//		}
//	}
//}
