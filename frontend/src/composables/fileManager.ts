import type { Ref } from 'vue'
import { ref, toRaw } from 'vue'

export default () => {
  const files: Ref<Array<UploadableFile>> = ref([])

  const addMergeFiles = (newFiles: Array<File>): Array<UploadableFile> => {
    const uploadable = newFiles.map(f => new UploadableFile(f))
    const toAdd = uploadable.filter(nf => !fileAlreadyPushed(nf))
    if(toAdd.length === 0) return files.value

    files.value = toRaw(files.value).concat(toAdd)
    return files.value
  }

  const fileAlreadyPushed = (file: UploadableFile): boolean => {
    return files.value.some(f => f.id === file.id)
  }

  const convertFileListArray = (files: FileList): Array<File> => {
    return [...Array.from(files)]
  }

  return { addMergeFiles, convertFileListArray, files }
}

export type UploadableFileStatus = 'PENDING' | 'SENDING' | 'FAILED' | 'SENT'

export class UploadableFile {
  file: File
  id: string
  url: string
  status: boolean

	constructor(file: File) {
		this.file = file
		this.id = `${file.name}-${file.size}-${file.lastModified}-${file.type}`
		this.url = URL.createObjectURL(file)
		this.status = false
	}
}
