package network

import (
	"aliens/common/util"
	modulebase "aliens/module/base"
	"aliens/common/data_structures/set"
	"aliens/protocol/base"
	"aliens/protocol"
	"aliens/cluster/center"
	"aliens/module/gate/cache"
	"aliens/module/dispatch/rpc"
)

var Manager = &networkManager{}

//const (
//	CommandRpcResponse = "resp"
//)

type networkManager struct {
	//handler *modulebase.Skeleton
	networks  *set.HashSet          //存储所有未验权的网络连接
	authNetworks map[int64]*Network //存储所有验权通过的网络连接
	node string //当前节点名
	//timeWheel *util.TimeWheel       //验权检查时间轮
}

var handler *modulebase.Skeleton

func Init(skeleton *modulebase.Skeleton) {
	handler = skeleton
	Manager.Init()
}

//开启权限,心跳等验证机制
func (this *networkManager) Init() {
	///this.handler = chanRpc
	//this.chanRpc.Register(CommandRpcResponse, this.handleResponse)
	this.node = center.ClusterCenter.GetNodeID()
	this.networks = set.NewHashSet()
	this.authNetworks = make(map[int64]*Network)

	//心跳精确到s
	//this.timeWheel = util.NewTimeWheel(time.Second, 60, this.dealAuthTimeout)
	//this.timeWheel.Start()
}
//
//func (this *networkManager) AsyncCall(f func(), c func()) {
//	this.handler.Go(f, c)
//}

func (this *networkManager) handleResponse(args []interface{}) {
	network := args[0].(*Network)
	response := args[1].(*base.Any)
	err, ok := args[2].(error)
	if ok {
		network.handleResponse(response, err)
	} else {
		network.handleResponse(response, nil)
	}
}

func (this *networkManager) BindService(authID int64, binds map[string]string) {
	auth := this.authNetworks[authID]
	if auth == nil {
		return
	}
	auth.BindService(binds)
}

func (this *networkManager) KickOut(authID int64, kickType protocol.KickType) {
	auth := this.authNetworks[authID]
	if auth == nil {
		return
	}
	auth.KickOut(kickType)
}

//推送消息
func (this *networkManager) Push(authID int64, message *base.Any) {
	auth := this.authNetworks[authID]
	if auth == nil {
		return
	}
	auth.Push(message)
}

//广播消息
func (this *networkManager)	Broadcast(message *base.Any) {
	for _, network := range this.authNetworks {
		network.Push(message)
	}
}

func (this *networkManager) AddNetwork(network *Network) {
	data := make(util.TaskData)
	data[0] = network
	//this.timeWheel.AddTimer(time.Duration(conf.Config.AuthTimeout)*time.Second, network, data)
	this.networks.Add(network)
}

func (this *networkManager) RemoveNetwork(network *Network) {
	if network.IsAuth() {
		delete(this.authNetworks, network.authID)
		cache.ClusterCache.CleanAuthGateID(network.authID)
	} else {
		//this.timeWheel.RemoveTimer(network)
		this.networks.Remove(network)
	}

}

func (this *networkManager) DealAuthTimeout() {
	this.networks.Range(func (element interface{}) {
		network := element.(*Network)
		//连接超过固定时长没有验证权限需要退出
		if network.IsAuthTimeout() {
			//log.Debug("Network auth timeout : %v", networker.GetRemoteAddr())
			network.KickOut(protocol.KickType_Timeout)
			this.networks.Remove(network)
		}
	})
}

//验权限
func (this *networkManager) auth(authID int64, network *Network) {
	//this.timeWheel.RemoveTimer(network)
	this.networks.Remove(network)

	oldNetwork, ok := this.authNetworks[authID]

	//顶号处理
	if ok {
		oldNetwork.KickOut(protocol.KickType_OtherSession)
	} else {
		node := cache.ClusterCache.GetAuthGateID(authID)
		//用户在其他网关节点登录 需要发送远程T人
		if node != this.node {
			kickMsg := &protocol.KickOut{
				AuthID:authID,
				KickType:protocol.KickType_OtherSession,
			}
			rpc.Gate.KickOut(node, kickMsg)
		}
	}
	cache.ClusterCache.SetAuthGateID(authID, this.node)
	this.authNetworks[authID] = network
	//cache.ClusterCache.SetAuthGateID(authID, center.ClusterCenter.GetNodeID())
}