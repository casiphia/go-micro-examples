package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/micro/go-micro/plugins/broker/rabbitmq/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/server"
)

// 定义一个发布消息的函数：每隔1秒发布一条消息
func loopPublish(event micro.Event) {
	for {
		time.Sleep(time.Duration(1) * time.Second)

		curUnix := strconv.FormatInt(time.Now().Unix(), 10)
		msg := "{\"Id\":" + curUnix + ",\"Name\":\"张三\"}"
		event.Publish(context.TODO(), msg)
	}
}

// 定义一个接收消息的函数：将收到的消息打印出来
func handle(ctx context.Context, msg interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint(r))
			log.Println(err)
		}
	}()

	log.Println(msg)
	return
}

func main() {
	rabbitmqUrl := "amqp://guest:guest@127.0.0.1:5672/"
	exchangeName := "amq.topic"
	subcribeTopic := "test"
	queueName := "rabbitmqdemo_test"

	// 默认是application/protobuf，这里演示用的是Json，所以要改下
	server.DefaultContentType = "application/json"

	// 创建 RabbitMQ Broker
	b := rabbitmq.NewBroker(
		broker.Addrs(rabbitmqUrl),           // RabbitMQ访问地址，含VHost
		rabbitmq.ExchangeName(exchangeName), // 交换机的名称
		rabbitmq.DurableExchange(),          // 消息在Exchange中时会进行持久化处理
		rabbitmq.PrefetchCount(1),           // 同时消费的最大消息数量
	)

	// 创建Service，内部会初始化一些东西，必须在NewSubscribeOptions前边
	service := micro.NewService(
		micro.Broker(b),
	)
	service.Init()

	// 初始化订阅上下文：这里不是必需的，订阅会有默认值
	subOpts := broker.NewSubscribeOptions(
		rabbitmq.DurableQueue(),   // 队列持久化，消费者断开连接后，消息仍然保存到队列中
		rabbitmq.RequeueOnError(), // 消息处理函数返回error时，消息再次入队列
		rabbitmq.AckOnSuccess(),   // 消息处理函数没有error返回时，go-micro发送Ack给RabbitMQ
	)

	// 注册订阅
	micro.RegisterSubscriber(
		subcribeTopic,    // 订阅的Topic
		service.Server(), // 注册到的rpcServer
		handle,           // 消息处理函数
		server.SubscriberContext(subOpts.Context), // RabbitMQ订阅参数，也可以不设置
		server.SubscriberQueue(queueName),         // 队列名称
	)

	// 发布事件消息
	event := micro.NewEvent(subcribeTopic, service.Client())
	go loopPublish(event)

	log.Println("Service is running ...")
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
