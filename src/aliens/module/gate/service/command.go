/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/8/22
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package service

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"aliens/gate"
)

type userPush struct {
	pushMsg interface{}
}

type NetworkInit struct {
	pid *actor.PID
	agent gate.Agent
}