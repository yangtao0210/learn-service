package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	pb "learn-service/grpc/server/proto"
	"net"
)

type SayHelloServer struct {
	pb.UnimplementedSayHelloServer
}

func (s *SayHelloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("未传输token")
	}
	var appID, appKey string
	if v, ok := meta["appid"]; ok {
		appID = v[0]
	}
	if v, ok := meta["appkey"]; ok {
		appKey = v[0]
	}
	if appID != "jacktao" || appKey != "guaika@12" {
		return nil, errors.New("token 不正确")
	}
	fmt.Println(req.RequestName, req.Age)
	return &pb.HelloResponse{ResponseMsg: "hello" + req.RequestName}, nil
}

func main() {
	// 1）服务端TLS认证：自签名文件和私钥文件
	// creds, _ := credentials.NewServerTLSFromFile("/Users/jacktaoyang/Desktop/personal/learn-service/grpc/key/test.pem", "/Users/jacktaoyang/Desktop/personal/learn-service/grpc/key/test.key")
	// grpcServer := grpc.NewServer(grpc.Creds(creds))
	// 1.开启端口监听
	listen, _ := net.Listen("tcp", ":9090")
	// 2.创建grpc服务时,将TSL
	grpcServer := grpc.NewServer()
	// 3.注册自定义服务到grcp中
	pb.RegisterSayHelloServer(grpcServer, &SayHelloServer{})

	// 4.启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Println("服务启动失败", err)
		return
	}
}
