package socket

import (
	"time"
	"aliens/common/util"
	"aliens/common/collection"
	"aliens/module/gate/conf"
	"sync"
)

var networkManager = &NetworkManager{}

func init() {
	networkManager.Init()
}

func Push(clientID string, message interface{}) {
	networkManager.Push(clientID, message)
}

type NetworkManager struct {
	sync.RWMutex
	networks  *collection.Map //存储所有未验权的网络连接
	authNetworks map[string]*network  //存储所有验权通过的网络连接
	timeWheel *util.TimeWheel //验权检查时间轮
}

//开启权限,心跳等验证机制
func (this *NetworkManager) Init() {
	this.Lock()
	defer this.Unlock()
	if this.timeWheel != nil {
		this.timeWheel.Stop()
	}
	this.networks = &collection.Map{}
	this.authNetworks = make(map[string]*network)

	//心跳精确到s
	this.timeWheel = util.NewTimeWheel(time.Second, 60, this.dealAuthTimeout)
	this.timeWheel.Start()
}

func (this *NetworkManager) dealAuthTimeout(data util.TaskData) {
	//network := data[0].(*network)
	//超过固定时长没有验证权限需要提出
	//if network.IsAuthTimeout() {
	//	log.Debug("network auth timeout : %v", network.GetRemoteAddr())
	//	network.Close()
	//	this.networks.Del(network)
	//}
}

//验权限
func (this *NetworkManager) Auth(network *network) {
	this.Lock()
	defer this.Unlock()
	this.networks.Del(network)
	this.authNetworks[network.GetID()] = network
}

//推送消息
func (this *NetworkManager) Push(id string, message interface{}) {
	this.RUnlock()
	defer this.RUnlock()
	auth := this.authNetworks[id]
	if auth == nil {
		return
	}
	auth.SendMessage(message)
}

func (this *NetworkManager) AddNetwork(network *network) {
	data := make(util.TaskData)
	data[0] = network
	this.timeWheel.AddTimer(time.Duration(conf.Config.AuthTimeout)*time.Second, network, data)
	this.Lock()
	defer this.Unlock()
	this.networks.Set(network, &struct{}{})
}

func (this *NetworkManager) RemoveNetwork(network *network) {
	this.Lock()
	defer this.Unlock()
	this.timeWheel.RemoveTimer(network)
	this.networks.Del(network)
}