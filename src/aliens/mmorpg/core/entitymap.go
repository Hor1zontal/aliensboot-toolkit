/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/3/21
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core

import (
	"bytes"
)

// EntityMap is the data structure for maintaining entity IDs to entities
type EntityMap map[EntityID]*Entity

// Add adds a new entity to EntityMap
func (em EntityMap) Add(entity *Entity) {
	em[entity.GetID()] = entity
}

// Del deletes an entity from EntityMap
func (em EntityMap) Del(id EntityID) {
	delete(em, id)
}

// Get returns the Entity of specified entity id in EntityMap
func (em EntityMap) Get(id EntityID) *Entity {
	return em[id]
}

// EntitySet is the data structure for a collection of entities
type EntitySet map[*Entity]struct{}

// Add adds an entity to the EntitySet
func (es EntitySet) Add(entity *Entity) {
	es[entity] = struct{}{}
}

// Del deletes an entity from the EntitySet
func (es EntitySet) Del(entity *Entity) {
	delete(es, entity)
}

func (es EntitySet) Len() int {
	return len(es)
}

// Contains returns if the entity is in the EntitySet
func (es EntitySet) Contains(entity *Entity) bool {
	_, ok := es[entity]
	return ok
}

func (es EntitySet) String() string {
	b := bytes.Buffer{}
	b.WriteString("{")
	first := true
	for entity := range es {
		if !first {
			b.WriteString(", ")
		} else {
			first = false
		}
		b.WriteString(entity.String())
	}
	b.WriteString("}")
	return b.String()
}

//// EntityIDSet is the data structure for a collection of entity IDs
//type EntityIDSet map[int32]struct{}
//
//// Add adds an entity id to EntityIDSet
//func (es EntityIDSet) Add(id int32) {
//	es[id] = struct{}{}
//}
//
//// Del removes an entity id from EntityIDSet
//func (es EntityIDSet) Del(id int32) {
//	delete(es, id)
//}
//
//// Contains checks if entity id is in EntityIDSet
//func (es EntityIDSet) Contains(id int32) bool {
//	_, ok := es[id]
//	return ok
//}
//
//// ToList convert EntityIDSet to a slice of entity IDs
//func (es EntityIDSet) ToList() []int32 {
//	list := make([]int32, 0, len(es))
//	for eid := range es {
//		list = append(list, eid)
//	}
//	return list
//}

