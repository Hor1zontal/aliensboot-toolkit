package main

import (
	"aliens/aliensbot"
	"aliens/aliensbot/module/database"
	"aliens/testserver/module/passport"
)

func init() {

}

func main() {

	aliens.Run(
		database.Module,
		passport.Module,
	)
}
