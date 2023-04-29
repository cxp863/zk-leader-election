package main

import (
	"google.golang.org/grpc"
	"net"
	idlsrv "zk-leader-election/src/zk_leader_election/idl/service"
	"zk-leader-election/src/zk_leader_election/server/service"
)

func main() {
	listen, err := net.Listen("tcp", ":10086")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	idlsrv.RegisterGreetServiceServer(grpcServer, &service.GreetServiceServer{})
	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}
