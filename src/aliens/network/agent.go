/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/6/13
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package network

type Agent interface {
	WriteMsg(msg interface{})
	//LocalAddr() net.Addr
	//RemoteAddr() net.Addr
	//Close()
	//Destroy()
	UserData() interface{}
	SetUserData(data interface{})
}