/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2017/4/19
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package util

import (
	"encoding/json"
	"math/rand"
)

func CopyMap(source map[int32]float32) map[int32]float32 {
	target := make(map[int32]float32)
	for key, value := range source {
		target[key] = value
	}
	return target
}

func CopyFloat64Map(source map[int32]float64) map[int32]float64 {
	target := make(map[int32]float64)
	for key, value := range source {
		target[key] = value
	}
	return target
}

func CopyInt32Map(source map[int32]int32) map[int32]int32 {
	target := make(map[int32]int32)
	for key, value := range source {
		target[key] = value
	}
	return target
}

func CopyInt64Map(source map[int32]int64) map[int32]int64 {
	target := make(map[int32]int64)
	for key, value := range source {
		target[key] = value
	}
	return target
}


func MargeMap(a map[int32]int32, b map[int32]int32) {
	for key, value := range b {
		a[key] = value
	}
}

func CopyMapInt32(a map[int32]int32) map[int32]int32 {
	results := make(map[int32]int32)
	for key, value := range a {
		results[key] = value
	}
	return results
}

func JSONCopy(marshaler interface{}, unMarshaler interface{}) error {
	data, error := json.Marshal(marshaler)
	if error != nil {
		return error
	}
	return json.Unmarshal(data, unMarshaler)
}

func RandomMultiWeight(weightMapping map[int32]int32, count int) []int32 {
	results := []int32{}
	for i := 0; i < count; i++ {
		result := RandomWeight(weightMapping)
		if result == 0 {
			return results
		}
		results = append(results, result)
		delete(weightMapping, result)
	}
	return results
}

func RandomWeight(weightMapping map[int32]int32) int32 {
	var totalWeight int32 = 0
	for _, weight := range weightMapping {
		totalWeight += weight
	}
	if totalWeight <= 0 {
		return 0
	}
	randomValue := rand.Int31n(totalWeight) + 1
	var currentValue int32 = 0
	for id, weight := range weightMapping {
		currentValue += weight
		if currentValue >= randomValue {
			return id
		}
	}
	return 0
}
