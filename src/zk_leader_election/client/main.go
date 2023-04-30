package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"sync"
	"sync/atomic"
	"time"
	idlsrv "zk-leader-election/src/zk_leader_election/idl/service"
)

func main() {
	conn, err := grpc.Dial("192.168.2.205:10086", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	client := idlsrv.NewGreetServiceClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), time.Second*1)
	value := send(ctx, client)
	fmt.Println(value)

}

func send(ctx context.Context, c idlsrv.GreetServiceClient) int64 {
	wg := sync.WaitGroup{}
	wg.Add(1)
	var value = atomic.Int64{}
	go func(wg *sync.WaitGroup) {
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				return
			default:
				_, err := c.Greeting(ctx, &idlsrv.GreetRequest{Message: "hello"})
				if err == nil {
					value.Add(1)
				}
			}
		}
	}(&wg)
	wg.Wait()
	return value.Load()
}
