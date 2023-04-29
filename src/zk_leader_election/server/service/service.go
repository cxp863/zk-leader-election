package service

import "context"
import idlsrv "zk-leader-election/src/zk_leader_election/idl/service"

type GreetServiceServer struct {
}

func (s *GreetServiceServer) Greeting(_ context.Context, req *idlsrv.GreetRequest) (*idlsrv.GreetResponse, error) {
	if req.Message == "" {
		return &idlsrv.GreetResponse{
			Code:    idlsrv.ResponseCode_kError,
			Message: "no data",
		}, nil
	}
	return &idlsrv.GreetResponse{
		Code:    idlsrv.ResponseCode_kOk,
		Message: "echo: " + req.Message,
	}, nil
}
