package tcp_conn

import (
	"context"
	"gim/config"
	"gim/internal/tcp_conn"
	"gim/pkg/logger"
	"gim/pkg/pb"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// TODO: 实现消息投递接口
type ConnIntServer struct{}

// Message 投递消息
func (s *ConnIntServer) DeliverMessage(ctx context.Context, req *pb.DeliverMessageReq) (*pb.DeliverMessageResp, error) {
	return &pb.DeliverMessageResp{}, tcp_conn.DeliverMessage(ctx, req)
}

// UnaryServerInterceptor 服务器端的单向调用的拦截器
//TODO: 用于拦截logic server发送过来的消息？？
func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	logger.Logger.Debug("interceptor", zap.Any("info", info), zap.Any("req", req), zap.Any("resp", resp))
	return resp, err
}

// StartRPCServer 启动rpc服务器
//TODO: 此处了解下gRPC这个rpc框架
func StartRPCServer() {
	listener, err := net.Listen("tcp", config.ConnConf.RPCListenAddr)
	//异常不处理，直接panic
	if err != nil {
		panic(err)
	}
	//TODO: grpc又是一个新玩意，需要深入了解下grpc的相关机制
	server := grpc.NewServer(grpc.UnaryInterceptor(UnaryServerInterceptor)) //这个地方后续再了解

	//注册这个服务
	pb.RegisterConnIntServer(server, &ConnIntServer{})
	logger.Logger.Debug("rpc服务已经开启")
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
