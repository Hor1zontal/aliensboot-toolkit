package main

import (
	"github.com/KylinHe/aliensboot"
	"github.com/KylinHe/aliensboot/module/database"
	"e.coding.net/aliens/aliensboot_testserver/module/room"
)

func init() {

}

func main() {

	aliens.Run(
		database.Module,
		room.Module,
	)
}
