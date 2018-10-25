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
	"aliens/protocol/room"
	"github.com/pkg/errors"
    "encoding/json"
)

//处理json请求
func handleJsonRequest(requestUrl string, data []byte) ([]byte, error) {
	switch requestUrl {
	    
		case "AllocFreeRoomSeat" :
			request := &room.AllocFreeRoomSeat{}
			response := &room.AllocFreeRoomSeatRet{}
			error := json.Unmarshal(data, request)
			if error != nil {
				return nil, error
			}
			handleAllocFreeRoomSeat(request, response, nil)
			return json.Marshal(response)
		
		case "JoinRoom" :
			request := &room.JoinRoom{}
			response := &room.JoinRoomRet{}
			error := json.Unmarshal(data, request)
			if error != nil {
				return nil, error
			}
			handleJoinRoom(request, response, nil)
			return json.Marshal(response)
		
		case "LeaveRoom" :
			request := &room.LeaveRoom{}
			response := &room.LeaveRoomRet{}
			error := json.Unmarshal(data, request)
			if error != nil {
				return nil, error
			}
			handleLeaveRoom(request, response, nil)
			return json.Marshal(response)
		
		default:
		    return nil, errors.New("unexpect request " + requestUrl)

	}
}
