package main

import (
	"aliens/module/cluster"
	"aliens/module/gate"
	"aliens"
	"aliens/module/statistics"
)

func init() {

}

func main() {

	aliens.Run(
		cluster.Module,
		statistics.Module,
		gate.Module,
	)

}
