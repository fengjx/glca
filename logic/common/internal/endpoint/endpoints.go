package endpoint

import (
	"context"

	"github.com/fengjx/daox"
	"github.com/fengjx/luchen"
	"github.com/go-kit/kit/endpoint"

	"github.com/fengjx/glca/logic/common/internal/service"
)

func Init(_ context.Context, httpServer *luchen.HTTPServer) {
	httpServer.Handler(newAdminCommonHandler())
}

func MakeQueryEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		query := request.(*daox.QueryRecord)
		pageVO, err := service.CommonService.Query(ctx, *query)
		if err != nil {
			return nil, err
		}
		return pageVO.ToAmisVO(), nil
	}
}

func MakeGetEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		record := request.(*daox.GetRecord)
		data, err := service.CommonService.Get(ctx, *record)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
}

func MakeInsertEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		record := request.(*daox.InsertRecord)
		id, err := service.CommonService.Insert(ctx, *record)
		if err != nil {
			return nil, err
		}
		return map[string]any{
			"id": id,
		}, nil
	}
}

func MakeUpdateEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		record := request.(*daox.UpdateRecord)
		affected, err := service.CommonService.Update(ctx, *record)
		if err != nil {
			return nil, err
		}
		return map[string]any{
			"affected": affected,
		}, nil
	}
}

func MakeDeleteEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		record := request.(*daox.DeleteRecord)
		affected, err := service.CommonService.Delete(ctx, *record)
		if err != nil {
			return nil, err
		}
		return map[string]any{
			"affected": affected,
		}, nil
	}
}
