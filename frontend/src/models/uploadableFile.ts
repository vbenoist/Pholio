import ExifReader from 'exifreader'

export type UploadableFileStatus = 'PENDING' | 'SENDING' | 'FAILED' | 'SENT'

export class UploadableFile {
  file: File
  id: string
  url: string
  status: UploadableFileStatus

  constructor(file: File) {
    this.file = file
    this.id = `${file.name}-${file.size}-${file.lastModified}-${file.type}`
    this.url = URL.createObjectURL(file)
    this.status = 'PENDING'
  }
}

export const extractImageExifDate = async (upFile: UploadableFile): Promise<Date> => {
  const defaultDate = new Date(upFile.file.lastModified)

  if(upFile.file.type !== 'image/jpeg' && upFile.file.type !== 'image/png') {
    return defaultDate
  }

  const fileBuf = await upFile.file.arrayBuffer()
  const tags = await ExifReader.load(fileBuf)

  if(!tags.DateTimeOriginal) {
    return defaultDate
  }

  let date = tags.DateTimeOriginal.description.split(" ")[0]
  date = date.replaceAll(":", "-")

  const time = tags.DateTimeOriginal.description.split(" ")[1]
  const offsetTime = tags.OffsetTimeOriginal?.description ?? "+00:00"
  const dateTime = `${date}T${time}${offsetTime}`

  return new Date(dateTime)
}
