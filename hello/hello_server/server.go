package main

import (
	"context"
	"fmt"
	"go-python-grpc/hello/hellopb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (s *server) Hello(ctx context.Context, req *hellopb.MsgRequest) (*hellopb.MsgResponse, error) {
	fmt.Printf("Hello function was invoked %v \n", req)

	msg := req.GetMessage()
	result := "hello " + msg

	res := &hellopb.MsgResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Printf("hello world \n")

	const port = ":50051"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hellopb.RegisterHelloServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
