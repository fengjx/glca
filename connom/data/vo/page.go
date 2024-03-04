package vo

type PageVO[T any] struct {
	List    []T   `json:"list"`
	Offset  int64 `json:"offset"`
	Limit   int64 `json:"limit"`
	Count   int64 `json:"count"`
	HasNext bool  `json:"has_next"`
}

func (v *PageVO[T]) ToAmisVO() *AmisPageVO[T] {
	if v == nil {
		return nil
	}
	return &AmisPageVO[T]{
		Items:   v.List,
		Count:   v.Count,
		HasNext: v.HasNext,
	}
}

type AmisPageVO[T any] struct {
	Items   []T   `json:"items"`
	Count   int64 `json:"count"`
	HasNext bool  `json:"hasNext"`
}
