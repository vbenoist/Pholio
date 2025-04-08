type PaginationQuery = {
	page: number
	perPage: number
	sortBy: string
	sortAsc: number
}

export type PaginatedResult<T> = {
  pagination: PaginationQuery,
  document: T
}

export type PaginatedResults<T> = {
  pagination: PaginationQuery,
  documents: T[]
}
