import type { ApiAddRecord } from '@/models/api/record'
import type { DetailedRecord, DraftRecord } from '@/models/record'
import type { DetailedRecord as ApiDetailedRecord } from '@/models/api/detailed-record'

export const draftRecordToApiRecord = (draftRecord: DraftRecord): ApiAddRecord => {
  return {
    draftId: draftRecord.draftId,
    description: draftRecord.description,
    location: draftRecord.location!,
    date: draftRecord.date,
  }
}

export const detailedRecordToApiDetailedRecord = (draftRecord: DetailedRecord): ApiDetailedRecord => {
  return {
    id: draftRecord.id,
    description: draftRecord.description,
    location: draftRecord.location!,
    date: draftRecord.date,
    folder: draftRecord.folder,
    stared: draftRecord.stared
  }
}
