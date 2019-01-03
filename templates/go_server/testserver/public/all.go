package main

import (
	"e.coding.net/aliens/aliensboot_testserver/module/game"
	"e.coding.net/aliens/aliensboot_testserver/module/gate"
	"e.coding.net/aliens/aliensboot_testserver/module/hall"
	"e.coding.net/aliens/aliensboot_testserver/module/passport"
	"e.coding.net/aliens/aliensboot_testserver/module/room"
	"github.com/KylinHe/aliensboot"
	"github.com/KylinHe/aliensboot/module/database"
)

func init() {

}

func main() {

	aliens.Run(
		database.Module,
		gate.Module,
		hall.Module,
		room.Module,
		passport.Module,
		game.Module,
	)
}
