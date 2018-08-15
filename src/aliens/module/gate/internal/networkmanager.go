package internal

import (
	"time"
	"aliens/common/util"
	"aliens/common/data_structures/map"
	"aliens/module/gate/conf"
	"aliens/module/cluster/cache"
	"aliens/cluster/center"
)

var networkManager = &NetworkManager{}

func init() {
	networkManager.Init()
}

type NetworkManager struct {
	//sync.RWMutex
	networks  *_map.Map   //存储所有未验权的网络连接
	authNetworks map[int64]*network //存储所有验权通过的网络连接
	timeWheel *util.TimeWheel       //验权检查时间轮
}

//开启权限,心跳等验证机制
func (this *NetworkManager) Init() {
	if this.timeWheel != nil {
		this.timeWheel.Stop()
	}
	this.networks = &_map.Map{}
	this.authNetworks = make(map[int64]*network)

	//心跳精确到s
	this.timeWheel = util.NewTimeWheel(time.Second, 60, this.dealAuthTimeout)
	this.timeWheel.Start()
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

//验权限
func (this *NetworkManager) auth(authID int64, network *network) {
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
	auth.WriteMsg(message)
}

func (this *NetworkManager) addNetwork(network *network) {
	data := make(util.TaskData)
	data[0] = network
	this.timeWheel.AddTimer(time.Duration(conf.Config.AuthTimeout)*time.Second, network, data)
	this.networks.Set(network, &struct{}{})
}

func (this *NetworkManager) removeNetwork(network *network) {
	network.Close()
	if network.IsAuth() {
		delete(this.authNetworks, network.authID)
		cache.ClusterCache.CleanAuthGateID(network.authID)
	} else {
		this.timeWheel.RemoveTimer(network)
		this.networks.Del(network)
	}
}