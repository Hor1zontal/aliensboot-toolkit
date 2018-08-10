/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/3/26
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package util

import (
	"aliens/mmorpg/entity"
	"aliens/protocol"
)


func BuildEntities(entitySet entity.EntitySet) []*protocol.Entity {
	results := []*protocol.Entity{}
	for entity, _ := range entitySet {
		results = append(results, BuildEntityProtocol(entity))
	}
	return results
}

func BuildEntityProtocol(entity *entity.Entity) *protocol.Entity {
	return &protocol.Entity{
		Id:        entity.GetID(),
		Position:  BuildVector(entity.GetPosition()),
		Direction: BuildVector(entity.GetDirection()),
	}
}

func BuildVector(vector entity.Vector3) *protocol.Vector {
	return &protocol.Vector{
		X: vector.X,
		Y: vector.Y,
		Z: vector.Z,
	}
}

func TransVector(vector *protocol.Vector) entity.Vector3 {
	return entity.Vector3{
		X: vector.GetX(),
		Y: vector.GetY(),
		Z: vector.GetZ(),
	}
}
