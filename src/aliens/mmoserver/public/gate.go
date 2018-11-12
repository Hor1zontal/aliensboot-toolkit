package main

import (
	"aliens/aliensbot"
	"aliens/mmoserver/module/gate"
)

func init() {

}

func main() {

	aliens.Run(
		gate.Module,
	)

}
