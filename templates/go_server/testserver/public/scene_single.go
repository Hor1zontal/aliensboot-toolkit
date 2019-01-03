package main

import (
	"github.com/KylinHe/aliensboot"
	"github.com/KylinHe/aliensboot/module/database"
	"e.coding.net/aliens/aliensboot_testserver/module/scene"
)

func init() {

}

func main() {

	aliens.Run(
		database.Module,
		scene.Module,
	)

}
