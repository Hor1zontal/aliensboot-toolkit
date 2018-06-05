/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 *
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package center

//服务中心，处理服务的调度和查询
import (
	"encoding/json"
	"github.com/samuel/go-zookeeper/zk"
	"sync"
	"time"
	"aliens/cluster/center/lbs"
	"gopkg.in/mgo.v2/bson"
	"aliens/log"
)

const NODE_SPLIT string = "/"

const SERVICE_NODE_NAME string = "service"

const DEFAULT_LBS string = lbs.LBS_STRATEGY_POLLING

type ZKServiceCenter struct {
	sync.RWMutex
	zkCon            *zk.Conn
	zkName           string
	serviceRoot      string
	serviceContainer map[string]*serviceCategory //服务容器 key 服务名

	nodeId  string //当前集群节点的id
	lbs string //default polling
	certFile string
	keyFile string
	commonName string
}

func (this *ZKServiceCenter) GetNodeID() string {
	return this.nodeId
}

//启动服务中心客户端
//func (this *ZKServiceCenter) Connect(address string, timeout int, zkName string, nodeID string) {
//	this.ConnectCluster([]string{address}, timeout, zkName, nodeID)
//}

func (this *ZKServiceCenter) ConnectCluster(config ClusterConfig) {
	if config.ID == "" {
		config.ID = bson.NewObjectId().Hex()
		//config.ID =
		//panic("cluster nodeID can not be empty")
	}
	if config.Timeout == 0 {
		config.Timeout = 10
	}
	this.lbs = config.LBS
	this.zkName = config.Name
	this.nodeId = config.ID
	this.certFile = config.CertFile
	this.keyFile = config.KeyFile
	this.commonName = config.CommonName
	//this.serviceFactory = serviceFactory
	c, _, err := zk.Connect(config.Servers, time.Duration(config.Timeout)*time.Second)
	if err != nil {
		panic(err)
	}
	this.serviceContainer = make(map[string]*serviceCategory)
	this.serviceRoot = NODE_SPLIT + this.zkName + NODE_SPLIT + SERVICE_NODE_NAME
	this.zkCon = c
	this.confirmNode(NODE_SPLIT + this.zkName)
	this.confirmNode(this.serviceRoot)
}

//func (this *ZKServiceCenter) SetLBS(lbs string) {
//	this.lbs = lbs
//}

func (this *ZKServiceCenter) IsConnect() bool {
	return this.zkCon != nil
}

func (this *ZKServiceCenter) assert() {
	if this.zkCon == nil {
		panic("mast start service center first")
	}
}

//关闭服务中心
func (this *ZKServiceCenter) Close() {
	if this.zkCon != nil {
		this.zkCon.Close()
	}
}

//更新服务
func (this *ZKServiceCenter) UpdateService(service IService) {
	this.Lock()
	defer this.Unlock()
	if this.serviceContainer[service.GetName()] == nil {
		this.serviceContainer[service.GetName()] = NewServiceCategory(service.GetName(), this.lbs, service.GetDesc())
	}
	this.serviceContainer[service.GetName()].updateService(service)
}

//根据服务类型获取一个空闲的服务节点
func (this *ZKServiceCenter) AllocService(serviceType string) IService {
	this.RLock()
	defer this.RUnlock()
	//TODO 后续要优化，考虑负载、空闲等因素
	serviceCategory := this.serviceContainer[serviceType]
	if serviceCategory == nil {
		return nil
	}
	return serviceCategory.allocService()
}

//
func (this *ZKServiceCenter) GetMasterService(serviceType string) IService {
	this.RLock()
	defer this.RUnlock()
	serviceCategory := this.serviceContainer[serviceType]
	if serviceCategory == nil {
		return nil
	}
	return serviceCategory.getMaster()
}

//func (this *ZKServiceCenter) CanHandle(name string, seq int32) bool {
//	serviceCategory := this.serviceContainer[name]
//	if serviceCategory == nil {
//		return false
//	}
//	return serviceCategory.canHandle(seq)
//}

func (this *ZKServiceCenter) GetService(serviceType string, serviceID string) IService {
	this.RLock()
	defer this.RUnlock()
	serviceCategory := this.serviceContainer[serviceType]
	if serviceCategory == nil {
		return nil
	}
	return serviceCategory.services[serviceID]
	////节点没有取第一个
	//if (service == nil) {
	//	serviceCategory.allocService()
	//}
	//return service
}

func (this *ZKServiceCenter) GetAllService(serviceType string) []IService {
	this.RLock()
	defer this.RUnlock()
	serviceCategory := this.serviceContainer[serviceType]
	if serviceCategory == nil {
		return nil
	}
	return serviceCategory.getAllService()
	////节点没有取第一个
	//if (service == nil) {
	//	serviceCategory.allocService()
	//}
	//return service
}

func (this *ZKServiceCenter) GetServiceInfo(serviceType string) []string {
	this.RLock()
	defer this.RUnlock()
	serviceCategory := this.serviceContainer[serviceType]
	if serviceCategory == nil {
		return nil
	}
	return serviceCategory.getNodes()
}

//订阅服务  能实时更新服务信息
func (this *ZKServiceCenter) SubscribeServices(serviceTypes ...string) {
	this.assert()
	for _, serviceType := range serviceTypes {
		this.SubscribeService(serviceType)
	}
}

func (this *ZKServiceCenter) SubscribeService(serviceName string) {
	path := this.serviceRoot + NODE_SPLIT + serviceName
	desc := this.confirmContentNode(path)
	serviceIDs, _, ch, err := this.zkCon.ChildrenW(path)
	if err != nil {
		log.Errorf("subscribe service %v error: %v", path, err)
		return
	}
	this.Lock()
	defer this.Unlock()
	oldContainer := this.serviceContainer[serviceName]
	serviceCategory := NewServiceCategory(serviceName, this.lbs, desc)
	for _, serviceID := range serviceIDs {
		data, _, err := this.zkCon.Get(path + NODE_SPLIT + serviceID)
		service := loadService(data, serviceID, serviceName)
		if service == nil {
			log.Errorf("%v unExpect service : %v", path, err)
			continue
		}
		if oldContainer != nil {
			oldService := oldContainer.takeoutService(service)
			if oldService != nil {
				oldService.SetID(service.GetID())
				serviceCategory.updateService(oldService)
				continue
			}
		}
		//新服务需要连接上才能更新
		if service.Connect() {
			serviceCategory.updateService(service)
		}
	}
	this.serviceContainer[serviceName] = serviceCategory
	go this.openListener(serviceName, path, ch)
}

func loadService(data []byte, id string, name string) IService {
	centerService := &centerService{}
	json.Unmarshal(data, centerService)
	centerService.SetID(id)
	centerService.SetName(name)
	switch centerService.Protocol {
	case GRPC:
		return &GRPCService{centerService: centerService}
	//case WEBSOCKET:
	//	return &wbService{centerService: centerService}
	//case HTTP:
	//	return &httpService{centerService: centerService}
	}
	return nil
}

func (this *ZKServiceCenter) openListener(serviceType string, path string, ch <-chan zk.Event) {
	event, _ := <-ch
	//更新服务节点信息
	if event.Type == zk.EventNodeChildrenChanged {
		this.SubscribeService(serviceType)
	}
}

//
func (this *ZKServiceCenter) confirmNode(path string, flags ...int32) bool {
	_, err := this.zkCon.Create(path, nil, 0, zk.WorldACL(zk.PermAll))
	return err == nil
}

func (this *ZKServiceCenter) confirmContentNode(path string, flags ...int32) string {
	_, err := this.zkCon.Create(path, nil, 0, zk.WorldACL(zk.PermAll))
	if err != nil {
		data, _, _ := this.zkCon.Get(path)
		return string(data)
	}
	return ""
}

func (this *ZKServiceCenter) confirmDataNode(path string, data []byte) bool {
	byteData := []byte(data)
	_, err := this.zkCon.Create(path, byteData, 0, zk.WorldACL(zk.PermAll))
	if err != nil {
		this.zkCon.Set(path, byteData, -1)
	}
	return err == nil
}

//发布服务
func (this *ZKServiceCenter) PublicService(service IService, unique bool) bool {
	this.assert()
	if !service.IsLocal() {
		log.Error("service info is invalid")
		return false
	}
	//path string, data []byte, version int32
	data, err := json.Marshal(service)
	if err != nil {
		log.Errorf("marshal json service data error : %v", err)
		return false
	}
	servicePath := this.serviceRoot + NODE_SPLIT + service.GetName()
	if unique {
		//TODO 可能有事务上的问题 需要优化
		child, _, _ := this.zkCon.Children(servicePath)
		if child != nil && len(child) > 0 {
			log.Errorf("unique service %v-%v already exist.", service.GetName(), child)
			return false
		}
	}

	this.confirmNode(servicePath)
	id, err := this.zkCon.Create(servicePath + NODE_SPLIT + service.GetID(), data,
		zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	if err != nil {
		log.Errorf("create service error : %v", err)
		return false
	}
	log.Infof("public %v success : %v", service.GetName(), id)
	//服务注册在容器
	this.UpdateService(service)
	return true
}