package service

import (
	"aliens/cluster/center"
	"aliens/module/cluster"
	"aliens/module/service2/conf"
	"google.golang.org/grpc"
	"aliens/protocol/service2"
	"aliens/log"
	"golang.org/x/net/context"
	"github.com/gogo/protobuf/proto"
)

//var Test1RPCService *center.gRPCService = nil
type RPCServiceServer interface {

}

type Service struct {
}

func (this *Service)Request(ctx context.Context, request *service2.Request2) (*service2.Response2, error) {
	log.Debug("call service2 : %v", request)
	return &service2.Response2{Response:proto.String(request.GetRequest())}, nil
}


func Init() {
	server := grpc.NewServer()
	service2.RegisterRPCServiceServer(server, &Service{})
	center.PublicRPCService(cluster.SERVICE_2, conf.Config.RPCPort, server)
}

func Close() {

}