/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/4/3
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package cache

import "time"

type DBUser struct {
	ID       int32  `bson:"_id" gorm:"AUTO_INCREMENT"`             //用户id
	Username string `bson:"username"  rorm:"uname"  unique:"true"` //用户名 渠道信息_渠道用户id存Username
	Password string `bson:"password"  rorm:"pwd"`                  //加密的密码
	Salt     string `bson:"salt"      rorm:"salt"`                 //加密的salt

	ChannelUID string `bson:"cuid"    rorm:"cuid"`    //用户的渠道的渠道用户id
	Channel    string `bson:"channel" rorm:"channel"` //用户的渠道信息 渠道用户id存Username
	Avatar	   string `bson:"avatar"  rorm:"avatar"` //用户的头像地址

	Mobile string `bson:"mobile"` //用户电话
	OpenID string `bson:"openid"` //微信OPENID 绑定微信填写
	IP     string `bson:"ip"`     //最后一次登录的ip
	Status  byte      `bson:"status"`  //用户状态 0正常  1封号
	RegTime time.Time `bson:"regtime"` //用户注册时间
}

type DBPassport struct {

}




func getUserID() int32 {
	return UserCache.HGetInt32("user", "abc")
}