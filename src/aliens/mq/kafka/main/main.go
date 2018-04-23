/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/4/21
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package main

import (
	"aliens/log"
	"aliens/mq/kafka"
)

func main() {
	log.Init("conf/aliens/log.xml")
	//kafka.Consumer("group1")
	kafka.AsyncProducer()

	kafka.AsyncProducer()
}
