package main

import (
	"aliens/testserver/module/gate"
	"aliens/aliensbot"
)

func init() {

}

func main() {

	aliens.Run(
		gate.Module,
	)

}
