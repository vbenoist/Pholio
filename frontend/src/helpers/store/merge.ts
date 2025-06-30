import type { Ref } from 'vue'
import {
  type GroupbyRecord,
  type DetailedRecord
} from '@/models/api/'
import { type StoreMergeMethod } from '@/composables/store/pagination'

export const mergeDetailedResults: StoreMergeMethod = <T>(content: Ref<T[]>, newContent: GroupbyRecord[] | DetailedRecord[]) => {
  const storeContent = content as Ref<DetailedRecord[]>
  const newApiContent = newContent as DetailedRecord[]
  storeContent.value = storeContent.value.concat(newApiContent)
}

export const mergeGroupByResults: StoreMergeMethod = <T>(content: Ref<T[]>, newContent: GroupbyRecord[] | DetailedRecord[]) => {
  const storeContent = content as Ref<GroupbyRecord[]>
  const newApiContent = newContent as GroupbyRecord[]

  newApiContent.forEach((grp) => {
    const scIdx = storeContent.value.findIndex((sc) => sc['group-by'] === grp['group-by'])

    if(scIdx === -1) {
      storeContent.value.push(grp)
    } else {
      storeContent.value[scIdx]['results'] = storeContent.value[scIdx]['results'].concat(grp['results'])
    }
  })
}
