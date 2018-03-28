package center

import (

	"gopkg.in/mgo.v2/bson"
)

var ClusterCenter *ServiceCenter = &ServiceCenter{} //服务调度中心

var nodeName = bson.NewObjectId().Hex()


func IsMaster(serviceType string) bool {
	service := ClusterCenter.GetMasterService(serviceType)
	if service == nil {
		return false
	}
	return service.GetID() == GetServerNode()
}

func GetServerNode() string {
	return nodeName
}

