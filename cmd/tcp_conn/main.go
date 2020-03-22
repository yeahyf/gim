package main

import (
	"gim/api/tcp_conn"
	"gim/config"
	tcp_conn2 "gim/internal/tcp_conn"
	"gim/pkg/rpc_cli"
	"gim/pkg/util"
)

func main() {
	//TODO: 在一个goroutine中异步启动rpc服务
	go func() {
		defer util.RecoverPanic()  //Recover的统一处理
		tcp_conn.StartRPCServer()  //启动RPC服务端
	}()

	// 初始化Rpc Client，作为访问logic服务器的客户端，
	// 直接看逻辑服务器的地址 config.ConnConf.LogicRPCAddrs
	rpc_cli.InitLogicIntClient(config.ConnConf.LogicRPCAddrs)

	// 启动长链接服务器，这个是对客户端而言的
	server := tcp_conn2.NewTCPServer(config.ConnConf.TCPListenAddr, 10)
	server.Start()
}
