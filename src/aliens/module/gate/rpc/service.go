package rpc

import (
	"aliens/cluster/center"
	"golang.org/x/net/context"
	"aliens/module/gate/socket"
	"aliens/protocol"
	"aliens/module/gate/conf"
)

var gateRPCService *center.GRPCService = nil

func Init() {
	gateRPCService = center.PublicGRPCService(conf.Config.Service, &gateService{})
}

func Close() {
	if gateRPCService != nil {
		gateRPCService.Close()
	}
}

//只处理推送消息
type gateService struct {
}

func (this *gateService) Request(ctx context.Context,request *protocol.Any) (*protocol.Any, error) {
	//处理推送
	if request.SessionId != "" {
		request.Id = uint16(request.MessageId)
		socket.Push(request.SessionId, request)
	}
	return nil, nil
}