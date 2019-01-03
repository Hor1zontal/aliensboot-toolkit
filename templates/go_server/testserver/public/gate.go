package main

import (
	"github.com/KylinHe/aliensboot"
	"e.coding.net/aliens/aliensboot_testserver/module/gate"
)

func init() {

}

func main() {

	aliens.Run(
		gate.Module,
	)

}
