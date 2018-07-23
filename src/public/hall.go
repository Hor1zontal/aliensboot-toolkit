package main

import (
	"github.com/name5566/leaf"
	"aliens/module/hall"
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
		hall.Module,
	)
}
