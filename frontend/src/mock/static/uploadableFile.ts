import fs from 'fs-js'
import { UploadableFile } from '@/models/uploadableFile'

export const mockUplaodableFile = (length: number): Array<UploadableFile> => {
  const result: Array<UploadableFile> = []

  const fData = fs.readFileSync('./src/mock/assets/midsize.jpeg')
  const file: File = new File(fData, 'midsize')

  for (let i = 0; i < length; i++) {
    result.push(new UploadableFile(file))
  }

  return result
}
