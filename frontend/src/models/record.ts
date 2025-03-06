import { v4 as uuidv4 } from 'uuid'
import type { UploadableFile } from "./uploadableFile"

// export class Record {
//   id: string
//   nativImgSrc: string
//   midImgSrc:   string
//   thumbImgSrc: string
//   description: string | null
//   location:    string
//   date:        Date

//   constructor(file: File) {

//   }
// }

export class DraftRecord {
  draftId: string
  description: string | null
  location: string | null
  date: Date
  file: UploadableFile

  constructor(file: UploadableFile) {
    this.draftId = uuidv4()
    this.description = null
    this.location = ''
    this.date = new Date()
    this.file = file
  }
}
