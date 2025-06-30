import type { Ref } from 'vue'
import { inject, ref } from 'vue'
import { defineStore } from 'pinia'
import type { DetailedRecord } from '@/models/api/detailed-record'
import { ApiResolver } from '@/plugins/apiResolver'
import PaginatedStore from '@/composables/store/pagination'
import { mergeDetailedResults } from '@/helpers/store/merge'

export const useDetailedStore = defineStore('detailedContent', () => {
  const apiResolver = inject('$apiResolver') as ApiResolver
  const content: Ref<Array<DetailedRecord>> = ref([])

  const { fetched,
    hasFetched,
    fetchContent,
    fetchNextContent,
    getContent,
    initContent
  } = PaginatedStore(apiResolver.fetchDetailed, mergeDetailedResults, content)

  return {
    fetched,
    hasFetched,
    fetchContent,
    fetchNextContent,
    getContent,
    initContent
  }
})
