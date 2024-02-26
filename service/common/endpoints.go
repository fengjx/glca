package common

import (
	"context"

	"github.com/fengjx/daox"
	"github.com/go-kit/kit/endpoint"
)

type endpoints struct {
}

func newEndpoints() *endpoints {
	return &endpoints{}
}

func (e *endpoints) MakeQueryEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		query := request.(*daox.QueryRecord)
		pageVO, err := GetInst().CommonSvc.Query(ctx, *query)
		if err != nil {
			return nil, err
		}
		return pageVO, nil
	}
}

func (e *endpoints) MakeGetEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		record := request.(*daox.GetRecord)
		data, err := GetInst().CommonSvc.Get(ctx, *record)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
}

func (e *endpoints) MakeInsertEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		record := request.(*daox.InsertRecord)
		id, err := GetInst().CommonSvc.Insert(ctx, *record)
		if err != nil {
			return nil, err
		}
		return map[string]any{
			"id": id,
		}, nil
	}
}

func (e *endpoints) MakeUpdateEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		record := request.(*daox.UpdateRecord)
		affected, err := GetInst().CommonSvc.Update(ctx, *record)
		if err != nil {
			return nil, err
		}
		return map[string]any{
			"affected": affected,
		}, nil
	}
}

func (e *endpoints) MakeDeleteEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		record := request.(*daox.DeleteRecord)
		affected, err := GetInst().CommonSvc.Delete(ctx, *record)
		if err != nil {
			return nil, err
		}
		return map[string]any{
			"affected": affected,
		}, nil
	}
}
