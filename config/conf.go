package config

import (
	"os"
)

var (
	LogicConf logicConf //逻辑服务器配置
	ConnConf  connConf  //TCP服务器配置
	WSConf    wsConf //WS服务器配置
)

// logic配置
//TODO:每个运行环境中都需要逻辑服务器的基本配置信息
type logicConf struct {
	MySQL                  string //rdb
	NSQIP                  string //mq
	RedisIP                string //cache
	//TODO: 以下地址？
	RPCIntListenAddr       string //logic服务器的RPC服务地址
	ClientRPCExtListenAddr string
	ServerRPCExtListenAddr string
	ConnRPCAddrs           string
}

// conn配置
type connConf struct {
	TCPListenAddr string //TODO: TCP监听端口，客户端访问地址
	RPCListenAddr string //TODO: RPC服务器地址，猜测用于与logic通讯
	LocalAddr     string //何用？
	LogicRPCAddrs string //逻辑服务器地址
}

// WS配置
type wsConf struct {
	WSListenAddr  string
	//以下地址是否与TCP服务器一致
	RPCListenAddr string
	LocalAddr     string
	LogicRPCAddrs string
}

//TODO: 此处会被第一个调用
func init() {
	//TODO: 通过环境变量来启动使用那种配置,实际上是使用了4个配置项目
	//TODO: 参看配置config目录下面的四个不同的配置文件
	//TODO：此方法的好处是直接通过系统参数来控制不同配置，一个包搞定所有
	//TODO：但是有些参数可能在线上需要临时调整，就不方便了，不过还是要学习下
	env := os.Getenv("gim_env")
	switch env {
	case "dev": //开发使用
		initDevConf()
	case "pre": //测试使用
		initPreConf()
	case "prod": //生产使用
		initProdConf()
	default:
		initLocalConf() //本地配置
	}
}
