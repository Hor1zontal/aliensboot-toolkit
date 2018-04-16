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
	"aliens/protocol/passport"
	"github.com/pkg/errors"
    "encoding/json"
)

//处理json请求
func handleJsonRequest(requestUrl string, data []byte) ([]byte, error) {
	switch requestUrl {
	    
		case "LoginRegister" :
			request := &passport.LoginRegister{}
			response := &passport.LoginRegisterRet{}
			error := json.Unmarshal(data, request)
			if error != nil {
				return nil, error
			}
			handleLoginRegister(request, response)
			return json.Marshal(response)
		
		case "LoginLogin" :
			request := &passport.LoginLogin{}
			response := &passport.LoginLoginRet{}
			error := json.Unmarshal(data, request)
			if error != nil {
				return nil, error
			}
			handleLoginLogin(request, response)
			return json.Marshal(response)
		
		default:
		    return nil, errors.New("unexpect request " + requestUrl)

	}
}
