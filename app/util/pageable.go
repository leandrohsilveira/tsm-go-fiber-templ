package util

type PageParams struct {
	Page  int32 `json:"page" form:"page" query:"page"`
	Limit int32 `json:"limit" form:"limit" query:"limit"`
}

type PageResult[T any] struct {
	Items []T   `json:"items"`
	Count int64 `json:"count"`
	Pages int64 `json:"pages"`
}

func NewPageResult[T any](items []T, count int64, limit int32) PageResult[T] {
	return PageResult[T]{
		Items: items,
		Count: count,
		Pages: calcPages(count, int64(limit)),
	}
}

func (params *PageParams) Offset() int32 {
	return params.Page*params.Limit - params.Limit
}

func (params *PageParams) SetDefaults() {
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.Limit <= 0 {
		params.Limit = 10
	}
}

func MapResult[T any, R any](result PageResult[T], mapper func(T) R) PageResult[R] {
	items := []R{}

	for _, item := range result.Items {
		mapped := mapper(item)
		items = append(items, mapped)
	}

	return PageResult[R]{
		Items: items,
		Count: result.Count,
		Pages: result.Pages,
	}
}

func calcPages(count int64, limit int64) int64 {
	if limit == 0 {
		return 0
	}
	if count%limit > 0 {
		return count/limit + 1
	}
	return count / limit
}
