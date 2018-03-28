package internal

import (
	"github.com/name5566/leaf/gate"
	"time"
	"aliens/log"
	"aliens/common/util"
	"aliens/common/collection"
)

var agentManager = &AgentManager{}

func init() {
	agentManager.Init(60)
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	//agentManager.AddAgent(a)
	_ = a
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	//agentManager.RemoveAgent(a)
	//userdata := a.UserData()
	//a.SetUserData(nil)
	_ = a
}


type AgentManager struct {
	agents    *collection.Map //存储所有的网络代理
	timeWheel *util.TimeWheel //心跳检查时间轮
	timeout   float64
}


//开启心跳检查
func (this *AgentManager) Init(timeout float64) {
	if this.timeWheel != nil {
		this.timeWheel.Stop()
	}
	this.agents = &collection.Map{}
	this.timeout = timeout

	//心跳精确到s
	this.timeWheel = util.NewTimeWheel(1*time.Second, 60, this.dealHeartbeat)
	this.timeWheel.Start()
}

func (this *AgentManager) dealHeartbeat(data util.TaskData) {
	agent := data[0].(gate.Agent)
	if agent.IsHearbeatTimeout(this.timeout) {
		log.Debug("agent heartbeat timeout : %v", agent.RemoteAddr())
		agent.Close()
		this.agents.Del(agent)
	}
}

func (this *AgentManager) AddAgent(agent gate.Agent) {
	//id := util.GenUUID()
	//this.agents.Set(id, agent)
	//agent.SetUserData(id)
	this.agents.Set(agent, &struct{}{})
	data := make(util.TaskData)
	data[0] = agent
	this.timeWheel.AddTimer(time.Duration(this.timeout)*time.Second, agent, data)
}

func (this *AgentManager) RemoveAgent(agent gate.Agent) {
	this.agents.Del(agent)
	this.timeWheel.RemoveTimer(agent)
}





