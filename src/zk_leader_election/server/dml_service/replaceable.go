package dml_service

import (
	"context"
	"errors"
	"sync"
	"zk-leader-election/src/zk_leader_election/idl/dml_service"
)

func newDMLServiceReplaceable() ReplaceableDMLService {
	return &DMLServiceReplaceable{
		optional: make(map[interface{}]DMLService, 0),
		lock:     &sync.RWMutex{},
	}
}

type DMLServiceReplaceable struct {
	current  DMLService
	optional map[interface{}]DMLService
	lock     *sync.RWMutex
}

func (s *DMLServiceReplaceable) RegisterServices(optional map[interface{}]DMLService) {
	s.lock.Lock()
	s.optional = optional
	s.lock.Unlock()
}

func (s *DMLServiceReplaceable) ActivateService(key interface{}) error {
	replace, exist := s.optional[key]
	if !exist {
		return errors.New("not found")
	}

	s.lock.Lock()
	s.current = replace
	s.lock.Unlock()
	return nil
}

func (s *DMLServiceReplaceable) Create(ctx context.Context, req *dml_service.CreateRequest) (*dml_service.CreateResponse, error) {
	s.lock.RLock()
	resp, err := s.current.Create(ctx, req)
	s.lock.RUnlock()
	return resp, err
}

func (s *DMLServiceReplaceable) Read(ctx context.Context, req *dml_service.ReadRequest) (*dml_service.ReadResponse, error) {
	s.lock.RLock()
	resp, err := s.current.Read(ctx, req)
	s.lock.RUnlock()
	return resp, err
}

func (s *DMLServiceReplaceable) Update(ctx context.Context, req *dml_service.UpdateRequest) (*dml_service.UpdateResponse, error) {
	s.lock.RLock()
	resp, err := s.current.Update(ctx, req)
	s.lock.RUnlock()
	return resp, err
}

func (s *DMLServiceReplaceable) Delete(ctx context.Context, req *dml_service.DeleteRequest) (*dml_service.DeleteResponse, error) {
	s.lock.RLock()
	resp, err := s.current.Delete(ctx, req)
	s.lock.RUnlock()
	return resp, err
}
