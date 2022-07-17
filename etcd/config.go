package main

import (
	"github.com/micro/go-micro/plugins/config/source/etcd/v2"
	"github.com/micro/go-micro/v2/config"
)

type Config struct {
	NetWork string `json:"NetWork"`
	Backend string `json:"Backend"`
}

func getConfig() Config {
	getJSON()
	conf := Config{}
	config.Scan(&conf)
	return conf
}

func getJSON() error {
	//配置中心使用etcd key/value 模式
	etcdSource := etcd.NewSource(
		//设置配置中心地址
		etcd.WithAddress("127.0.0.1:2379"),
		//设置前缀，不设置默认为 /micro/config
		//etcd.WithPrefix(),
		//是否移除前缀，这里设置为true 表示可以不带前缀直接获取对应配置
		etcd.StripPrefix(true),
	)
	//加载配置
	return config.Load(etcdSource)
}
