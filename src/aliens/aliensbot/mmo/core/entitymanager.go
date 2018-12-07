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
	"aliens/aliensbot/mmo/unit"
	"errors"
	"fmt"
	"reflect"
)

var EntityManager = newEntityManager()

type _EntityManager struct {

	entities       EntityMap //所有实体 id-entity

	entitiesByType map[EntityType]EntityMap //实体按类型分类 type[id-entity]

	entitiesDesc   map[EntityType]*EntityDesc //实体元数据 type-entity_meta
}

func newEntityManager() *_EntityManager {
	return &_EntityManager{
		entities:       EntityMap{},
		entitiesByType: map[EntityType]EntityMap{},
		entitiesDesc:   map[EntityType]*EntityDesc{},
	}
}

func (em *_EntityManager) put(entity *Entity) {
	em.entities.Add(entity)
	etype := entity.GetType()
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
	if entities, ok := em.entitiesByType[e.GetType()]; ok {
		entities.Del(eid)
	}
}

func (em *_EntityManager) traverseByType(eType EntityType, cb func(e *Entity)) {
	entities := em.entitiesByType[eType]
	for _, e := range entities {
		cb(e)
	}
}

// GenEntityID generates a new EntityID
func (em *_EntityManager) genEntityID() EntityID {
	return EntityID(util.GenUUID())
}

func (em *_EntityManager) GetEntity(id EntityID) *Entity {
	return em.entities.Get(id)
}

//处理远程调用
func (em *_EntityManager) HandleRemoteEntityCall(caller EntityID, id EntityID, method string, args[][]byte) (*Entity, error) {
	entity := em.GetEntity(id)
	if entity == nil {
		return nil, nil
	}
	return entity, entity.onCallFromRemote(caller, method, args)
}


//处理本地调用
func (em *_EntityManager) HandleLocalEntityCall(id EntityID, method string, args []interface{}) (*Entity, error) {
	entity := em.GetEntity(id)
	if entity == nil {
		return nil, nil
	}
	return entity, entity.OnCallFromLocal(method, args)
}

//func (em *_EntityManager) UNRegisterEntity(entity IEntity) {
//
//	delete(em.entitiesDesc, typeName)
//}

// RegisterEntity registers custom entity type and define entity behaviors
func (em *_EntityManager) RegisterEntity(entity IEntity) *EntityDesc {
	entityVal := reflect.ValueOf(entity)
	entityType := entityVal.Type()
	if entityType.Kind() == reflect.Ptr {
		entityType = entityType.Elem()
	}
	typeName := EntityType(entityType.Name())

	if desc, ok := em.entitiesDesc[typeName]; ok {
		log.Warnf("RegisterEntity: Entity type %s already registered", typeName)
		return desc
	}

	methodDesc := methodDescMap{}
	entityTypeDesc := &EntityDesc{
		name:         typeName,
		useAOI:       false,
		entityType:   entityType,
		methodDesc:   methodDesc,
		clientAttrs:  set.StringSet{},
		allAttrs:     set.StringSet{},
		persistAttrs: set.StringSet{},
	}
	em.entitiesDesc[typeName] = entityTypeDesc

	entityPtrType := reflect.PtrTo(entityType)
	numMethods := entityPtrType.NumMethod()

	for i := 0; i < numMethods; i++ {
		method := entityPtrType.Method(i)
		methodDesc.visit(method)
	}

	//// define entity Attrs
	entity.DescribeEntityType(entityTypeDesc)
	log.Infof(">>> RegisterEntity %s => %s <<<", typeName, entityType.Name())
	return entityTypeDesc
}


//从元数据中初始化一个实体
func (em *_EntityManager) CreateEntity(entityType EntityType, space *Space, pos unit.Vector, entityID EntityID) (*Entity, error) {
	entityDesc, ok := em.entitiesDesc[entityType]

	if !ok {
		return nil, errors.New(fmt.Sprintf("unknown entity type: %s", entityType))
	}

	//没有实体id自定义生成
	if entityID == "" {
		entityID = em.genEntityID()
	}


	entityInstance := reflect.New(entityDesc.entityType)
	entity := reflect.Indirect(entityInstance).FieldByName("Entity").Addr().Interface().(*Entity)
	//entity := &entity1
	entity.desc = entityDesc

	entity.init(entityID, entityInstance)
	em.put(entity)

	log.Debugf("Entity %s created.", entity)

	//entity.OnCreated()
	entity.I.OnCreated()

	if space != nil {
		space.enter(entity, pos)
	}

	return entity, nil
}
