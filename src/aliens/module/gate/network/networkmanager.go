package network

import (
	"time"
	"aliens/common/util"
	"aliens/module/gate/conf"
	"aliens/common/data_structures/set"
	"aliens/chanrpc"
	"aliens/protocol/base"
)

var Manager = &networkManager{}

const (
	CommandRpcResponse = "resp"
)

type networkManager struct {
	chanRpc *chanrpc.Server
	networks  *set.HashSet             //存储所有未验权的网络连接
	authNetworks map[int64]*Network //存储所有验权通过的网络连接
	timeWheel *util.TimeWheel       //验权检查时间轮
}

//开启权限,心跳等验证机制
func (this *networkManager) Init(chanRpc *chanrpc.Server) {
	this.chanRpc = chanRpc
	this.chanRpc.Register(CommandRpcResponse, this.handleResponse)
	if this.timeWheel != nil {
		this.timeWheel.Stop()
	}
	this.networks = set.NewHashSet()
	this.authNetworks = make(map[int64]*Network)

	//心跳精确到s
	this.timeWheel = util.NewTimeWheel(time.Second, 60, this.dealAuthTimeout)
	this.timeWheel.Start()
}

func (this *networkManager) acceptResponse(network *Network, response *base.Any, err error) {
	this.chanRpc.Go(CommandRpcResponse, network, response, err)
}


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

//推送消息
func (this *networkManager) Push(id int64, message interface{}) {
	auth := this.authNetworks[id]
	if auth == nil {
		return
	}
	auth.SendMessage(message)
}

func (this *networkManager) AddNetwork(network *Network) {
	data := make(util.TaskData)
	data[0] = network
	this.timeWheel.AddTimer(time.Duration(conf.Config.AuthTimeout)*time.Second, network, data)
	this.networks.Add(network)
}

func (this *networkManager) RemoveNetwork(network *Network) {
	if network.IsAuth() {
		delete(this.authNetworks, network.authID)
		//cache.ClusterCache.CleanAuthGateID(network.authID)
	} else {
		this.timeWheel.RemoveTimer(network)
		this.networks.Remove(network)
	}
}

func (this *networkManager) dealAuthTimeout(data util.TaskData) {
	//Network := data[0].(*Network)
	//超过固定时长没有验证权限需要退出
	//if Network.IsAuthTimeout() {
	//	log.Debug("Network auth timeout : %v", Network.GetRemoteAddr())
	//	Network.Close()
	//	this.networks.Del(Network)
	//}
}

//验权限
func (this *networkManager) auth(authID int64, network *Network) {
	//TODO T人
	this.timeWheel.RemoveTimer(network)
	this.networks.Remove(network)
	this.authNetworks[authID] = network
	//cache.ClusterCache.SetAuthGateID(authID, center.ClusterCenter.GetNodeID())
}