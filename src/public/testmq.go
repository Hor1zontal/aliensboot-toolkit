/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/4/27
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package main

import (
	"aliens/log"
	"aliens/mq"
	"time"
)

func main() {

	service := "gate"

	config := mq.Config{Address:[]string{"127.0.0.1:9092"}, Timeout:5}
	_, err := mq.NewConsumer(mq.TYPE_KAFKA, config, service, "1", handleProxy)
	log.Debug(err)

	producer, err := mq.NewProducer(mq.TYPE_KAFKA, config)
	log.Debug(err)
	producer.SendMessage(service, "1", []byte{1,2,3})
	producer.Broadcast(service, []byte{3,2,1})

	time.Sleep(time.Hour)
}

func handleProxy(data []byte) error {
	log.Debug(data)
	return nil
}
