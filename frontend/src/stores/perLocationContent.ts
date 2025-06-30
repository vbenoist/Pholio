import type { Ref } from 'vue'
import { inject, ref } from 'vue'
import { defineStore } from 'pinia'
import type { GroupbyRecord } from '@/models/record'
import { ApiResolver } from '@/plugins/apiResolver'
import PaginatedStore from '@/composables/store/pagination'
import { mergeGroupByResults } from '@/helpers/store/merge'

export const usePerLocationStore = defineStore('perLocationContent', () => {
  const apiResolver = inject('$apiResolver') as ApiResolver
  const content: Ref<Array<GroupbyRecord>> = ref([])

  const { fetched,
    hasFetched,
    fetchContent,
    fetchNextContent,
    getContent,
    initContent
  } = PaginatedStore(apiResolver.fetchPerLocation, mergeGroupByResults, content)

  return {
    fetched,
    hasFetched,
    fetchContent,
    fetchNextContent,
    getContent,
    initContent
  }
})

