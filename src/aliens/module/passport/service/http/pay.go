package http

import (
	"net/http"
	"gok/passportserver/helper"
	"gok/log"
	"gok/passportserver/conf"
	"gok/rpc"
	"aliens/common/character"
	"gok/passportserver/cache"
	"gok/network/msg/protocol"
	"github.com/gogo/protobuf/proto"
	"sync"
	"gok/passportserver/db"
	"time"
)

var PAY_LOCK = &sync.Mutex{}

func Pay(responseWriter http.ResponseWriter, request *http.Request) {
	payProxy(request)
	sendToClient(responseWriter, "success")
}

func payProxy(request *http.Request) {
	request.ParseForm()

	memo := request.FormValue(PARAM_MEMO)
	appID := request.FormValue(PARAM_APPID)
	amount := request.FormValue(PARAM_AMOUNT)
	orderId := request.FormValue(PARAM_ORDER)
	sign := request.FormValue(PARAM_SIGN)
	accountID := request.FormValue(PARAM_ACCOUNTID)

	signText := "accountId" + accountID + "&amount=" + amount + "&appId=" + appID + "&cpOrderId=" + orderId + conf.Server.AppKey
	signResult := helper.MD5Hash(signText)
	log.Debug("signResult %v : sign : %v", signResult, sign)
	//TODO 验证失败返回

	shopID := character.StringToInt32(memo)
	if (shopID == 0) {
		log.Debug("shopID not found %v", shopID)
		return
	}
	uid := character.StringToInt32(accountID)
	if (uid == 0 || !cache.UserCache.IsUserExist(uid)) {
		log.Debug("user not found %v", uid)
		return
	}

	PAY_LOCK.Lock()
	defer PAY_LOCK.Unlock()
	order := &db.DBOrder{ID: orderId}
	if db.DatabaseHandler.IDExist(order) {
		log.Debug("%v: order %v repeat", uid, orderId)
		return
	}

	rpc.UserServiceProxy.Call(uid, cache.UserCache.GetUserNode(uid), &protocol.C2GS{
		Sequence: []int32{524},
		AddShopItem: proto.Int32(shopID),
	})

	order.UserID = uid
	order.ProductID = shopID
	order.Amount = character.StringToFloat64(amount)
	order.CreateTime = time.Now()
	db.DatabaseHandler.Insert(order)
}


