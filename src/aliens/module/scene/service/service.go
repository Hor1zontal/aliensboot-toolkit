package service

import (
	"google.golang.org/grpc"
	"aliens/cluster/center"
	"aliens/module/cluster"
	"aliens/module/scene/conf"
	"aliens/protocol/scene"
	"aliens/mmorpg"
	"aliens/module/scene/entity"
	"github.com/gogo/protobuf/proto"
	"aliens/module/scene/util"
	"golang.org/x/net/context"
)

//var Test1RPCService *center.gRPCService = nil
type RPCServiceServer interface {

}

type SceneService struct {
}



func (this *SceneService)Request(ctx context.Context,request *scene.SceneRequest) (*scene.SceneResponse, error) {

	response := &scene.SceneResponse{
		Session:request.Session,
	}
	if request.GetSpaceEnter() != nil {
		message := request.GetSpaceEnter()
		entity := mmorpg.SpaceManager.CreateEntity(message.GetSpaceID(), &entity.PlayerEntity{}, util.TransVector(message.GetPosition()), util.TransVector(message.GetDirection()))
		response.SpaceEnterRet = &scene.SpaceEnterRet{
			EntityID:proto.Int32(entity.GetID()),
		}
	} else if request.GetSpaceLeave() != nil {
		message := request.GetSpaceLeave()
		mmorpg.SpaceManager.LeaveEntity(message.GetSpaceID(), message.GetEntityID())
		response.SpaceLeaveRet = &scene.SpaceLeaveRet{

		}
	} else if request.GetSpaceMove() != nil {
		message := request.GetSpaceMove()
		mmorpg.SpaceManager.MoveEntity(message.GetSpaceID(), message.GetEntityID(), util.TransVector(message.GetPosition()), util.TransVector(message.GetDirection()))
		response.SpaceMoveRet = &scene.SpaceMoveRet{

		}
	} else if request.GetGetState() != nil {
		message := request.GetGetState()
		neighbors := mmorpg.SpaceManager.GetEntityState(message.GetSpaceID(), message.GetEntityID())
		response.GetStateRet = &scene.GetStateRet{
			Neighbors: util.BuildEntities(neighbors),
		}
	}

	return response, nil
	//log.Debug("call mmorpg : %v", request)
	//return &mmorpg.Response1{Response:proto.String(request.GetRequest())}, nil
}


func Init() {
	server := grpc.NewServer()
	scene.RegisterRPCServiceServer(server, &SceneService{})

	//s.RegisterService(&_RPCService_serviceDesc, srv)
	center.PublicRPCService(cluster.SERVICE_SCENE, conf.Config.RPCPort, server)
}

func Close() {

}