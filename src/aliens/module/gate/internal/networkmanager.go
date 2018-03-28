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
	skeleton.RegisterChanRPC("Newnetwork", rpcNewnetwork)
	skeleton.RegisterChanRPC("Closenetwork", rpcClosenetwork)
}

func rpcNewnetwork(args []interface{}) {
	a := args[0].(*network)
	//networkManager.AddNetwork(a)
	_ = a
}

func rpcClosenetwork(args []interface{}) {
	a := args[0].(*network)
	//networkManager.RemoveNetwork(a)
	//userdata := a.UserData()
	//a.SetUserData(nil)
	_ = a
}


type NetworkManager struct {
	networks  *collection.Map //存储所有的网络代理
	timeWheel *util.TimeWheel //验权检查时间轮
}


//开启心跳检查
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
	network := data[0].(*network)
	//超过固定时长没有验证权限需要提出
	if network.IsAuthTimeout() {
		//log.Debug("network heartbeat timeout : %v", network.RemoteAddr())
		network.Close()
		this.networks.Del(network)
	}
}

func (this *NetworkManager) AddNetwork(network *network) {
	this.networks.Set(network, &struct{}{})
	data := make(util.TaskData)
	data[0] = network
	this.timeWheel.AddTimer(time.Duration(conf.Config.AuthTimeout)*time.Second, network, data)
}

func (this *NetworkManager) RemoveNetwork(network *network) {
	this.networks.Del(network)
	this.timeWheel.RemoveTimer(network)
}





