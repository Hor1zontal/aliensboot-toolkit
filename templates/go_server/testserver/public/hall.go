package main

import (
	"github.com/KylinHe/aliensboot"
	"github.com/KylinHe/aliensboot/module/database"
	"e.coding.net/aliens/aliensboot_testserver/module/hall"
)

func init() {

}

func main() {

	aliens.Run(
		database.Module,
		hall.Module,
	)
}
