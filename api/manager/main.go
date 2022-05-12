package main

import (
	"api/manager/baseInfo"
	"api/manager/cmd"
	"api/manager/events"
	"api/manager/monitor"
	"fmt"
	"log"
	"net"
)

const (
	Address = ":50052"
	Network = "tcp"
)

//var ServiceName = "manager.service"

func main() {
	// 监听端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		fmt.Println(err)
	}
	host, sport, err := net.SplitHostPort(listener.Addr().String())
	if err != nil {
		log.Println(err)
	}

	// 初始话本机基本信息
	baseInfo.InfoInit()

	// 用来注册和启动服务
	go func() {

	}()

	// 事件监听
	events.Init()
	// 服务轮询协程
	monitor.PollStart()
}
