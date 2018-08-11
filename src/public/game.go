package main

import (
	"aliens/module/cluster"
	"time"
	"github.com/name5566/leaf"
	"math/rand"
	"aliens/module/game"
	"aliens/module/statistics"
	"aliens/module/database"
)

func init() {

}

func main() {
	//defer log.Close()
	//log.Init("conf/aliens/log.xml")
	rand.Seed(time.Now().UnixNano())
	leaf.Run(
		cluster.Module,
		statistics.Module,
		database.Module,
		game.Module,
	)
}
