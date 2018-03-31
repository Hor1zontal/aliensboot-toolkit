/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/3/31
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package template



type ProtoMessage struct {
	PackageName string
	RequestName string
	ResponseName string

	Handlers map[int]*ProtoHandler
}


type ProtoHandler struct {
	//Name string
	Desc string
	ORequest string
	OResponse string
}
