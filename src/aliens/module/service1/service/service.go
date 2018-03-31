package service

import (
	"aliens/cluster/center"
	"aliens/module/cluster"
	"aliens/module/service1/conf"
	"google.golang.org/grpc"
	"aliens/protocol/service1"
	"aliens/log"
	"golang.org/x/net/context"
	"github.com/gogo/protobuf/proto"
)

//var Test1RPCService *center.gRPCService = nil
type RPCServiceServer interface {

}

type Service struct {
}

func (this *Service)Request(ctx context.Context,request *service1.Request1) (*service1.Response1, error) {
	log.Debug("call mmorpg : %v", request)
	return &service1.Response1{Response:proto.String(request.GetRequest())}, nil
}

//type ServiceP struct {
//}
//
//func (this *ServiceP) Request(server mmorpg.RPCService_RequestServer) error {
////	log.Debug("call mmorpg : %v", request)
////	return &mmorpg.Response1{Response:proto.String(request.GetRequest())}, nil
//	//server.se nd
//	request, _ := server.Recv()
//	log.Debug("call mmorpg : %v", request)
//	server.Send(&mmorpg.Response1{Response:proto.String(request.GetRequest())})
//
//	return  nil
//}


func Init() {
	server := grpc.NewServer()
	service1.RegisterRPCServiceServer(server, &Service{})
	center.PublicGRPCService(cluster.SERVICE_1, conf.Config.RPCPort, server)
}

func Close() {

}