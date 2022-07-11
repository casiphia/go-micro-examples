package main

import (
	"log"

	"github.com/micro/go-micro/plugins/broker/rabbitmq/v2"
	"github.com/micro/go-micro/v2/broker"
)

func eventHandle(event broker.Event) (err error) {
	topic := event.Topic()
	log.Println(topic)
	body := event.Message().Body
	header := make(map[string]string)

	//为了兼容其他框架发送的mq消息
	msgTopic := header["Micro-Topic"]
	if msgTopic == "" {
		header["Micro-Topic"] = topic
	}
	log.Println(body)
	return
}

func main() {
	rabbitmqUrl := "amqp://guest:guest@127.0.0.1:5672/"
	exchangeName := "amq.topic"
	subcribeTopic := "test"
	queueName := "rabbitmqdemo_test"

	// 创建 RabbitMQ Broker
	b := rabbitmq.NewBroker(
		broker.Addrs(rabbitmqUrl),           // RabbitMQ访问地址，含VHost
		rabbitmq.ExchangeName(exchangeName), // 交换机的名称
		rabbitmq.DurableExchange(),          // 消息在Exchange中时会进行持久化处理
		rabbitmq.PrefetchCount(10),          // 同时消费的最大消息数量
	)

	err := b.Connect()
	if err != nil {
		panic(err)
	}

	b.Subscribe(subcribeTopic, eventHandle,
		rabbitmq.DurableQueue(),   // 队列持久化，消费者断开连接后，消息仍然保存到队列中
		broker.Queue(queueName),   // 队列名称
		rabbitmq.RequeueOnError(), // 消息处理函数返回error时，消息再次入队列
		rabbitmq.AckOnSuccess(),   // 消息处理函数没有error返回时，go-micro发送Ack给RabbitMQ
	)

	log.Println("Consumer is running ...")

	select {}
}
