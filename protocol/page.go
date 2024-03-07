package protocol

import "github.com/fengjx/glca/connom/kit"

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
		Items: kit.OrElse[[]T](len(v.List) > 0, v.List, make([]T, 0)),
		Total: v.Count,
	}
}

type AmisPageVO[T any] struct {
	Items []T   `json:"items"`
	Total int64 `json:"total"`
}
