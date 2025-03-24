type ApiRecord = {
  description: string | null
  location: string
  date: Date
}

export type ApiAddRecord = ApiRecord & {
  draftId: string
}

export type ApiGetRecord = ApiRecord & {
  id: string
}
