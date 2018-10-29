/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/8/15
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package cache

func Test() {
	//
	//Users      []*User `protobuf:"bytes,13,rep,name=users" json:"users,omitempty"`
	//User       *User   `protobuf:"bytes,14,opt,name=user" json:"user,omitempty"`
	//Data       []byte  `protobuf:"bytes,15,opt,name=data,proto3" json:"data,omitempty"`
	//Intdata    []int32 `protobuf:"varint,16,rep,packed,name=intdata" json:"intdata,omitempty"`

	//user1 := protocol.Test{
	//	Id:5,
	//	Username:"2",
	//}
	//
	//user := &protocol.User{
	//	Id:999,
	//	Username:"2",
	//	Password:"3",
	//	Salt:"4",
	//	Channeluid:"5",
	//	Channel:"6",
	//	Avatar:"7",
	//	Mobile:"8",
	//	Openid:"9",
	//	Ip:"10",
	//	Status:11,
	//	RegTime:12,
	//	Test:user1,
	//	Test1:[]protocol.Test{user1},
	//}
	//
	//startTime := time.Now()
	//for i:=0; i <100000; i ++ {
	//	PassportCache.HSetUser(999, user)
	//}
	//
	//log.Debug(time.Now().Sub(startTime).Seconds())
	//
	//
	//startTime = time.Now()
	//user2 := &protocol.User{}
	//PassportCache.HGetUser(999, user2)
	//log.Debug(time.Now().Sub(startTime).Seconds())
	//log.Debug(user2)
}
