/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/3/30
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package service

import (
	"aliens/aliensbot/log"
	"aliens/testserver/module/gate/network"
	"aliens/testserver/protocol"
)

//
func handleBindService(request *protocol.BindService) {
	log.Debugf("bind : %v - %v", request.GetAuthID(), request.GetBinds())
	network.Manager.BindService(request.GetAuthID(), request.GetBinds())
}
