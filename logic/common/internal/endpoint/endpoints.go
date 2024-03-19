package endpoint

import (
	"context"

	"github.com/fengjx/daox"
	"github.com/fengjx/go-halo/utils"
	"github.com/fengjx/luchen"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"

	"github.com/fengjx/glca/connom/errno"
	"github.com/fengjx/glca/logic/common/internal/service"
	"github.com/fengjx/glca/protocol"
)

func Init(_ context.Context, httpServer *luchen.HTTPServer) {
	httpServer.Handler(newAdminCommonHandler())
}

func MakeQueryEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		query := request.(*protocol.QueryReq)
		pageVO, err := service.CommonService.Query(ctx, query.QueryRecord)
		if err != nil {
			return nil, err
		}
		return pageVO.ToAmisResp(), nil
	}
}

func MakeGetEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		record := request.(*protocol.GetReq)
		data, err := service.CommonService.Get(ctx, record.GetRecord)
		if err != nil {
			return nil, err
		}
		return &protocol.GetResp{
			Record: data,
		}, nil
	}
}

func MakeInsertEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		record := request.(*protocol.InsertReq)
		id, err := service.CommonService.Insert(ctx, record.InsertRecord)
		if err != nil {
			return nil, err
		}
		return &protocol.InsertResp{
			ID: id,
		}, nil
	}
}

func MakeUpdateEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		record := request.(*protocol.UpdateReq)
		affected, err := service.CommonService.Update(ctx, record.UpdateRecord)
		if err != nil {
			return nil, err
		}
		return &protocol.UpdateResp{
			Affected: affected,
		}, nil
	}
}

// MakeBatchUpdateEndpoint 批量更新，要求必须包含主键字段
func MakeBatchUpdateEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log := luchen.Logger(ctx)
		req := request.(*protocol.BatchUpdateWithIDReq)
		config, ok := service.CommonService.GetTableConfig(req.TableName)
		if !ok {
			log.Warn(errno.TableNotSupportErr.Msg, zap.String("table_name", req.TableName))
			return nil, errno.TableNotSupportErr
		}
		ret := map[any]int64{}
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
		return &protocol.BatchUpdateResp{
			Affected: ret,
		}, nil
	}
}

func MakeDeleteEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		record := request.(*protocol.DeleteReq)
		affected, err := service.CommonService.Delete(ctx, record.DeleteRecord)
		if err != nil {
			return nil, err
		}
		return &protocol.DeleteResp{
			Affected: affected,
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
		return &protocol.DeleteByIDsResp{
			Affected: affected,
		}, nil
	}
}
