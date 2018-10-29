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
	"aliens/common/util"
	"aliens/log"
)

type ClientID struct {
	gateID   string
	authID   int64
	entityID EntityID
}

// EntityID type
type EntityID string

// IsNil returns if EntityID is nil
func (id EntityID) IsNil() bool {
	return id == ""
}

// MustEntityID assures a string to be EntityID
func MustEntityID(id string) EntityID {
	if len(id) != util.UUID_LENGTH {
		log.Panicf("%s of len %d is not a valid entityD (len=%d)", id, len(id), util.UUID_LENGTH)
	}
	return EntityID(id)
}
