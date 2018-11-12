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
	"aliens/mmoserver/module/gate/network"
	"aliens/mmoserver/protocol"
)

//
func handleBindService(request *protocol.BindService) {
	network.Manager.BindService(request.GetAuthID(), request.GetBinds())
}
