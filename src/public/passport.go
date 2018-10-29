package main

import (
	"aliens"
	"aliens/module/passport"
)

func init() {

}

func main() {

	aliens.Run(
		passport.Module,
	)
}
