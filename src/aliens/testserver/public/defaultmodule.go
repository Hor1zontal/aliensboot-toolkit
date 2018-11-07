package main

import (
	"aliens/aliensbot"
	"aliens/aliensbot/module/database"
	"aliens/testserver/module/defaultmodule"
)

func init() {

}

func main() {

	aliens.Run(
		database.Module,
		defaultmodule.Module,
	)
}
