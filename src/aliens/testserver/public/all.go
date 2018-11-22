package main

import (
	"aliens/aliensbot"
	"aliens/aliensbot/module/database"
	"aliens/testserver/module/game"
	"aliens/testserver/module/gate"
	"aliens/testserver/module/hall"
	"aliens/testserver/module/passport"
	"aliens/testserver/module/room"
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
