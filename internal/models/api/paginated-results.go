package apimodels

import mongopagination "github.com/gobeam/mongo-go-pagination"

type PaginatedResult[T any] struct {
	Pagination mongopagination.PaginationData `json:"pagination"`
	Document   T                              `json:"document"`
}

type PaginatedResults[T any] struct {
	Pagination mongopagination.PaginationData `json:"pagination"`
	Documents  []T                            `json:"documents"`
}
