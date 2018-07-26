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
	"aliens/protocol/game"
	"github.com/pkg/errors"
    "encoding/json"
)

//处理json请求
func handleJsonRequest(authID int64, requestUrl string, data []byte) ([]byte, error) {
	switch requestUrl {
	    
		case "LoginRole" :
			request := &game.LoginRole{}
			response := &game.LoginRoleRet{}
			error := json.Unmarshal(data, request)
			if error != nil {
				return nil, error
			}
			handleLoginRole(authID, request, response)
			return json.Marshal(response)
		
		case "CreateRole" :
			request := &game.CreateRole{}
			response := &game.CreateRoleRet{}
			error := json.Unmarshal(data, request)
			if error != nil {
				return nil, error
			}
			handleCreateRole(authID, request, response)
			return json.Marshal(response)
		
		case "RemoveRole" :
			request := &game.RemoveRole{}
			response := &game.RemoveRoleRet{}
			error := json.Unmarshal(data, request)
			if error != nil {
				return nil, error
			}
			handleRemoveRole(authID, request, response)
			return json.Marshal(response)
		
		case "GetUserInfo" :
			request := &game.GetUserInfo{}
			response := &game.GetUserInfoRet{}
			error := json.Unmarshal(data, request)
			if error != nil {
				return nil, error
			}
			handleGetUserInfo(authID, request, response)
			return json.Marshal(response)
		
		default:
		    return nil, errors.New("unexpect request " + requestUrl)

	}
}
