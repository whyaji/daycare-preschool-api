package types

type PaginationFilter struct {
	Page    int
	Limit   int
	Filters map[string]any
}
