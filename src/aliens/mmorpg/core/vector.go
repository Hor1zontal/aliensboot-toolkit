/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/3/21
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core

import (
	"aliens/protocol"
	"math"
)

// Yaw is the type of entity direction
type Yaw float32

// protocol.Vector is type of entity position
//type protocol.Vector struct {
//	X float32
//	Y float32
//	Z float32
//}
//
//func (p Vector) String() string {
//	return fmt.Sprintf("(%.2f, %.2f, %.2f)", p.X, p.Y, p.Z)
//}

// DistanceTo calculates distance between two positions
func DistanceTo(p protocol.Vector, o protocol.Vector) float32 {
	dx := p.X - o.X
	dy := p.Y - o.Y
	dz := p.Z - o.Z
	return float32(math.Sqrt(float64(dx*dx + dy*dy + dz*dz)))
}

// Sub calculates protocol.Vector p - protocol.Vector o
func Sub(p protocol.Vector, o protocol.Vector) protocol.Vector {
	return protocol.Vector{p.X - o.X, p.Y - o.Y, p.Z - o.Z}
}

func Add(p protocol.Vector, o protocol.Vector) protocol.Vector {
	return protocol.Vector{p.X + o.X, p.Y + o.Y, p.Z + o.Z}
}

// Mul calculates protocol.Vector p * m
func Mul(p protocol.Vector, m float32) protocol.Vector {
	return protocol.Vector{p.X * m, p.Y * m, p.Z * m}
}

// DirToYaw convert direction represented by protocol.Vector to Yaw
func DirToYaw(dir protocol.Vector) Yaw {
	Normalized(&dir)
	yaw := math.Acos(float64(dir.X))
	if dir.Z < 0 {
		yaw = math.Pi*2 - yaw
	}
	yaw = yaw / math.Pi * 180 // convert to angle
	if yaw <= 90 {
		yaw = 90 - yaw
	} else {
		yaw = 90 + (360 - yaw)
	}

	return Yaw(yaw)
}

func Normalize(p *protocol.Vector) {
	d := float32(math.Sqrt(float64(p.X*p.X + p.Y + p.Y + p.Z*p.Z)))
	if d == 0 {
		return
	}
	p.X /= d
	p.Y /= d
	p.Z /= d
}

func Normalized(p *protocol.Vector) {
	Normalize(p)
}
