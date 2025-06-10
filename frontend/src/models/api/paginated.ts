type PaginationQueryT = {
	page: number
	perPage: number
	sortBy: string
	sortAsc: number
}

export type PaginatedResult<T> = {
  pagination: PaginationQueryT,
  document: T
}

export type PaginatedResults<T> = {
  pagination: PaginationQueryT,
  documents: T[]
}

// const paginationQueryDefaults = {
//   page: 1,
//   perPage: 20,
//   sortBy: '',
//   sortAsc: 1
// }

// export class PaginationQuery {
//   page: number
// 	perPage: number
// 	sortBy: string
// 	sortAsc: number

//   constructor({
//     page = paginationQueryDefaults.page,
//     perPage = paginationQueryDefaults.perPage,
//     sortBy = paginationQueryDefaults.sortBy,
//     sortAsc = paginationQueryDefaults.sortAsc
//   } = { ...paginationQueryDefaults }) {
//     this.page = page
//     this.perPage = perPage
//     this.sortBy = sortBy
//     this.sortAsc = sortAsc
//   }
// }


export class PaginationQuery {
  page: number
	perPage: number
	sortBy: string
	sortAsc: number

  constructor(cparams: Partial<PaginationQueryT> | null = null) {
    this.page = cparams?.page ?? 1
    this.perPage = cparams?.perPage ?? 20
    this.sortBy = cparams?.sortBy ?? ''
    this.sortAsc = cparams?.sortAsc ?? 1
  }
}
