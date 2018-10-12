/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/8/31
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core

import "encoding/json"

type Attr map[string]interface{}

func (attr *Attr) Marshal() ([]byte, error) {
	return json.Marshal(attr)
}

func (attr *Attr) Unmarshal(data []byte) error {
	return json.Unmarshal(data, attr)
}