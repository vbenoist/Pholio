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

/*
Avoid "null" response. An empty array will result in [] after json marshalling.
A simple declaration of an empty array will result in "null"
*/
func (pr *PaginatedResults[T]) AvoidNullResults() {
	if len(pr.Documents) == 0 {
		pr.Documents = []T{}
	}
}
