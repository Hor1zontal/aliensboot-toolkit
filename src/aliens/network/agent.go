/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/6/11
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package network

import "net"


type UDPAgent struct {

	conn *net.UDPConn

	udpAddr *net.UDPAddr

	userData interface{}
}


//发送数据
func (this *UDPAgent) WriteData(data []byte) {
	this.conn.WriteToUDP(data, this.udpAddr)
}

func (this *UDPAgent) SetUserData(userData interface{}) {
	this.userData = userData
}

func (this *UDPAgent) UserData() interface{} {
	return this.userData
}

func (this *UDPAgent) GetID() string {
	if this.udpAddr == nil {
		return ""
	}
	return this.udpAddr.String()
}


