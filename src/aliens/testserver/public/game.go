package main

import (
	"aliens/aliensbot"
	"aliens/aliensbot/module/database"
	"aliens/testserver/module/game"
)

func init() {

}

func main() {
	aliens.Run(
		database.Module,
		game.Module,
	)
}
