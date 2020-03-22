package tcp_conn

import (
	"gim/pkg/logger"
	"gim/pkg/util"
	"net"
)

// TCPServer TCP服务器 结构体，简单封装
type TCPServer struct {
	Address            string // 端口
	AcceptGoroutineNum int    // 接收建立连接的goroutine数量
}

// NewTCPServer 创建TCP服务器
func NewTCPServer(address string, acceptGoroutineNum int) *TCPServer {
	return &TCPServer{
		Address:            address,
		AcceptGoroutineNum: acceptGoroutineNum,
	}
}

// Start 启动服务器
func (t *TCPServer) Start() {
	addr, err := net.ResolveTCPAddr("tcp", t.Address)
	if err != nil {
		logger.Sugar.Error(err)
		panic(err)
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		logger.Sugar.Error("error listening:", err.Error())
		return
	}
	for i := 0; i < t.AcceptGoroutineNum; i++ {
		go t.Accept(listener)
	}
	logger.Sugar.Info("tcp server start")
	select {}
}

// Accept 接收客户端的TCP长连接的建立
func (t *TCPServer) Accept(listener *net.TCPListener) {
	defer util.RecoverPanic()

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			logger.Sugar.Error(err)
			continue
		}

		//设置为长链接
		err = conn.SetKeepAlive(true)
		if err != nil {
			logger.Sugar.Error(err)
		}
		//此处用于统一处理封包和拆包问题
		connContext := NewConnContext(conn)
		//异步处理这个连接的请求
		go connContext.DoConn()
	}
}
