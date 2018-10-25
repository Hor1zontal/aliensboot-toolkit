package main

import (
	"aliens/module/cluster"
	"aliens/module/game"
	"aliens/module/statistics"
	"aliens/module/database"
	"aliens"
)

func init() {

}

func main() {
	aliens.Run(
		database.Module,
		game.Module,
	)
}
