package main

import (
	"github.com/KylinHe/aliensboot"
	"github.com/KylinHe/aliensboot/module/database"
	"e.coding.net/aliens/aliensboot_testserver/module/defaultmodule"
)

func init() {

}

func main() {

	aliens.Run(
		database.Module,
		defaultmodule.Module,
	)
}
