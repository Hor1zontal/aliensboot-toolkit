/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/3/30
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package service

import (
	"aliens/protocol/hall"
	"aliens/module/cluster/rpc"
	"aliens/protocol/room"
)


func handleQuickMatch(request *hall.QuickMatch, response *hall.QuickMatchRet) {
	ret := rpc.Proxy_room.AllocFreeRoom(&room.AllocFreeRoom{GameID:1})
	response.Result = hall.HallResult_success
	response.MatchResult = &hall.QuickMatchResult{
		RoomID:ret.GetRoomID(),
		Address:ret.GetAddress(),
		Token:ret.GetToken(),
	}
}
