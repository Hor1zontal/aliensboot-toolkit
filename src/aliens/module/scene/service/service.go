package service

import (
	"aliens/cluster/center"
	"aliens/module/scene/conf"
)

var sceneRPCService *center.GRPCService = nil

func Init() {
	sceneRPCService = center.PublicGRPCService(conf.Config.Service, conf.Config.RPCPort, &sceneService{})
}

func Close() {
	if sceneRPCService != nil {
		sceneRPCService.Close()
	}
}