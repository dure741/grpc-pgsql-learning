package main

import (
	"context"
	"grpctest/hello"
	"log"
	"time"

	"google.golang.org/grpc"
)

//实现helloserver中的不完整(未定义内容)
//type HelloServer hello.HelloServer

//=== === === === === === === === === === === === === === === === === ===
const port = ":8888"

var (
	clientconn *grpc.ClientConn
	ctx        context.Context
	cancel     context.CancelFunc
	err        error
	greet      *hello.Greet
	fuck       *hello.Fuck
)

func main() {
	if clientconn, err = grpc.Dial("127.0.0.1"+port, grpc.WithInsecure(), grpc.WithBlock()); err != nil {
		log.Fatalf("conn server err :%v", err)
	}
	defer clientconn.Close()

	client := hello.NewHelloClient(clientconn)
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if greet, err = client.SyaHello(ctx, &hello.Greet{Name: "hi,this is a greeting!"}); err != nil {
		log.Fatalf("get error :%v", err)
	}
	log.Printf("get msg :%s", greet.Name)

	if fuck, err = client.FuckYou(ctx, &hello.Greet{Name: "hello???"}); err != nil {
		log.Fatalf("get error :%v", err)
	}
	log.Printf("get msg :%s", fuck.Curse)
}
