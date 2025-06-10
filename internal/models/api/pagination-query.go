package apimodels

type PaginationQuery struct {
	Page         int64
	PerPage      int64
	SortBy       string
	SortAsc      int8 // sort applied on all dbs elements
	SortAscGroup int8 // sort applied on group-by value
}
