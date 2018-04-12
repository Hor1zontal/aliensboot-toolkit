package service

import (
	"aliens/cluster/center"
	"aliens/module/scene/conf"
)

var sceneRPCService *center.GRPCService = nil

func Init() {
	sceneRPCService = center.PublicGRPCService(conf.Config.Service, &sceneService{})
}

func Close() {
	if sceneRPCService != nil {
		sceneRPCService.Close()
	}
}