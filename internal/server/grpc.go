package server

import (
	"github.com/raylin666/go-gin-api/pkg/grpc"
	"go-server/config"
	"go-server/grpc/system/rpc/server"
	"go-server/grpc/system/rpc/system"
	go_grpc "google.golang.org/grpc"
)

// 创建 GRPC 服务
func NewGrpcServer() {
	// 创建 gRPC 系统服务
	grpc.NewServer(grpc.Server{
		Network: config.Get().Grpc.System.Network,
		Host:    config.Get().Grpc.System.Host,
		Port:    config.Get().Grpc.System.Port,
		RegisterServer: func(g *go_grpc.Server) {
			system.RegisterSystemServer(g, &server.System{})
		},
	})
}