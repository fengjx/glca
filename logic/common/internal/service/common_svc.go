package service

import (
	"context"
	"sync"

	"github.com/fengjx/daox"
	"github.com/fengjx/go-halo/json"
	"github.com/fengjx/luchen"
	"go.uber.org/zap"

	"github.com/fengjx/glca/connom/data/vo"
	"github.com/fengjx/glca/integration/db"
	"github.com/fengjx/glca/logic/common/commdto"
)

var CommonService *commonService

func init() {
	CommonService = newCommonService()
}

type commonService struct {
	sync.Mutex
	tableConfigMap map[string]commdto.TableConfig
}

type TableConfig struct {
	TableName          string
	Columns            []string
	InsertFieldsFilter daox.FieldsFilter
	InsertDataWrapper  daox.DataWrapper[map[string]any, map[string]any]
	SelectFieldsFilter daox.FieldsFilter
	SelectDataWrapper  daox.DataWrapper[any, any]
	UpdateFieldsFilter daox.FieldsFilter
	UpdateDataWrapper  daox.DataWrapper[map[string]any, map[string]any]
}

func newCommonService() *commonService {
	inst := &commonService{}
	inst.tableConfigMap = make(map[string]commdto.TableConfig)
	return inst
}

func (svc *commonService) RegisterTableConfig(config commdto.TableConfig) {
	svc.Lock()
	CommonService.tableConfigMap[config.TableName] = config
	luchen.RootLogger().Infof("register table config [%s]", config.TableName)
	svc.Unlock()
}

func (svc *commonService) Query(ctx context.Context, query daox.QueryRecord) (*vo.PageVO[map[string]any], error) {
	log := luchen.Logger(ctx)
	defaultDB := db.GetDefaultDB()
	config := svc.tableConfigMap[query.TableName]
	if len(query.Fields) == 0 {
		query.Fields = config.Columns
	}
	list, page, err := daox.FindListMap(ctx, defaultDB, query)
	if err != nil {
		log.Error("common query err", zap.Any("query", json.ToJsonDelay(query)), zap.Error(err))
		return nil, err
	}
	pageVO := &vo.PageVO[map[string]any]{
		List:    list,
		Offset:  page.Offset,
		Limit:   page.Limit,
		Count:   page.Count,
		HasNext: page.HasNext,
	}
	log.Info("page vo", zap.Any("page", pageVO))
	return pageVO, nil
}

func (svc *commonService) Get(ctx context.Context, record daox.GetRecord) (map[string]any, error) {
	log := luchen.Logger(ctx)
	defaultDB := db.GetDefaultDB()
	config := svc.tableConfigMap[record.TableName]
	if len(record.Fields) == 0 {
		record.Fields = config.Columns
	}
	data, err := daox.GetMap(ctx, defaultDB, record)
	if err != nil {
		log.Error("common get err", zap.Any("record", json.ToJsonDelay(record)), zap.Error(err))
		return nil, err
	}
	return data, nil
}

func (svc *commonService) Insert(ctx context.Context, record daox.InsertRecord) (int64, error) {
	log := luchen.Logger(ctx)
	defaultDB := db.GetDefaultDB()
	tableName := record.TableName
	fieldsFilter := func() daox.FieldsFilter {
		return func(ctx context.Context) []string {
			disableFields := []string{"id", "ctime"}
			if cfg, ok := svc.tableConfigMap[tableName]; ok && cfg.InsertFieldsFilter != nil {
				disableFields = append(disableFields, cfg.InsertFieldsFilter(ctx)...)
			}
			return disableFields
		}
	}

	dataWrapper := func() daox.DataWrapper[map[string]any, map[string]any] {
		return func(ctx context.Context, src map[string]any) map[string]any {
			if cfg, ok := svc.tableConfigMap[tableName]; ok && cfg.InsertDataWrapper != nil {
				return cfg.InsertDataWrapper(ctx, src)
			}
			return src
		}
	}

	id, err := daox.Insert(
		ctx,
		defaultDB,
		record,
		daox.WithInsertFieldsFilter(fieldsFilter()),
		daox.WithInsertDataWrapper(dataWrapper()),
	)
	if err != nil {
		log.Error("common insert err", zap.Any("record", json.ToJsonDelay(record)), zap.Error(err))
		return 0, err
	}
	log.Info("insert record", zap.Any("record", record))
	return id, nil
}

func (svc *commonService) Update(ctx context.Context, record daox.UpdateRecord) (int64, error) {
	log := luchen.Logger(ctx)
	defaultDB := db.GetDefaultDB()
	affected, err := daox.Update(ctx, defaultDB, record)
	if err != nil {
		log.Error("common update err", zap.Any("record", json.ToJsonDelay(record)), zap.Error(err))
		return 0, err
	}
	log.Info("update record", zap.Any("record", record))
	return affected, nil
}

func (svc *commonService) Delete(ctx context.Context, record daox.DeleteRecord) (int64, error) {
	log := luchen.Logger(ctx)
	defaultDB := db.GetDefaultDB()
	affected, err := daox.Delete(ctx, defaultDB, record)
	if err != nil {
		log.Error("common update err", zap.Any("record", json.ToJsonDelay(record)), zap.Error(err))
		return 0, err
	}
	log.Info("delete record", zap.Any("record", record))
	return affected, nil
}
