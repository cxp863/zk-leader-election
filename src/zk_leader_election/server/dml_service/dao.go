package dml_service

import (
	"context"
	"zk-leader-election/src/zk_leader_election/idl/dml_service"
)

type dmlServiceDao struct {
	store map[string]string
}

func newDMLServiceDAO() DMLService {
	return &dmlServiceDao{store: make(map[string]string, 0)}
}

func (s *dmlServiceDao) Create(_ context.Context, req *dml_service.CreateRequest) (*dml_service.CreateResponse, error) {
	if req.Data != nil {
		s.store[req.Data.Key] = req.Data.Value
		return &dml_service.CreateResponse{Status: &dml_service.ResponseStatus{
			Code:    dml_service.ResponseCode_kOk,
			Message: "",
		}}, nil
	}
	return &dml_service.CreateResponse{Status: &dml_service.ResponseStatus{
		Code:    dml_service.ResponseCode_kInvalidParamError,
		Message: "data is nil",
	}}, nil
}

func (s *dmlServiceDao) Read(_ context.Context, req *dml_service.ReadRequest) (*dml_service.ReadResponse, error) {
	value, exist := s.store[req.Key]
	if exist {
		return &dml_service.ReadResponse{Status: &dml_service.ResponseStatus{
			Code:    dml_service.ResponseCode_kOk,
			Message: "",
		}, Data: &dml_service.DemoData{
			Key:   req.Key,
			Value: value,
		}}, nil
	}
	return &dml_service.ReadResponse{Status: &dml_service.ResponseStatus{
		Code:    dml_service.ResponseCode_kNotFound,
		Message: "",
	}}, nil
}

func (s *dmlServiceDao) Update(_ context.Context, req *dml_service.UpdateRequest) (*dml_service.UpdateResponse, error) {
	if req.Data == nil {
		return &dml_service.UpdateResponse{Status: &dml_service.ResponseStatus{
			Code:    dml_service.ResponseCode_kInvalidParamError,
			Message: "data is nil",
		}}, nil
	}

	_, exist := s.store[req.Data.Key]
	if exist {
		s.store[req.Data.Key] = req.Data.Value
		return &dml_service.UpdateResponse{Status: &dml_service.ResponseStatus{
			Code:    dml_service.ResponseCode_kOk,
			Message: "",
		}}, nil
	}
	return &dml_service.UpdateResponse{Status: &dml_service.ResponseStatus{
		Code:    dml_service.ResponseCode_kNotFound,
		Message: "",
	}}, nil
}

func (s *dmlServiceDao) Delete(_ context.Context, req *dml_service.DeleteRequest) (*dml_service.DeleteResponse, error) {
	_, exist := s.store[req.Key]
	if exist {
		delete(s.store, req.Key)
		return &dml_service.DeleteResponse{Status: &dml_service.ResponseStatus{
			Code:    dml_service.ResponseCode_kOk,
			Message: "",
		}}, nil
	}
	return &dml_service.DeleteResponse{Status: &dml_service.ResponseStatus{
		Code:    dml_service.ResponseCode_kNotFound,
		Message: "",
	}}, nil
}
