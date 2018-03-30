package main

import (
	"aliens/module/cluster"
	"time"
	"github.com/name5566/leaf"
	"math/rand"
	"aliens/log"
	"aliens/module/scene"
)

func init() {

}

func main() {
	defer log.Close()
	log.Init("conf/aliens/log.xml")
	rand.Seed(time.Now().UnixNano())
	leaf.Run(
		cluster.Module,
		//gate.Module,
		//service1.Module,
		//service2.Module,
		scene.Module,
	)
}
