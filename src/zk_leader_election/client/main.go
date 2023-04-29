package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	idlsrv "zk-leader-election/src/zk_leader_election/idl/service"
)

func main() {

	conn, err := grpc.Dial("localhost:10086", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	c := idlsrv.NewGreetServiceClient(conn)

	resp, err := c.Greeting(context.Background(), &idlsrv.GreetRequest{Message: "hello"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Message)
}
