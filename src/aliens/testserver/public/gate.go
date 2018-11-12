package main

import (
	"aliens/aliensbot"
	"aliens/testserver/module/gate"
)

func init() {

}

func main() {

	aliens.Run(
		gate.Module,
	)

}
