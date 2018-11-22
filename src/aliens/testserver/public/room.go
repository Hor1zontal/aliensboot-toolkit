package main

import (
	"aliens/aliensbot"
	"aliens/aliensbot/module/database"
	"aliens/testserver/module/room"
)

func init() {

}

func main() {

	aliens.Run(
		database.Module,
		room.Module,
	)
}
