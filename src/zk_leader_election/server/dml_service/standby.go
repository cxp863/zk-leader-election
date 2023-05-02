package dml_service

import (
	"context"
	"zk-leader-election/src/zk_leader_election/idl/dml_service"
)

func newDMLServiceStandBy() DMLService {
	return &DMLServiceStandBy{}
}

type DMLServiceStandBy struct {
}

func (s *DMLServiceStandBy) Create(ctx context.Context, req *dml_service.CreateRequest) (*dml_service.CreateResponse, error) {
	return &dml_service.CreateResponse{Status: &dml_service.ResponseStatus{
		Code:    dml_service.ResponseCode_kNotLeaderError,
		Message: "not leader",
	}}, nil
}

func (s *DMLServiceStandBy) Read(ctx context.Context, req *dml_service.ReadRequest) (*dml_service.ReadResponse, error) {
	return &dml_service.ReadResponse{Status: &dml_service.ResponseStatus{
		Code:    dml_service.ResponseCode_kNotLeaderError,
		Message: "not leader",
	}}, nil
}

func (s *DMLServiceStandBy) Update(ctx context.Context, req *dml_service.UpdateRequest) (*dml_service.UpdateResponse, error) {
	return &dml_service.UpdateResponse{Status: &dml_service.ResponseStatus{
		Code:    dml_service.ResponseCode_kNotLeaderError,
		Message: "not leader",
	}}, nil
}

func (s *DMLServiceStandBy) Delete(ctx context.Context, req *dml_service.DeleteRequest) (*dml_service.DeleteResponse, error) {
	return &dml_service.DeleteResponse{Status: &dml_service.ResponseStatus{
		Code:    dml_service.ResponseCode_kNotLeaderError,
		Message: "not leader",
	}}, nil
}
