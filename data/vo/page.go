package vo

type PageVO[T any] struct {
	List    []T   `json:"list"`
	Offset  int64 `json:"offset"`
	Limit   int64 `json:"limit"`
	Count   int64 `json:"count"`
	HasNext bool  `json:"has_next"`
}
