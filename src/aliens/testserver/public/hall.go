package main

import (
	"aliens/aliensbot"
	"aliens/aliensbot/module/database"
	"aliens/testserver/module/hall"
)

func init() {

}

func main() {

	aliens.Run(
		database.Module,
		hall.Module,
	)
}
