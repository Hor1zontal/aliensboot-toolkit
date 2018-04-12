/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/4/12
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package cache

import (
	"aliens/common/util"
	"aliens/common/database"
)

const (
	USER_KEY_PREFIX string = "gu"
	USERNAME_KEY_PREFIX string = "uk_"

	UPROP_DESC string = "desc"         //用户签名
	UPROP_NICKNAME string = "nname"         //用户昵称
	UPROP_ICON string = "icon"      	//图标
	UPROP_ONLINE string = "online"		//用户是否登录
	UPROP_ONLINE_TIME string = "otime"  //用户上次的登录时间戳
	UPROP_INTO_EVENT string = "inevent" //用户当前进入的事件副本
	UPROP_FAITH string = "faith"      	//用户信仰值
	UPROP_POWER string = "power"      	//用户法力值
	UPROP_POWER_LIMIT string = "p_limit"      	//用户法力值上限
	UPROP_LEVEL string = "level"      	//用户神力值
	UPROP_AVATAR string = "avatar"      //用户头像
	UPROP_HISTORY_ATK string = "h_atk"  //星球被攻击的历史记录
	UPROP_BUILDING_LEVEL string = "b_level"  //星球建筑总等级
	UPROP_BELIEVER_COUNT string = "b_count"  //星球信徒数量

	UPROP_PUBLIC_ITEM string = "pitem"      	//用户发布的物品
	UPROP_PUBLIC_TIMESTAMP string = "ptstamp"   //用户发布物品的时间戳

	UPROP_SALE string = "h_sale"  //用户挂售的圣物

	UPROP_STAR string = "star"      	//用户当前星球id
	UPROP_STARTYPE string = "tstar"      	//用户当前的星球类型

	UPROP_BUFF_MANA_LIMIT string = "buffml"      	//用户当前的BUFF法力上限加成
	UPROP_BUFF_MANA_INTERVAL string = "buffmi"      	//用户当前的BUFF法力回复间隔

	UPROP_NODE string = "un"   			//用户内存所在的节点信息

	FLAG_LOADUSER string = "flu_"   	//标识，是否加载用户数据到缓存

	UPROP_TOKEN string = "token"      	//登录令牌


)

func GetUserKey(uid int64) string {
	return USER_KEY_PREFIX + util.Int64ToString(uid)
}

//设置用户会话token
func (this *cacheManager) SetUserToken(uid int64, token string) bool {
	return this.redisClient.HSet(GetUserKey(uid), UPROP_TOKEN, token)
}

//获取用户会话token
func (this *cacheManager) GetUserToken(uid int64) string {
	return this.redisClient.HGet(GetUserKey(uid), UPROP_TOKEN)
}
//
////设置用户属性
//func (this *cacheManager) SetUserAttr(uid int64, propKey string, value interface{}) bool {
//	return this.redisClient.HSet(GetUserKey(uid), propKey, value)
//}

//设置用户头像
func (this *cacheManager) SetUserAvatar(uid int64, avatar string) bool {
	return this.redisClient.HSet(GetUserKey(uid), UPROP_AVATAR, avatar)
}

//获取用户头像
func (this *cacheManager) GetUserAvatar(uid int64) string {
	return this.redisClient.HGet(GetUserKey(uid), UPROP_AVATAR)
}

//设置用户会话token
func (this *cacheManager) SetUserNickname(uid int64, nickname string) bool {
	return this.redisClient.HSet(GetUserKey(uid), UPROP_NICKNAME, nickname)
}

//获取用户昵称
func (this *cacheManager) GetUserNickname(uid int64) string {
	return this.redisClient.HGet(GetUserKey(uid), UPROP_NICKNAME)
}

//设置用户个人简介
func (this *cacheManager) SetUserDesc(uid int64, desc string) bool {
	return this.redisClient.HSet(GetUserKey(uid), UPROP_DESC, desc)
}

//获取用户个人简介
func (this *cacheManager) GetUserDesc(uid int64) string {
	return this.redisClient.HGet(GetUserKey(uid), UPROP_DESC)
}

//用户名是否存在
func (this *cacheManager) IsUsernameExist(username string) bool {
	return this.GetUidByUsername(username) != 0
}

func (this *cacheManager) SetUsernameUidMapping(username string, uid int64) bool {
	return this.redisClient.SetData(USERNAME_KEY_PREFIX + username, uid)
}

func (this *cacheManager) GetUidByUsername(username string) int64 {
	return int64(this.redisClient.GetDataInt64(USERNAME_KEY_PREFIX + username))
}

//获取用户所有信息数据
func (this *cacheManager) HSetUser(uid int64, data database.IData) {
	this.redisClient.HSetData(GetUserKey(uid), data)
}

//设置用户所有信息数据
func (this *cacheManager) HGetUser(uid int64, data database.IData) {
	this.redisClient.HGetData(GetUserKey(uid), data)
}

//用户是否存在
func (this *cacheManager) IsUserExist(uid int64) bool {
	result, _ := this.redisClient.Exists(GetUserKey(uid))
	return result
}

//用户是否在线
func (this *cacheManager) IsUserOnline(uid int64) bool {
	return this.GetUserAttrBool(uid, UPROP_ONLINE)
}

func (this *cacheManager) GetUserAttrBool(uid int64, propKey string) bool {
	return this.redisClient.HGetBool(GetUserKey(uid), propKey)
}


