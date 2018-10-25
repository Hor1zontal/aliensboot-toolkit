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
	"aliens/protocol/hall"
	"github.com/pkg/errors"
    "encoding/json"
)

//处理json请求
func handleJsonRequest(requestUrl string, data []byte) ([]byte, error) {
	switch requestUrl {
	    
		case "QuickMatch" :
			request := &hall.QuickMatch{}
			response := &hall.QuickMatchRet{}
			error := json.Unmarshal(data, request)
			if error != nil {
				return nil, error
			}
			handleQuickMatch(request, response)
			return json.Marshal(response)
		
		default:
		    return nil, errors.New("unexpect request " + requestUrl)

	}
}
