package aliens

import (
	"os"
	"os/signal"
	"aliens/config"
	"aliens/module"
	"aliens/log"
	"aliens/console"
)

func Run(mods ...module.Module) {

	log.Infof("AliensBot %v starting up", config.Version)
	// module
	for i := 0; i < len(mods); i++ {
		module.Register(mods[i])
	}
	module.Init()
	// console
	console.Init()

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	log.Infof("AliensBot closing down (signal: %v)", sig)
	console.Destroy()
	module.Destroy()
}
