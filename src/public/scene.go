package main

import (
	"aliens/module/cluster"
	"aliens/module/scene"
	"aliens"
)

func init() {

}

func main() {

	aliens.Run(
		cluster.Module,
		scene.Module,
	)
}
