package main

import (
	"aliens/aliensbot"
	"aliens/aliensbot/module/database"
	"aliens/testserver/module/match"
)

func init() {

}

func main() {

	aliens.Run(
		database.Module,
		match.Module,
	)
}
