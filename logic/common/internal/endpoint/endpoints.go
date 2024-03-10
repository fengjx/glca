package endpoint

import (
	"context"

	"github.com/fengjx/daox"
	"github.com/fengjx/go-halo/utils"
	"github.com/fengjx/luchen"
	"go.uber.org/zap"

	"github.com/fengjx/glca/connom/endpoint"
	"github.com/fengjx/glca/connom/errno"
	"github.com/fengjx/glca/logic/common/internal/service"
	"github.com/fengjx/glca/protocol"
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

func MakeBatchUpdateEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log := luchen.Logger(ctx)
		req := request.(*protocol.BatchUpdateReq)
		config, ok := service.CommonService.GetTableConfig(req.TableName)
		if !ok {
			log.Warn(errno.TableNotSupportErr.Msg, zap.String("table_name", req.TableName))
			return nil, errno.TableNotSupportErr
		}
		ret := map[any]any{}
		if len(req.Rows) == 0 {
			return ret, nil
		}
		for _, row := range req.Rows {
			id, exist := row[config.PrimaryKey]
			if !exist {
				continue
			}
			affected, err := service.CommonService.Update(ctx, daox.UpdateRecord{
				TableName: req.TableName,
				Row:       row,
				Conditions: []daox.Condition{
					{
						Op:            daox.OpAnd,
						Field:         config.PrimaryKey,
						ConditionType: daox.ConditionTypeEq,
						Vals:          []any{id},
					},
				},
			})
			if err != nil {
				log.Error("update err",
					zap.String("table_name",
						req.TableName), zap.Error(err),
				)
				return ret, err
			}
			ret[id] = affected
		}
		return ret, nil
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

func MakeDeleteByIDsEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log := luchen.Logger(ctx)
		req := request.(*protocol.DeleteByIDsReq)
		config, ok := service.CommonService.GetTableConfig(req.TableName)
		if !ok {
			log.Warn(errno.TableNotSupportErr.Msg, zap.String("table_name", req.TableName))
			return nil, errno.TableNotSupportErr
		}
		ids := req.IDs
		if len(ids) == 0 {
			arr := utils.SplitTrim(req.IDStr, ",")
			for _, item := range arr {
				ids = append(ids, item)
			}
		}
		if len(ids) == 0 {
			return nil, errno.ArgsErr
		}
		record := daox.DeleteRecord{
			TableName: req.TableName,
			Conditions: []daox.Condition{
				{
					Op:            daox.OpAnd,
					Field:         config.PrimaryKey,
					ConditionType: daox.ConditionTypeIn,
					Vals:          ids,
				},
			},
		}
		affected, err := service.CommonService.Delete(ctx, record)
		if err != nil {
			return nil, err
		}
		log.Infof("delete table[%s], ids:[%v]", req.TableName, ids)
		return map[string]any{
			"affected": affected,
		}, nil
	}
}
