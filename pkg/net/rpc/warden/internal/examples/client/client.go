package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/ryanli-me/kratos/pkg/log"
	"github.com/ryanli-me/kratos/pkg/net/rpc/warden"
	pb "github.com/ryanli-me/kratos/pkg/net/rpc/warden/internal/proto/testproto"
)

// usage: ./client -grpc.target=test.service=127.0.0.1:8080
func main() {
	log.Init(&log.Config{Stdout: true})
	flag.Parse()
	conn, err := warden.NewClient(nil).Dial(context.Background(), "127.0.0.1:8081")
	if err != nil {
		panic(err)
	}
	cli := pb.NewGreeterClient(conn)
	normalCall(cli)
}

func normalCall(cli pb.GreeterClient) {
	reply, err := cli.SayHello(context.Background(), &pb.HelloRequest{Name: "tom", Age: 23})
	if err != nil {
		panic(err)
	}
	fmt.Println("get reply:", *reply)
}
