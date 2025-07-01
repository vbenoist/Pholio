import type { Ref } from 'vue'
import { inject, ref } from 'vue'
import { defineStore } from 'pinia'
import { cloneDeep } from 'lodash'
import type { DetailedRecord } from '@/models/record'
import type { PaginatedResults, PaginationQuery } from '@/models/api'
import { UploadableFile } from '@/models/uploadableFile'
import { ApiResolver, type ApiResolverCallable } from '@/plugins/apiResolver'
import PaginatedStore from '@/composables/store/pagination'
import { mergeDetailedResults } from '@/helpers/store/merge'

export const useDetailedStore = defineStore('detailedContent', () => {
  const apiResolver = inject('$apiResolver') as ApiResolver
  const content: Ref<Array<DetailedRecord>> = ref([])

  const customResolver: ApiResolverCallable = async (pageParams: Partial<PaginationQuery> | null): Promise<PaginatedResults<DetailedRecord> | null> => {
    const apiResults = await apiResolver.fetchDetailed(pageParams)
    if(!apiResults) return null

    const prs: Promise<void>[] = []

    apiResults.documents.forEach(detRecord => {
      prs.push(new Promise((resolve, _) => {
        resolveThumbRecord(cloneDeep(detRecord))
          .then(drWithThumb => {
            const idx = apiResults.documents.findIndex((dr) => dr.id === drWithThumb.id)
            if(idx === -1) return
            apiResults.documents[idx] = drWithThumb
            resolve()
          })
      }))
    })

    await Promise.all(prs)
    return apiResults
  }

  const resolveThumbRecord = async(record: DetailedRecord): Promise<DetailedRecord> => {
    const imgThumb = await apiResolver.getLinkedThumb(record.id)
    if(!imgThumb) return record


    return {
      ...record,
      status: 'PENDING',
      file: new UploadableFile(imgThumb)
    }
  }

  const { fetched,
    hasFetched,
    fetchContent,
    fetchNextContent,
    getContent,
    initContent
  } = PaginatedStore(customResolver, mergeDetailedResults, content)

  return {
    fetched,
    hasFetched,
    fetchContent,
    fetchNextContent,
    getContent,
    initContent
  }
})
