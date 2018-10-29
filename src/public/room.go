package main

import (
	"aliens"
	"aliens/module/cluster"
	"aliens/module/gate"
	"aliens/module/room"
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
