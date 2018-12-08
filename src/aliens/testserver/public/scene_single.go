package main

import (
	"aliens/aliensbot"
	"aliens/aliensbot/module/database"
	"aliens/testserver/module/scene"
)

func init() {

}

func main() {

	aliens.Run(
		database.Module,
		scene.Module,
	)

}
