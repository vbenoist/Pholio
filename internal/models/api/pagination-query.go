package apimodels

type PaginationQuery struct {
	Page    int64
	PerPage int64
	SortBy  string
	SortAsc int8
}
