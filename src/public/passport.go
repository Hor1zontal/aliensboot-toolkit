package main

import (
	"aliens/module/passport"
	"aliens"
)

func init() {

}

func main() {

	aliens.Run(
		passport.Module,
	)
}
