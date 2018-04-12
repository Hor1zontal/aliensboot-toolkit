package cluster

import "aliens/module/cluster/conf"

const (
	SERVICE_SCENE = "scene"
	SERVICE_PASSPORT = "scene"
	SERVICE_GATE = "gate"
)


func GetID() string {
	return conf.NodeName
}

