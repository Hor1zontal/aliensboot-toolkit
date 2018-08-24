package module

import (
	"sync"
	"aliens/exception"
	"aliens/config"
)

type Module interface {
	GetName() string  //module name
	GetConfig() interface{}  //module config data

	OnInit()
	OnDestroy()
	Run(closeSig chan bool)
}

type module struct {
	mi       Module
	closeSig chan bool
	wg       sync.WaitGroup
}

var mods []*module

func Register(mi Module) {
	m := new(module)
	m.mi = mi
	m.closeSig = make(chan bool, 1)
	mods = append(mods, m)
}

func Init() {
	for i := 0; i < len(mods); i++ {
		config.LoadConfigData(mods[i].mi.GetName(), mods[i].mi.GetConfig())
	}

	for i := 0; i < len(mods); i++ {
		mods[i].mi.OnInit()
	}

	for i := 0; i < len(mods); i++ {
		m := mods[i]
		m.wg.Add(1)
		go run(m)
	}
}

func Destroy() {
	for i := len(mods) - 1; i >= 0; i-- {
		m := mods[i]
		m.closeSig <- true
		m.closeSig <- true
		m.wg.Wait()
		destroy(m)
	}
}

func run(m *module) {
	m.mi.Run(m.closeSig)
	m.wg.Done()
}

func destroy(m *module) {
	defer func() {
		exception.CatchStackDetail()
	}()
	m.mi.OnDestroy()
}
