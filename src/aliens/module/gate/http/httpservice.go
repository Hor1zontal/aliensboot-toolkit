/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/4/2
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package http

import (
	"net/http"
	"aliens/log"
	"aliens/module/gate/conf"
)

func Init() {
	if conf.Config.Http.Address != "" {
		go func() {
			//http.HandleFunc("/", httpHandle)
			log.Info(http.ListenAndServe(conf.Config.Http.Address, nil))
		}()
	}
}


func Close() {

}

//添加弹幕信息
//func httpHandle(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()
//	body, _ := ioutil.ReadAll(r.Body)
//	response, error := route.HandleUrlMessage(r.RequestURI, body)
//	if error != nil {
//		response = []byte(error.Error())
//	}
//	_, err := w.Write(response)
//	if err != nil {
//		log.Debug(err.Error())
//	}
//}




