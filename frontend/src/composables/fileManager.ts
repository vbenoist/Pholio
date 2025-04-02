import type { Ref } from 'vue'
import { ref, toRaw } from 'vue'
import { UploadableFile } from '@/models/uploadableFile'

export default () => {
  const files: Ref<Array<UploadableFile>> = ref([])

  const addMergeFiles = (newFiles: Array<File>): Array<UploadableFile> => {
    const uploadable = newFiles.map((f) => new UploadableFile(f))
    const toAdd = uploadable.filter((nf) => !fileAlreadyPushed(nf))
    if (toAdd.length === 0) return files.value

    files.value = toRaw(files.value).concat(toAdd)
    return files.value
  }

  const fileAlreadyPushed = (file: UploadableFile): boolean => {
    return files.value.some((f) => f.id === file.id)
  }

  const convertFileListArray = (files: FileList): Array<File> => {
    return [...Array.from(files)]
  }

  return { addMergeFiles, convertFileListArray, files }
}
