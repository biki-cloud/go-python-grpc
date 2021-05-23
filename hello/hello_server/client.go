package main

import (
	"context"
	"fmt"
	"go-python-grpc/hello/hellopb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("hello world")

	// create client connection
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connetct: %v", err)
	}

	defer func(cc *grpc.ClientConn) {
		err := cc.Close()
		if err != nil {

		}
	}(cc)

	// create client
	c := hellopb.NewHelloServiceClient(cc)

	doRequest(c)
}

func doRequest(c hellopb.HelloServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	fmt.Println("I'm Stephan mark")
	// create request
	req := &hellopb.MsgRequest{
		Message: "bikibikibiki",
	}
	// Greetファンクションはサーバーにリクエストを送ってレスポンスを返す。
	res, err := c.Hello(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Hello RPC: %v", err)
	}

	log.Printf("Responce from Greet: %v \n", res.Result)
}
