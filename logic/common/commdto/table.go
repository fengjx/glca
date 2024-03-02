package commdto

import "github.com/fengjx/daox"

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
