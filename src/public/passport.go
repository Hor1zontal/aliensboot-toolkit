package main

import (
	"aliens/module/cluster"
	"aliens/module/passport"
	"aliens"
)

func init() {

}

func main() {

	aliens.Run(
		cluster.Module,
		passport.Module,
	)
}
