package protocol

import "github.com/fengjx/daox"

type QueryReq struct {
	daox.QueryRecord
}

type AmisPageResp[T any] struct {
	Items []T   `json:"items"`
	Total int64 `json:"total"`
}

type GetReq struct {
	daox.GetRecord
}

type GetResp struct {
	Record any `json:"record"`
}

type InsertReq struct {
	daox.InsertRecord
}

type InsertResp struct {
	ID any `json:"id"`
}

type UpdateReq struct {
	daox.UpdateRecord
}

type UpdateResp struct {
	Affected int64 `json:"affected"`
}

type BatchUpdateWithIDReq struct {
	TableName string           `json:"table_name"`
	Rows      []map[string]any `json:"rows"`
}

type BatchUpdateResp struct {
	Affected map[any]int64 `json:"affected"`
}

type DeleteByIDsReq struct {
	TableName string `json:"table_name"`
	IDs       []any  `json:"ids"`
	IDStr     string `json:"id_str"`
}

type DeleteByIDsResp struct {
	Affected int64 `json:"affected"`
}

type DeleteReq struct {
	daox.DeleteRecord
}

type DeleteResp struct {
	Affected int64 `json:"affected"`
}
