package main

import (
	"github.com/KylinHe/aliensboot"
	"github.com/KylinHe/aliensboot/module/database"
	"e.coding.net/aliens/aliensboot_testserver/module/game"
	"e.coding.net/aliens/aliensboot_testserver/module/gate"
	"e.coding.net/aliens/aliensboot_testserver/module/passport"
	"e.coding.net/aliens/aliensboot_testserver/module/scene"
)

func init() {

}

func main() {

	aliens.Run(
		database.Module,
		gate.Module,
		passport.Module,
		game.Module,
		scene.Module,
	)

}
