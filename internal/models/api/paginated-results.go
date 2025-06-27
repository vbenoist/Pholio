package apimodels

import mongopagination "github.com/vbenoist/pholio/pkg/helpers/database"

type PaginatedResult[T any] struct {
	Pagination mongopagination.PaginationData `json:"pagination"`
	Document   T                              `json:"document"`
}

type PaginatedResults[T any] struct {
	Pagination mongopagination.PaginationData `json:"pagination"`
	Documents  []T                            `json:"documents"`
}
