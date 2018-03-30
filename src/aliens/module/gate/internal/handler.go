package internal

import (
	"github.com/name5566/leaf/gate"
	"aliens/cluster/message"
	"aliens/module/cluster"
	"aliens/module/gate/msg"
	"reflect"
	"github.com/gogo/protobuf/types"
)

var router = make(map[uint16]message.IMessageService)

var Processor = msg.NewMsgProcessor() //protobuf.NewProcessor()

func Init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	skeleton.RegisterChanRPC(reflect.TypeOf(&types.Any{}), handleMessage)

	Processor.Register(0, "aliens/scene.SceneRequest")
	Processor.Register(1, "aliens/scene.SceneResponse")

	//types.UnmarshalAny()
	Processor.SetByteOrder(true)

	//types.UnmarshalAny()


	//&types.Any{TypeUrl: "type.googleapis.com/scene.SceneRequest", Value:   }, nil

	//TODO register router
	RegisterRouter(0, 1, cluster.SERVICE_SCENE)
	//RegisterRouter(2,3, cluster.SERVICE_1)
	//RegisterRouter(4, 5,  cluster.SERVICE_2)
	//RegisterRouter(0, &mmorpg.Request1{}, 1, &mmorpg.Response1{}, message.NewRemoteService(cluster.SERVICE_1))
	//RegisterRouter(2, &service2.Request2{}, 3, &service2.Response2{}, message.NewRemoteService(cluster.SERVICE_2))
}

//注册消息和服务映射关系
func RegisterRouter(requestID uint16, responseID uint16, serviceID string) {
	//Processor.Register(responseID, &any.Any{})
	//requestType := reflect.TypeOf(request)
	//responseType := reflect.TypeOf(response)

	//log.Debug("register request %v-%v  response %v-%v", requestNo, requestType, responseNo, responseType)
	//center.RegisterRPCClientFactory(serviceID, clientFactory)


	//"type.googleapis.com/scene.SceneRequest"
	router[requestID] = message.NewRemoteService(serviceID)

}


func rpcNewAgent(args []interface{}) {
	agent := args[0].(gate.Agent)
	if agent.UserData() == nil {
		//打开缓存大小为5的收消息管道
		network := newNetwork(agent)
		agent.SetUserData(network)
		networkManager.AddNetwork(network)
	}

	_ = agent
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	networkManager.RemoveNetwork(a.UserData().(*network))
	//userdata := a.UserData()
	//a.SetUserData(nil)
	_ = a
}


func handleMessage(args []interface{}) {
	request := args[0]
	// 消息的发送者
	gateAgent := args[1].(gate.Agent)
	userdata := gateAgent.UserData()
	switch userdata.(type) {
		case *network:
			userdata.(*network).AcceptMessage(request)
			break
	}
}