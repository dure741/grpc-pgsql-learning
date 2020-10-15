package main

import (
	"context"
	"fmt"
	"grpctest/hello"
	"log"
	"net"

	"google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	status "google.golang.org/grpc/status"
)

//实现helloserver中的不完整(未定义内容)
//type HelloServer hello.HelloServer
type HelloServer struct {
	// Greet hello.Greet
	// Fuck hello.Fuck
}

// type Greet hello.Greet
// type Fuck hello.Fuck

func (h *HelloServer) SyaHello(ctx context.Context, greet *hello.Greet) (*hello.Greet, error) {
	return greet, status.Errorf(codes.OK, "ok")
}

func (h *HelloServer) FuckYou(ctx context.Context, greet *hello.Greet) (*hello.Fuck, error) {
	fuck := &hello.Fuck{Curse: "Fuck U!!!!!!!!!!!"}
	//return nil, status.Errorf(codes.Unimplemented, "Unimplemented!")
	return fuck, status.Errorf(codes.OK, "ok")
}

//=== === === === === === === === === === === === === === === === === ===
const port = ":8888"

func main() {
	//ctx := context.Background()
	log.Println("Start")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hello.RegisterHelloServer(s, &HelloServer{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		fmt.Println("error!!")
	}
	fmt.Println("listen finished!")
}
