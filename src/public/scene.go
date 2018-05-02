package main

import (
	"aliens/module/cluster"
	"time"
	"github.com/name5566/leaf"
	"math/rand"
	"aliens/module/scene"
)

func init() {

}

func main() {
	rand.Seed(time.Now().UnixNano())
	leaf.Run(
		cluster.Module,
		scene.Module,
	)
}
