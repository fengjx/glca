package protocol

type BatchUpdateReq struct {
	TableName string           `json:"table_name"`
	Rows      []map[string]any `json:"rows"`
}

type DeleteByIDsReq struct {
	TableName string `json:"table_name"`
	IDs       []any  `json:"ids"`
	IDStr     string `json:"id_str"`
}
