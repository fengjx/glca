package dao

import (
	"context"
	"reflect"

	"github.com/fengjx/daox"

	"github.com/fengjx/glca/integration/db"
	"github.com/fengjx/glca/logic/common/commpub"
	"github.com/fengjx/glca/logic/sys/internal/data/entity"
	"github.com/fengjx/glca/logic/sys/internal/data/meta"
)

var ConfigDao = newSysConfigDao()

type sysConfigDao struct {
	*daox.Dao
}

func registerConfigTableConfig() {
	m := meta.SysConfigMeta
	fieldsFilter := func(ctx context.Context) []string {
		return nil
	}
	insertDataWrapper := func(ctx context.Context, src map[string]any) map[string]any {
		return src
	}

	updateDataWrapper := func(ctx context.Context, src map[string]any) map[string]any {
		return src
	}

	cfg := commpub.TableConfig{
		TableName:          m.TableName(),
		PrimaryKey:         m.PrimaryKey(),
		InsertFieldsFilter: fieldsFilter,
		InsertDataWrapper:  insertDataWrapper,
		UpdateFieldsFilter: fieldsFilter,
		UpdateDataWrapper:  updateDataWrapper,
	}
	commpub.CommonAPI.RegisterTableConfig(cfg)
}

func newSysConfigDao() *sysConfigDao {
	inst := &sysConfigDao{}
	inst.Dao = daox.NewDAO(
		db.GetDefaultDB(),
		"sys_config",
		"id",
		reflect.TypeOf(&entity.SysConfig{}),
	)
	return inst
}
