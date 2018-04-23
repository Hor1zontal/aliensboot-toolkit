package service

import (
	"aliens/module/gate/route"
	"aliens/module/gate/socket"
	"aliens/module/cluster/dispatch"
	"aliens/module/gate/conf"
	"aliens/protocol"
)

//var gateRPCService *center.GRPCService = nil

func Init() {
	dispatch.RegisterConsumer(conf.Config.Service.Name, HandlePush)
	//gateRPCService = center.PublicGRPCService(conf.Config.Service, &gateService{})
}

func Close() {
	dispatch.UnregisterConsumer(conf.Config.Service.Name)
}

//只处理推送消息
func HandlePush(request *protocol.Any) error {
	if request.SessionId != "" {
		request.Id = route.GetPushID(request.TypeUrl)
		socket.Push(request.SessionId, request)
	}
	return nil
}

//type gateService struct {
//}
//
//func (this *gateService) Request(ctx context.Context,request *protocol.Any) (*protocol.Any, error) {
//	//处理推送
//	if request.SessionId != "" {
//		request.Id = route.GetPushID(request.TypeUrl)
//		socket.Push(request.SessionId, request)
//	}
//	return nil, nil
//}
