package cmd

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"net"
	rpcSrv "zk-leader-election/src/zk_leader_election/idl/dml_service"
	implSrv "zk-leader-election/src/zk_leader_election/server/dml_service"
)

var rootCmd = &cobra.Command{
	Use:   "zk-leader-election",
	Short: "a demo of the leader election via zk",
	Long:  "a demo of the leader election via zk",
	Run:   rootFunc,
}

func rootFunc(cmd *cobra.Command, args []string) {
	listen, err := net.Listen("tcp", ":10086")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	rpcSrv.RegisterDMLServiceServer(grpcServer, implSrv.NewReplaceableDMLService())
	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
