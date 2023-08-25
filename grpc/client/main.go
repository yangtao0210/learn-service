package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "learn-service/grpc/server/proto"
	"log"
)

// ClientToken 自定义TOKEN认证:
// 1.实现credentials接口的方法
// 2.连接服务端时需要携带token信息
// 3.服务端会根据连接的元数据检测token是否有效
type ClientToken struct {
}

func (c ClientToken) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appId":  "jacktao",
		"appKey": "guaika@12",
	}, nil
}
func (c ClientToken) RequireTransportSecurity() bool {
	return false
}
func main() {
	// 1）客户端TLS认证：自签名文件
	//creds, _ := credentials.NewClientTLSFromFile("/Users/jacktaoyang/Desktop/personal/learn-service/grpc/key/test.pem", "*.jack")
	//conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(creds))
	// 2) TOEKN验证
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientToken)))
	// 链接到server端：此处禁用安全传输，没有加密和验证
	conn, err := grpc.Dial("127.0.0.1:9090", opts...)
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	defer conn.Close()

	// 创建链接客户端
	client := pb.NewSayHelloClient(conn)
	// 执行rpc调用(方法会在服务端实现并返回结果)
	res, err := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "jack"})
	if err != nil {
		return
	}
	log.Println(res.ResponseMsg)
}
