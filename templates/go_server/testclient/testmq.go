/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/4/27
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package testclient

import (
	"github.com/KylinHe/aliensboot/log"
	"github.com/KylinHe/aliensboot/mq"
	"time"
)

func main() {

	service := "gate"

	config := mq.Config{Type: mq.KAFKA, Address: []string{"127.0.0.1:9092"}, Timeout: 5}
	_, err := mq.NewConsumer(config, service, "1", handleProxy)
	log.Debug(err)

	producer, err := mq.NewProducer(config)
	log.Debug(err)

	time.Sleep(10 * time.Second)
	producer.SendMessage(service, "1", []byte{1, 2, 3, 4})
	producer.Broadcast(service, []byte{3, 2, 1})

	time.Sleep(time.Hour)
}

func handleProxy(data []byte) error {
	log.Debugf("accept data %v", data)
	return nil
}
