package internal

import (
	"time"
	"aliens/common/util"
	"aliens/common/collection"
	"aliens/module/gate/conf"
)

var networkManager = &NetworkManager{}

func init() {
	networkManager.Init()
}


type NetworkManager struct {
	networks  *collection.Map //存储所有的网络代理
	timeWheel *util.TimeWheel //验权检查时间轮
}

//开启权限,心跳等验证机制
func (this *NetworkManager) Init() {
	if this.timeWheel != nil {
		this.timeWheel.Stop()
	}
	this.networks = &collection.Map{}

	//心跳精确到s
	this.timeWheel = util.NewTimeWheel(1*time.Second, 60, this.dealAuthTimeout)
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

func (this *NetworkManager) AddNetwork(network *network) {
	data := make(util.TaskData)
	data[0] = network
	this.timeWheel.AddTimer(time.Duration(conf.Config.AuthTimeout)*time.Second, network, data)
	this.networks.Set(network, &struct{}{})
}

func (this *NetworkManager) RemoveNetwork(network *network) {
	this.timeWheel.RemoveTimer(network)
	this.networks.Del(network)
}





