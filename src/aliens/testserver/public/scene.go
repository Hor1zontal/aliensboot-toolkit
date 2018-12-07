package main

import (
	"aliens/aliensbot"
	"aliens/aliensbot/module/database"
	"aliens/testserver/module/game"
	"aliens/testserver/module/gate"
	"aliens/testserver/module/passport"
	"aliens/testserver/module/scene"
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
