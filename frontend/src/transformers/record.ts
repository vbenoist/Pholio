import type { ApiAddRecord } from "@/models/api/record"
import type { DraftRecord } from "@/models/record"

export const draftRecordToApiRecord = (draftRecord: DraftRecord): ApiAddRecord => {
  return {
    draftId: draftRecord.draftId,
    description: draftRecord.description,
    location: draftRecord.location!,
    date: draftRecord.date
  }
}
