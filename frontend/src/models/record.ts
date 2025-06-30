import { v4 as uuidv4 } from 'uuid'
import { extractImageExifDate, type UploadableFile } from './uploadableFile'
import type {
  ApiGetRecord,
  DetailedRecord as ApiDetailedRecord
} from './api'

export type DraftRecordStatus = 'PENDING' | 'SENDING' | 'FAILED' | 'SENT'

export class DraftRecord {
  draftId: string
  description: string | null
  location: string | null
  date: Date
  file: UploadableFile
  status: DraftRecordStatus
  saved: boolean

  constructor(file: UploadableFile) {
    this.draftId = uuidv4()
    this.description = null
    this.location = ''
    this.date = new Date(file.file.lastModified)
    this.file = file
    this.status = 'PENDING'
    this.saved = false
  }

  public extractExifDate = async() => {
    const date = await extractImageExifDate(this.file)
    this.date = date
  }
}

export type Record = Omit<ApiGetRecord, "date"> & {
  date: Date
}

export type GroupbyRecord = {
  'group-by': string,
  'results': Record[]
}

export type DetailedRecord = Omit<ApiDetailedRecord, "date"> & {
  date: Date,
  file: UploadableFile | null
  status: DraftRecordStatus
}

export type RecentlyRecord = {
  lastly: Array<Record>
  lately: Array<Record>
}
