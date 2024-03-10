package commpub

import (
	"github.com/fengjx/daox"
	"github.com/fengjx/luchen"
)

type TableConfig struct {
	TableName              string
	PrimaryKey             string
	SelectColumns          []string
	SelectFieldsFilter     daox.FieldsFilter
	SelectDataWrapper      daox.DataWrapper[any, any]
	SelectConditionWrapper ConditionWrapper
	InsertFieldsFilter     daox.FieldsFilter
	InsertDataWrapper      daox.DataWrapper[map[string]any, map[string]any]
	UpdateFieldsFilter     daox.FieldsFilter
	UpdateDataWrapper      daox.DataWrapper[map[string]any, map[string]any]
	UpdateConditionWrapper ConditionWrapper
}

// ConditionWrapper where 条件二次组装
type ConditionWrapper func([]daox.Condition) []daox.Condition

var CommonAPI commonAPI

type commonAPI interface {
	// RegisterTableConfig 注册通用正删改查配置
	RegisterTableConfig(config TableConfig)
}

func SetCommonAPI(impl commonAPI) {
	luchen.RootLogger().Info("set CommonAPI")
	CommonAPI = impl
}
