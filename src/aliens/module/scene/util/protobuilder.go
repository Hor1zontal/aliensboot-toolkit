/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/3/26
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package util

import (
	"aliens/protocol"
	"aliens/mmorpg/core"
)

func BuildEntities(entitySet core.EntitySet) []*protocol.Entity {
	results := make([]*protocol.Entity, entitySet.Len())
	for entity, _ := range entitySet {
		results = append(results, BuildEntityProtocol(entity))
	}
	return results
}

func BuildEntityProtocol(entity *core.Entity) *protocol.Entity {
	return &protocol.Entity{
		//Id:       int64(entity.GetID()),
		Position: entity.GetPosition(),
		Direction: entity.GetDirection(),
	}
}

//func BuildVector(vector entity.Vector3) *protocol.Vector {
//	return &protocol.Vector{
//		X: vector.X,
//		Y: vector.Y,
//		Z: vector.Z,
//	}
//}
//
//func TransVector(vector *protocol.Vector) entity.Vector3 {
//	return entity.Vector3{
//		X: vector.GetX(),
//		Y: vector.GetY(),
//		Z: vector.GetZ(),
//	}
//}
