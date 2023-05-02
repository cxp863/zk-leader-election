package dml_service

import (
	"context"
	"zk-leader-election/src/zk_leader_election/idl/dml_service"
)

type DMLService interface {
	Create(context.Context, *dml_service.CreateRequest) (*dml_service.CreateResponse, error)
	Read(context.Context, *dml_service.ReadRequest) (*dml_service.ReadResponse, error)
	Update(context.Context, *dml_service.UpdateRequest) (*dml_service.UpdateResponse, error)
	Delete(context.Context, *dml_service.DeleteRequest) (*dml_service.DeleteResponse, error)
}

type ReplaceableDMLService interface {
	DMLService
	RegisterServices(map[interface{}]DMLService)
	ActivateService(key interface{}) error
}
