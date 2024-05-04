package main

import (
	"fmt"
	"github.com/LXJ0000/go-rpc/app/user/bootstrap"
	"github.com/LXJ0000/go-rpc/app/user/internal/repository"
	"github.com/LXJ0000/go-rpc/app/user/internal/service"
	pb "github.com/LXJ0000/go-rpc/idl/pb/user"
	"github.com/LXJ0000/go-rpc/utils/discovery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	// etcd 地址
	etcdAddress := []string{env.Etcd.Address}

	// 服务注册
	etcdRegister := discovery.NewRegister(etcdAddress, logrus.New())
	defer etcdRegister.Stop()

	server := grpc.NewServer()
	defer server.Stop()

	// 绑定service
	userRepository := repository.NewUserRepository(app.Orm)
	userService := service.NewUserServiceServer(userRepository)
	pb.RegisterUserServiceServer(server, userService)

	grpcAddress := env.Services["user"].Addr
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}

	userNode := discovery.Server{
		Name: env.Domain["user"].Name,
		Addr: grpcAddress,
	}
	if _, err := etcdRegister.Register(userNode, 10); err != nil {
		panic(fmt.Sprintf("start server failed, err: %v", err))
	}

	logrus.Info("server started listen on ", grpcAddress)
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
