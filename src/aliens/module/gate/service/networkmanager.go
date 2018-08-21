package service

import (
	"time"
	"aliens/common/util"
	"aliens/common/data_structures/map"
	"aliens/module/gate/conf"
	"aliens/module/cluster/cache"
	"aliens/cluster/center"
)

var Manager = &NetworkManager{}

func init() {
	Manager.Init()
}

type NetworkManager struct {
	//sync.RWMutex
	networks  *_map.Map             //存储所有未验权的网络连接
	authNetworks map[int64]*Network //存储所有验权通过的网络连接
	timeWheel *util.TimeWheel       //验权检查时间轮
}

//开启权限,心跳等验证机制
func (this *NetworkManager) Init() {
	if this.timeWheel != nil {
		this.timeWheel.Stop()
	}
	this.networks = &_map.Map{}
	this.authNetworks = make(map[int64]*Network)

	//心跳精确到s
	this.timeWheel = util.NewTimeWheel(time.Second, 60, this.dealAuthTimeout)
	this.timeWheel.Start()
}

func (this *NetworkManager) dealAuthTimeout(data util.TaskData) {
	//Network := data[0].(*Network)
	//超过固定时长没有验证权限需要退出
	//if Network.IsAuthTimeout() {
	//	log.Debug("Network auth timeout : %v", Network.GetRemoteAddr())
	//	Network.Close()
	//	this.networks.Del(Network)
	//}
}

//验权限
func (this *NetworkManager) auth(authID int64, network *Network) {
	this.timeWheel.RemoveTimer(network)
	this.networks.Del(network)
	this.authNetworks[authID] = network
	cache.ClusterCache.SetAuthGateID(authID, center.ClusterCenter.GetNodeID())
}

//推送消息
func (this *NetworkManager) push(id int64, message interface{}) {
	auth := this.authNetworks[id]
	if auth == nil {
		return
	}
	auth.SendMessage(message)
}

func (this *NetworkManager) AddNetwork(network *Network) {
	data := make(util.TaskData)
	data[0] = network
	this.timeWheel.AddTimer(time.Duration(conf.Config.AuthTimeout)*time.Second, network, data)
	this.networks.Set(network, &struct{}{})
}

func (this *NetworkManager) RemoveNetwork(network *Network) {
	if network.IsAuth() {
		delete(this.authNetworks, network.authID)
		cache.ClusterCache.CleanAuthGateID(network.authID)
	} else {
		this.timeWheel.RemoveTimer(network)
		this.networks.Del(network)
	}
}