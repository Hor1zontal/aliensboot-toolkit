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
	"aliens/protocol/scene"
)


func BuildEntities(entitySet entity.EntitySet) []*scene.Entity {
	results := []*scene.Entity{}
	for entity, _ := range entitySet {
		results = append(results, BuildEntityProtocol(entity))
	}
	return results
}

func BuildEntityProtocol(entity *entity.Entity) *scene.Entity {
	return &scene.Entity{
		Id:        entity.GetID(),
		Position:  BuildVector(entity.GetPosition()),
		Direction: BuildVector(entity.GetDirection()),
	}
}

func BuildVector(vector entity.Vector3) *scene.Vector {
	return &scene.Vector{
		X: vector.X,
		Y: vector.Y,
		Z: vector.Z,
	}
}

func TransVector(vector *scene.Vector) entity.Vector3 {
	return entity.Vector3{
		X: vector.GetX(),
		Y: vector.GetY(),
		Z: vector.GetZ(),
	}
}
