package main

import (
	"aliens/module/room"
	"aliens/module/cluster"
	"aliens/module/gate"
	"aliens"
)

func init() {

}

func main() {

	aliens.Run(
		cluster.Module,
		gate.Module,
		room.Module,
	)
}
