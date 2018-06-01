/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/5/25
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package game

import (
	"net"
	"fmt"
)

type Player struct {
	id uint32

	network *net.UDPAddr //UDP连接

	lostFrame []uint32

	//id agent.ID
	game *Game  //当前加入的游戏

}

//是否有丢帧请求
func (this *Player) haveLostFrame() bool {
	return this.lostFrame != nil
}

//发动数据
func (this *Player) sendData(data []byte) {

	//_, err = socket.WriteToUDP(data, this.network)
}

//发动同步帧
func (this *Player) sendFrame(frame *Frame) {

	this.lostFrame = nil
}

//
func (this *Player) acceptCommand(command *Command) {
	command.PlayerID = this.id

	//_, err = socket.WriteToUDP(data, this.network)
}

func (this *Player) acceptLostFrame(command *RequestLostFrame) {


	//_, err = socket.WriteToUDP(data, this.network)
}


func init() {
	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("监听失败!", err)
		return
	}
	defer socket.Close()

	for {
		// 读取数据
		data := make([]byte, 1024)
		read, remoteAddr, err := socket.ReadFromUDP(data)
		if err != nil {
			fmt.Println("读取数据失败!", err)
			continue
		}
		fmt.Println(read, remoteAddr)
		fmt.Printf("%s\n\n", data)

		// 发送数据
		senddata := []byte("hello client!")
		_, err = socket.WriteToUDP(senddata, remoteAddr)
		if err != nil {
			return
			fmt.Println("发送数据失败!", err)
		}
	}

}
