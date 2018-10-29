package main

import (
	"aliens"
	"aliens/module/cluster"
	"aliens/module/scene"
)

func init() {

}

func main() {

	aliens.Run(
		cluster.Module,
		scene.Module,
	)
}
