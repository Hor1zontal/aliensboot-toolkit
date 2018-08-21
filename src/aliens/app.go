/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/8/21
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package aliens

import (
	"os"
	"os/signal"
	"aliens/config"
	"aliens/module"
	"aliens/log"
	"aliens/console"
	"fmt"
	"flag"
)

var (
	debug = false
)

func Run(mods ...module.Module) {
	flag.BoolVar(&debug, "debug", false, "debug flag")
	flag.Parse()
	log.SetDebug(debug)


	logo := `
	╔═║║  ╝╔═╝╔═ ╔═╝╔═ ╔═║═╔╝
	╔═║║  ║╔═╝║ ║══║╔═║║ ║ ║
	╝ ╝══╝╝══╝╝ ╝══╝══ ══╝ ╝
	`
	fmt.Println(logo)

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
