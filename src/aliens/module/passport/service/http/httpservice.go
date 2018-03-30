/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2017/11/15
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package http

import (
	"net/http"
	"gok/passportserver/conf"
	"gok/log"
)

func Init() {

	go func() {
		//http.Handle("/api/", http.FileServer(http.Dir("./../doc/http/")))

		http.HandleFunc("/api/user/loginoauth", ChannelLogin)
		http.HandleFunc("/api/user/pay", Pay)

		log.Info("%v", http.ListenAndServe(conf.Server.HTTPAddress, nil))
	}()

}

func Close() {

}

func sendToClient(responseWriter http.ResponseWriter, content string) {
	_, err := responseWriter.Write([]byte(content))
	if err != nil {
		log.Debug(err.Error())
	}
}

