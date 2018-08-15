package main

import (
	"aliens/module/hall"
	"aliens/module/cluster"
	"aliens"
)

func init() {

}

func main() {

	aliens.Run(
		cluster.Module,
		hall.Module,
	)
}
