package dml_service

import (
	"context"
	"time"
	"zk-leader-election/src/zk_leader_election/idl/dml_service"
)

func newDMLServiceContext(service DMLService, timeout time.Duration) DMLService {
	return &DMLServiceContext{
		inner:   service,
		timeout: timeout,
	}
}

type DMLServiceContext struct {
	inner   DMLService
	timeout time.Duration
}

func (s *DMLServiceContext) Create(ctx context.Context, req *dml_service.CreateRequest) (*dml_service.CreateResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	return s.inner.Create(ctx, req)
}

func (s *DMLServiceContext) Read(ctx context.Context, req *dml_service.ReadRequest) (*dml_service.ReadResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	return s.inner.Read(ctx, req)
}

func (s *DMLServiceContext) Update(ctx context.Context, req *dml_service.UpdateRequest) (*dml_service.UpdateResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	return s.inner.Update(ctx, req)
}

func (s *DMLServiceContext) Delete(ctx context.Context, req *dml_service.DeleteRequest) (*dml_service.DeleteResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	return s.inner.Delete(ctx, req)
}
