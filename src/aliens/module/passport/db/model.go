package db

import (
	"time"
)

const (
	//collection constants
	C_USER = "user"
	//increment id constants
	//ID_USER = "uid"
)

//角色
type DBUser struct {
	ID       int64  `bson:"_id" gorm:"AUTO_INCREMENT"`             //用户id
	Username string `bson:"username"  rorm:"uname"  unique:"true"` //用户名 渠道信息_渠道用户id存Username
	Password string `bson:"password"  rorm:"pwd"`                  //加密的密码
	Salt     string `bson:"salt"      rorm:"salt"`                 //加密的salt

	ChannelUID string `bson:"cuid"    rorm:"cuid"`    //用户的渠道的渠道用户id
	Channel    string `bson:"channel" rorm:"channel"` //用户的渠道信息 渠道用户id存Username
	Avatar	   string `bson:"avatar"  rorm:"avatar"`  //用户的头像地址

	Mobile string `bson:"mobile"` //用户电话
	OpenID string `bson:"openid"` //微信OPENID 绑定微信填写
	IP     string `bson:"ip"`     //最后一次登录的ip
	Status  byte      `bson:"status"`  //用户状态 0正常  1封号
	RegTime time.Time `bson:"regtime"` //用户注册时间
}

func (this *DBUser) Name() string {
	return C_USER
}

func (this *DBUser) GetID() interface{} {
	return this.ID
}


//type DBOrder struct {
//	ID         string    `bson:"_id"`        //订单id
//	UserID     int32     `bson:"userid"`     //用户id
//	ProductID  int32     `bson:"productid"`  //充值商品id
//	Amount     float64   `bson:"amount"`     //充值金额
//	CreateTime time.Time `bson:"createTime"` //订单创建时间
//}
//
//func (this *DBOrder) Name() string {
//	return "order"
//}
//
//func (this *DBOrder) GetID() interface{} {
//	return this.ID
//}
