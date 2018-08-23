package service

import (
	"aliens/common/util"
	"aliens/common/data_structures/map"
	"github.com/AsynkronIT/protoactor-go/actor"
	"aliens/gate"
	"aliens/chanrpc"
)

var manager = &NetworkManager{}

var chanRpc *chanrpc.Server = nil

const (
	CommandAgentRemote = "removeAgent"
	CommandAgentAuth = "authAgent"
)

type NetworkManager struct {
	//sync.RWMutex
	networks  *_map.Map               //存储所有未验权的网络连接
	authNetworks map[int64]*actor.PID //存储所有验权通过的网络连接
	//timeWheel *util.TimeWheel       //验权检查时间轮
}

//开启权限,心跳等验证机制
func (this *NetworkManager) Init(chanRpc *chanrpc.Server) {
	//if this.timeWheel != nil {
	//	this.timeWheel.Stop()
	//}
	chanRpc = chanRpc
	this.networks = &_map.Map{}
	this.authNetworks = make(map[int64]*actor.PID)

	chanRpc.Register(CommandAgentRemote, this.removeAgent)
	chanRpc.Register(CommandAgentAuth, this.authAgent)
	chanRpc.Register(gate.CommandAgentNew, this.newAgent)
	chanRpc.Register(gate.CommandAgentClose, this.closeAgent)
	chanRpc.Register(gate.CommandAgentMsg, this.handleMessage)

	//心跳精确到s
	//this.timeWheel = util.NewTimeWheel(time.Second, 60, this.dealAuthTimeout)
	//this.timeWheel.Start()
}

func (this *NetworkManager) dealAuthTimeout(data util.TaskData) {
	//network := data[0].(*network)
	//超过固定时长没有验证权限需要退出
	//if network.IsAuthTimeout() {
	//	log.Debug("network auth timeout : %v", network.GetRemoteAddr())
	//	network.Close()
	//	this.networks.Del(network)
	//}
}

//推送消息给指定玩家
func (this *NetworkManager) push(id int64, message interface{}) {
	auth := this.authNetworks[id]
	if auth == nil {
		//TODO 推送的玩家不在线的处理方式
		return
	}
	auth.Tell(&userPush{pushMsg:message})
}

//
func (this *NetworkManager) broadcast(message interface{}) {
	msg := &userPush{pushMsg:message}
	for _, network := range this.authNetworks {
		network.Tell(msg)
	}
}

//新的agent处理
func (this *NetworkManager) newAgent(args []interface{}) {
	agent := args[0].(gate.Agent)
	if agent.UserData() == nil {
		pid := actor.Spawn(actor.FromProducer(func() actor.Actor {return &network{}}))
		//data := make(util.TaskData)
		//data[0] = network
		//this.timeWheel.AddTimer(time.Duration(conf.Config.AuthTimeout)*time.Second, network, data)
		this.networks.Set(pid, &struct{}{})
		pid.Tell(&NetworkInit{agent:agent,pid:pid})
		agent.SetUserData(pid)
	}
}

//agent授权
func (this *NetworkManager) authAgent(args []interface{}) {
	authID := args[0].(int64)
	pid := args[1].(*actor.PID)
	this.networks.Del(pid)
	this.authNetworks[authID] = pid
}

//由network 处理释放流程后主动通知
func (this *NetworkManager) removeAgent(args []interface{}) {
	authID := args[0].(int64)
	pid := args[1].(*actor.PID)
	if authID != 0 {
		delete(this.authNetworks, authID)
	} else {
		this.networks.Del(pid)
	}
}

//关闭连接处理
func (this *NetworkManager) closeAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	pid, ok := a.UserData().(*actor.PID)
	a.SetUserData(nil)
	if ok {
		pid.Stop()
	}
}

//消息处理
func (this *NetworkManager) handleMessage(args []interface{}) {
	request := args[0]
	//消息的发送者
	gateAgent := args[1].(gate.Agent)
	pid, ok := gateAgent.UserData().(*actor.PID)
	if ok {
		pid.Tell(request)
	}
}



