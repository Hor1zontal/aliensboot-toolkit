package main

import (
	"aliens/aliensbot"
	"aliens/testserver/module/passport"
	"aliens/aliensbot/module/database"
)

func init() {

}

func main() {

	aliens.Run(
		database.Module,
		passport.Module,
	)
}
