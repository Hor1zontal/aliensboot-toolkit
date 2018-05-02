/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/4/11
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package cache

const (
	userGatePrefix string = "ug_"     //
)

//设置客户端所在的网关id
//func (this *cacheManager) SetClientGateID(clientID string, gateID string) bool {
//	return this.redisClient.SetData(clientID, gateID)
//}
//
////清楚用户和网关的对应关系
//func (this *cacheManager) CleanClientGateID(clientID string, gateID string) bool {
//	return this.redisClient.DelData(clientID)
//}
//
////获取客户端所在的网关id
//func (this *cacheManager) GetClientGateID(clientID string) string {
//	return this.redisClient.GetData(clientID)
//}
