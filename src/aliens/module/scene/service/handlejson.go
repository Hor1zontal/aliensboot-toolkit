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
	"aliens/protocol/scene"
	"github.com/pkg/errors"
    "encoding/json"
)

//处理json请求
func handleJsonRequest(requestUrl string, data []byte) ([]byte, error) {
	switch requestUrl {
	    
		case "SpaceLeave" :
			request := &scene.SpaceLeave{}
			response := &scene.SpaceLeaveRet{}
			error := json.Unmarshal(data, request)
			if error != nil {
				return nil, error
			}
			handleSpaceLeave(request, response)
			return json.Marshal(response)
		
		case "GetState" :
			request := &scene.GetState{}
			response := &scene.GetStateRet{}
			error := json.Unmarshal(data, request)
			if error != nil {
				return nil, error
			}
			handleGetState(request, response)
			return json.Marshal(response)
		
		case "SpaceMove" :
			request := &scene.SpaceMove{}
			response := &scene.SpaceMoveRet{}
			error := json.Unmarshal(data, request)
			if error != nil {
				return nil, error
			}
			handleSpaceMove(request, response)
			return json.Marshal(response)
		
		case "SpaceEnter" :
			request := &scene.SpaceEnter{}
			response := &scene.SpaceEnterRet{}
			error := json.Unmarshal(data, request)
			if error != nil {
				return nil, error
			}
			handleSpaceEnter(request, response)
			return json.Marshal(response)
		
		default:
		    return nil, errors.New("unexpect request " + requestUrl)

	}
}
