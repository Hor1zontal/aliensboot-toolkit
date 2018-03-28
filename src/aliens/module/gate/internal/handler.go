package internal

import (
	"github.com/name5566/leaf/gate"
	"reflect"
	"aliens/module/gate/conf"
	"aliens/protocol/service1"
	"aliens/protocol/service2"
	"aliens/cluster/message"
	"github.com/name5566/leaf/network/protobuf"
	"aliens/module/cluster"
	"aliens/log"
	"aliens/cluster/center"
	"google.golang.org/grpc"
	"aliens/protocol/scene"
)

var router = make(map[reflect.Type]message.IMessageService)
//消息路由列表

var Processor = protobuf.NewProcessor()

func Init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	Processor.SetByteOrder(true)

	//TODO register router
	RegisterRouter(0, &scene.SceneRequest{}, 1, &scene.SceneResponse{}, cluster.SERVICE_SCENE, SceneFactory)
	RegisterRouter(2, &service1.Request1{}, 3, &service1.Response1{}, cluster.SERVICE_1, Service1Factory)
	RegisterRouter(4, &service2.Request2{}, 5, &service2.Response2{}, cluster.SERVICE_2, Service2Factory)
	//RegisterRouter(0, &mmorpg.Request1{}, 1, &mmorpg.Response1{}, message.NewRemoteService(cluster.SERVICE_1))
	//RegisterRouter(2, &service2.Request2{}, 3, &service2.Response2{}, message.NewRemoteService(cluster.SERVICE_2))
}

func SceneFactory(cc *grpc.ClientConn) interface{} {
	return scene.NewRPCServiceClient(cc)
}

func Service1Factory(cc *grpc.ClientConn) interface{} {
	return service1.NewRPCServiceClient(cc)
}

func Service2Factory(cc *grpc.ClientConn) interface{} {
	return service2.NewRPCServiceClient(cc)
}

//注册消息和服务映射关系
func RegisterRouter(requestID uint16, request interface{}, responseID uint16, response interface{},
				serviceID string, clientFactory func(cc *grpc.ClientConn) interface{}) {
	requestNo := Processor.Register(requestID, request)
	responseNo := Processor.Register(responseID, response)
	requestType := reflect.TypeOf(request)
	responseType := reflect.TypeOf(response)

	log.Debug("register request %v-%v  response %v-%v", requestNo, requestType, responseNo, responseType)
	center.RegisterRPCClientFactory(serviceID, clientFactory)

	router[requestType] = message.NewRemoteService(serviceID)
	skeleton.RegisterChanRPC(requestType, handleMessage)
}

type NetworkMessageHandler struct {
	networkChannel message.IMessageChannel
}

func (this *NetworkMessageHandler) HandleMessage(request interface{}) {
	requestType := reflect.TypeOf(request)
	messageService := router[requestType]
	if messageService == nil {
		log.Debug("unexpect request : %v", request)
		//TODO 返回错误信息，或者T人
		return
	}
	//response := reflect.New(responseType).Elem().Interface()
	response, error := messageService.HandleMessage(request)
	if error != nil {
		log.Debug("handle service error : %v", error)
		//TODO 返回错误信息，或者T人
		return
	}
	//log.Debug("service %v request : %v response : %v",messageService.GetType(), request, response)
	this.networkChannel.WriteMsg(response)
}

func handleMessage(args []interface{}) {
	request := args[0]
	// 消息的发送者
	gateAgent := args[1].(gate.Agent)
	userdata := gateAgent.UserData()
	switch userdata.(type) {
	case message.IChannelHandler:
		userdata.(message.IChannelHandler).AcceptMessage(request)
		break
	default:
		//打开缓存大小为5的收消息管道
		channelHandler := message.OpenChannelHandler(gateAgent, &NetworkMessageHandler{gateAgent}, conf.Config.MessageChannelLimit)
		gateAgent.SetUserData(channelHandler)
		channelHandler.AcceptMessage(request)
	}
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	userdata := a.UserData()
	//userdata.(message.IChannelHandler).AcceptMessage(&scene.SceneRequest{
	//	SpaceLeave:&scene.SpaceLeave{
	//		SpaceID:proto.Int32(),
	//		EntityID:
	//	},
	//})
	a.SetUserData(nil)
	switch userdata.(type) {
	case message.IChannelHandler:
		userdata.(message.IChannelHandler).GateClose(a)
		break
	}
	_ = a
}

