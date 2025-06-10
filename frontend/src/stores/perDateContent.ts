import type { Ref } from 'vue'
import { inject, ref } from 'vue'
import { defineStore } from 'pinia'
import type { GroupbyRecord } from '@/models/api/groupby-record'
import { ApiResolver } from '@/plugins/apiResolver'
import PaginatedStore from '@/composables/store/pagination'

export const usePerDateStore = defineStore('perDateContent', () => {
  const apiResolver = inject('$apiResolver') as ApiResolver
  const content: Ref<Array<GroupbyRecord>> = ref([])

  const { fetched,
    hasFetched,
    fetchContent,
    fetchNextContent,
    getContent,
    initContent
  } = PaginatedStore(apiResolver.fetchPerDate, content)

  return {
    fetched,
    hasFetched,
    fetchContent,
    fetchNextContent,
    getContent,
    initContent
  }
})
