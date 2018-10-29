package main

import (
	"aliens"
	"aliens/module/database"
	"aliens/module/game"
)

func init() {

}

func main() {
	aliens.Run(
		database.Module,
		game.Module,
	)
}
