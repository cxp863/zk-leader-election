package dml_service

import (
	"context"
	"go.uber.org/zap"
	"zk-leader-election/src/zk_leader_election/idl/dml_service"
)

func newDMLServiceLog(service DMLService) DMLService {
	return &DMLServiceLog{
		inner:  service,
		logger: zap.NewNop(),
	}
}

type DMLServiceLog struct {
	inner  DMLService
	logger *zap.Logger
}

func (s *DMLServiceLog) Create(ctx context.Context, req *dml_service.CreateRequest) (*dml_service.CreateResponse, error) {
	resp, err := s.inner.Create(ctx, req)
	if err != nil || resp.Status == nil || resp.Status.Code != 0 {
		s.logger.Warn("Create failed")
	}
	return resp, err
}

func (s *DMLServiceLog) Read(ctx context.Context, req *dml_service.ReadRequest) (*dml_service.ReadResponse, error) {
	resp, err := s.inner.Read(ctx, req)
	if err != nil || resp.Status == nil || resp.Status.Code != 0 {
		s.logger.Warn("Read failed")
	}
	return resp, err
}

func (s *DMLServiceLog) Update(ctx context.Context, req *dml_service.UpdateRequest) (*dml_service.UpdateResponse, error) {
	resp, err := s.inner.Update(ctx, req)
	if err != nil || resp.Status == nil || resp.Status.Code != 0 {
		s.logger.Warn("Update failed")
	}
	return resp, err
}

func (s *DMLServiceLog) Delete(ctx context.Context, req *dml_service.DeleteRequest) (*dml_service.DeleteResponse, error) {
	resp, err := s.inner.Delete(ctx, req)
	if err != nil || resp.Status == nil || resp.Status.Code != 0 {
		s.logger.Warn("Delete failed")
	}
	return resp, err
}
