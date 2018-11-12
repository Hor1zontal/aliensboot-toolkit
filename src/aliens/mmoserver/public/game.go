package main

import (
	"aliens/aliensbot"
	"aliens/aliensbot/module/database"
	"aliens/mmoserver/module/game"
)

func init() {

}

func main() {
	aliens.Run(
		database.Module,
		game.Module,
	)
}
