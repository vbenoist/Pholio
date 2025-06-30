type ApiRecord = {
  description: string | null
  location: string
  date: string
}

export type ApiAddRecord = ApiRecord & {
  draftId: string
}

export type ApiGetRecord = ApiRecord & {
  id: string
}
