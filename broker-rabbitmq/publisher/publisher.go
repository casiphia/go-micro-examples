package main

import (
	"log"
	"strconv"
	"time"

	"github.com/micro/go-micro/plugins/broker/rabbitmq/v2"
	"github.com/micro/go-micro/v2/broker"
)

func main() {
	rabbitmqUrl := "amqp://guest:guest@127.0.0.1:5672/"
	exchangeName := "amq.topic"
	routeTopic := "test"

	// 创建 RabbitMQ Broker
	b := rabbitmq.NewBroker(
		broker.Addrs(rabbitmqUrl),           // RabbitMQ访问地址，含VHost
		rabbitmq.ExchangeName(exchangeName), // 交换机的名称
		rabbitmq.DurableExchange(),          // 消息在Exchange中时会进行持久化处理
	)

	// 先连接，才能用
	err := b.Connect()
	if err != nil {
		log.Println(err)
	}

	//循环发送消息
	loopPublish(b, routeTopic)

	log.Println("Publisher is working ...")

	select {}
}

func loopPublish(b broker.Broker, topic string) {
	//间隔1s发送消息
	tick := time.NewTicker(time.Second)
	for range tick.C {
		publish(b, topic)
	}
}

//消息发送
func publish(b broker.Broker, topic string) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("error:", err)
		}
	}()

	log.Println("开始发送消息.")

	curUnix := strconv.FormatInt(time.Now().Unix(), 10)
	omsg := "{\"Id\":" + curUnix + ",\"Name\":\"张三\"}"
	msg := &broker.Message{
		Body: []byte(omsg),
	}

	log.Println("消息类容：", omsg)
	err := b.Publish(topic, msg)
	if err != nil {
		log.Println(err)
	}
}
