package main

import (
	"time"
	"github.com/name5566/leaf"
	"aliens/log"
	"math/rand"
	"aliens/module/gate"
	"aliens/module/cluster"
	"aliens/module/service1"
	"aliens/module/service2"
	"aliens/module/scene"
)

func main() {
	defer log.Close()
	log.Init("conf/aliens/log.xml")
	rand.Seed(time.Now().UnixNano())

	leaf.Run(
		cluster.Module,
		gate.Module,
		service1.Module,
		service2.Module,
		scene.Module,
	)
}
