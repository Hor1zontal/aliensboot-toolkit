/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
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
	"aliens/cluster/center"
)

var (
	debug = false
	configPath = ""  //配置文件根目录，默认当前
	tag = ""
)

func init() {
	flag.BoolVar(&debug, "debug", false, "debug flag")
	flag.StringVar(&configPath, "config", "config", "configuration path")
	flag.StringVar(&tag, "tag", "aliens", "log tag")
	flag.Parse()

}

func Run(mods ...module.Module) {
	baseConfig := config.Init(configPath)

	log.Init(debug, tag, baseConfig.PathLog)

	center.ClusterCenter.ConnectCluster(baseConfig.Cluster)

	logo := `
	╔═║║  ╝╔═╝╔═ ╔═╝╔═ ╔═║═╔╝
	╔═║║  ║╔═╝║ ║══║╔═║║ ║ ║
	╝ ╝══╝╝══╝╝ ╝══╝══ ══╝ ╝
	`
	fmt.Println(logo)

	log.Infof("AliensBot %v starting up", config.Version)

	//module.Register(database.Module)

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
	//close cluster
	center.ClusterCenter.Close()
}
