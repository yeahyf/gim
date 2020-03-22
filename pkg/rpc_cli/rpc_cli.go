package rpc_cli

import (
	"context"
	"fmt"
	"gim/pkg/grpclib"
	"gim/pkg/logger"
	"gim/pkg/pb"

	"google.golang.org/grpc"
)

var (
	LogicIntClient   pb.LogicIntClient  //逻辑服务器的客户端
	ConnectIntClient pb.ConnIntClient   //TCP服务器的客户端
)

func InitLogicIntClient(addr string) {
	conn, err := grpc.DialContext(context.TODO(), addr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor))
	if err != nil {
		logger.Sugar.Error(err)
		panic(err)
	}

	LogicIntClient = pb.NewLogicIntClient(conn)
}

func InitConnIntClient(addr string) {
	conn, err := grpc.DialContext(context.TODO(), addr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, grpclib.Name)))
	if err != nil {
		logger.Sugar.Error(err)
		panic(err)
	}

	ConnectIntClient = pb.NewConnIntClient(conn)
}
