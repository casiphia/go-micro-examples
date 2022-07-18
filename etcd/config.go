/*
 * @Descripttion: 828 project
 * @Version: 1.0.0
 * @Author: wangxiaodiao
 * @Email: 413586280@qq.com
 * @Date: 2022-07-18 10:31:52
 * @LastEditors: wangxiaodiao
 * @LastEditTime: 2022-07-18 11:46:51
 * @FilePath: /go-micro-examples/etcd/config.go
 */
package main

import (
	"github.com/micro/go-micro/plugins/config/source/etcd/v2"
	"github.com/micro/go-micro/v2/config"
)

type Demo struct {
	NetWork string `json:"netWork"`
	Backend string `json:"backend"`
}

func getConfig() Demo {
	getJSON("demo")
	conf := Demo{}
	config.Scan(&conf)
	return conf
}

func getJSON(pr string) error {
	//配置中心使用etcd key/value 模式
	etcdSource := etcd.NewSource(
		//设置配置中心地址
		etcd.WithAddress("127.0.0.1:2379"),
		//设置前缀，不设置默认为 /micro/config
		etcd.WithPrefix("/micro/config/"+pr),
		//是否移除前缀，这里设置为true 表示可以不带前缀直接获取对应配置,StripPrefix 从语义上看是去掉前缀的意思，如果没有去掉前缀，则会保留micro、etcd这两个key
		etcd.StripPrefix(true),
	)
	//加载配置
	return config.Load(etcdSource)
}

func watch(pr string) error {
	w, err := config.Watch("/micro/config/", "demo")
	if err != nil {
		return err
	}
	// wait for next value
	v, err := w.Next()
	if err != nil {
		return err
	}
	var demo Demo
	v.Scan(&demo)
	return nil
}
