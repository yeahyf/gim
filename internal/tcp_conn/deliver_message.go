package tcp_conn

import (
	"context"
	"gim/pkg/grpclib"
	"gim/pkg/logger"
	"gim/pkg/pb"
)

// 将消息发送给需要接收的用户，这里需要获取到他的连接，连接放在一个map里边，
// 通过DeviceId来进行索引
func DeliverMessage(ctx context.Context, req *pb.DeliverMessageReq) error {
	// 获取设备对应的TCP连接
	conn := load(req.DeviceId)
	if conn == nil {
		logger.Sugar.Warn("ctx id nil")
		return nil
	}

	// 发送消息
	conn.Output(pb.PackageType_PT_MESSAGE, grpclib.GetCtxRequstId(ctx), nil, req.Message)
	return nil
}
