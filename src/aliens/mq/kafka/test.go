/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/4/8
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package kafka


// syncProducer 同步生产者
// 并发量小时，可以用这种方式
//func syncProducer() {
//	config := sarama.NewConfig()
//	//  config.Producer.RequiredAcks = sarama.WaitForAll
//	//  config.Producer.Partitioner = sarama.NewRandomPartitioner
//	config.Producer.Return.Successes = true
//	config.Producer.Timeout = 5 * time.Second
//	p, err := sarama.NewSyncProducer(strings.Split("localhost:9092", ","), config)
//	defer p.Close()
//	if err != nil {
//		log.Error(err)
//		return
//	}
//
//	v := "sync: " + strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10000))
//	fmt.Fprintln(os.Stdout, v)
//	msg := &sarama.ProducerMessage{
//		Topic: topics,
//		Value: sarama.ByteEncoder(v),
//	}
//	if _, _, err := p.SendMessage(msg); err != nil {
//		log.Error(err)
//		return
//	}
//}



