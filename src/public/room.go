package main

import (
	"github.com/name5566/leaf"
	"aliens/module/room"
	"aliens/module/cluster"
)

func init() {

}

func main() {
	//defer log.Close()
	//log.Init("conf/aliens/log.xml")
	//rand.Seed(time.Now().UnixNano())
	leaf.Run(
		cluster.Module,
		room.Module,
	)
}
