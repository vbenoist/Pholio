import type {
  DetailedRecord,
  DraftRecord,
  GroupbyRecord,
  Record
} from '@/models/record'
import type {
  ApiAddRecord,
  ApiGetRecord,
  DetailedRecord as ApiDetailedRecord,
  GroupbyRecord as ApiGroupbyRecord
} from '@/models/api'

export const draftRecordToApiRecord = (draftRecord: DraftRecord): ApiAddRecord => {
  return {
    draftId: draftRecord.draftId,
    description: draftRecord.description,
    location: draftRecord.location!,
    date: draftRecord.date.toISOString(),
  }
}

export const detailedRecordToApiDetailedRecord = (draftRecord: DetailedRecord): ApiDetailedRecord => {
  return {
    id: draftRecord.id,
    description: draftRecord.description,
    location: draftRecord.location!,
    date: draftRecord.date.toISOString(),
    folder: draftRecord.folder,
    stared: draftRecord.stared
  }
}

export const apiRecordToRecord = (apiRecord: ApiGetRecord[]): Record[] => {
  return apiRecord.map((record) => {
    return {
      ...record,
      date: new Date(record.date)
    }
  })
}

export const apiGroupedRecordToGroupedRecord = (apiRecord: ApiGroupbyRecord[]): GroupbyRecord[] => {
  return apiRecord.map((record) => {
    return {
      'group-by': record['group-by'],
      results: record.results.map((result) => {
        return {
          ...result,
          date: new Date(result.date)
        }
      })
    }
  })
}

export const apiDetailedRecordToDetailedRecord = (apiRecord: ApiDetailedRecord[]): DetailedRecord[] => {
  return apiRecord.map((record) => {
    return {
      ...record,
      date: new Date(record.date),
      file: null,
      status: 'PENDING'
    }
  })
}
