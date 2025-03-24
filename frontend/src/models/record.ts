import { v4 as uuidv4 } from 'uuid'
import type { UploadableFile } from "./uploadableFile"

export type DraftRecordStatus = 'PENDING' | 'FAILED' | 'SENT'

export class DraftRecord {
  draftId: string
  description: string | null
  location: string | null
  date: Date
  file: UploadableFile
  status: DraftRecordStatus

  constructor(file: UploadableFile) {
    this.draftId = uuidv4()
    this.description = null
    this.location = ''
    this.date = new Date(file.file.lastModified)
    this.file = file
    this.status = 'PENDING'
  }
}
