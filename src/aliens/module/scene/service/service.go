package service

import (
	"aliens/cluster/center"
	"aliens/module/scene/conf"
)




//func (this *SceneService)HandleRequest(ctx context.Context,request *scene.SceneRequest) (*scene.SceneResponse, error) {
//
//
//	response := &scene.SceneResponse{
//		Session:request.Session,
//	}
//	//request.re
//	//types.MarshalAny(&scene.SceneResponse)
//
//	if request.GetSpaceEnter() != nil {
//		message := request.GetSpaceEnter()
//		entity := mmorpg.SpaceManager.CreateEntity(message.GetSpaceID(), &entity.PlayerEntity{}, util.TransVector(message.GetPosition()), util.TransVector(message.GetDirection()))
//		response.SpaceEnterRet = &scene.SpaceEnterRet{
//			EntityID:entity.GetID(),
//		}
//	} else if request.GetSpaceLeave() != nil {
//		message := request.GetSpaceLeave()
//		mmorpg.SpaceManager.LeaveEntity(message.GetSpaceID(), message.GetEntityID())
//		response.Response.SpaceLeaveRet = &scene.SpaceLeaveRet{
//
//		}
//	} else if request.GetSpaceMove() != nil {
//		message := request.GetSpaceMove()
//		mmorpg.SpaceManager.MoveEntity(message.GetSpaceID(), message.GetEntityID(), util.TransVector(message.GetPosition()), util.TransVector(message.GetDirection()))
//		response.SpaceMoveRet = &scene.SpaceMoveRet{
//
//		}
//	} else if request.GetGetState() != nil {
//		message := request.GetGetState()
//		neighbors := mmorpg.SpaceManager.GetEntityState(message.GetSpaceID(), message.GetEntityID())
//		response.GetStateRet = &scene.GetStateRet{
//			Neighbors: util.BuildEntities(neighbors),
//		}
//	}
//
//	return response, nil
//	//log.Debug("call mmorpg : %v", request)
//	//return &mmorpg.Response1{Response:proto.String(request.GetRequest())}, nil
//}


func Init() {
	//s.RegisterService(&_RPCService_serviceDesc, srv)
	center.PublicGRPCService(conf.Config.Service, conf.Config.RPCPort, &sceneService{})
}

func Close() {

}