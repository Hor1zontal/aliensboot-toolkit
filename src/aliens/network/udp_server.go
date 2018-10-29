/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/6/8
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package network

import (
	"aliens/log"
	"net"
	"strconv"
)

type UDPServer struct {
	conn *net.UDPConn

	protocol string
	address  string

	agents map[string]*UDPAgent

	handle func(data []byte, addr *UDPAgent)
}

func (this *UDPServer) Start(config Config, handle func(data []byte, addr *UDPAgent)) {
	this.init(config, handle)
	go this.run()
}

func (this *UDPServer) init(config Config, handle func(data []byte, addr *UDPAgent)) {
	this.handle = handle
	this.protocol = config.Protocol
	this.address = config.Address + ":" + strconv.Itoa(config.Port)

	this.agents = make(map[string]*UDPAgent)

	udpAddr, err := net.ResolveUDPAddr(this.protocol, this.address)
	if err != nil {
		log.Fatalf("star udp server error : %v", err)
	}
	udpConn, err2 := net.ListenUDP("udp", udpAddr)
	if err2 != nil {
		log.Fatalf("star udp server error : %v", err2)
	}
	this.conn = udpConn
}

func (this *UDPServer) run() {
	defer this.conn.Close()
	for {
		buf := make([]byte, 512)
		//读取数据
		len, udpAddr, err := this.conn.ReadFromUDP(buf)
		if err != nil {
			log.Errorf("read udp msg err : %v", err)
		}

		agent := this.agents[udpAddr.String()]
		if agent == nil {
			agent = &UDPAgent{conn: this.conn, udpAddr: udpAddr}
			this.agents[udpAddr.String()] = agent
		}
		this.handle(buf[0:len], agent)
	}
}
