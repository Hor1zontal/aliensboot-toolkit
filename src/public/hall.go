package main

import (
	"aliens"
	"aliens/module/cluster"
	"aliens/module/hall"
)

func init() {

}

func main() {

	aliens.Run(
		cluster.Module,
		hall.Module,
	)
}
