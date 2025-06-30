import type { Ref } from 'vue'
import { inject, ref } from 'vue'
import { defineStore } from 'pinia'
import { ApiResolver } from '@/plugins/apiResolver'
import type { RecentlyRecord } from '@/models/record'
import { PaginationQuery } from '@/models/api'

export const useRecentlyContentStore = defineStore('recentlyContent', () => {
  const content: Ref<RecentlyRecord | null> = ref(null)
  const fetched: Ref<boolean> = ref(false)
  const reachedEnd: Ref<boolean> = ref(false)
  const pagination: Ref<PaginationQuery> = ref(new PaginationQuery({ perPage: 3 }))
  const apiResolver = inject('$apiResolver') as ApiResolver

  const initContent = async () => {
    const res = await apiResolver.fetchRecently(pagination.value)
    if(!res) return

    content.value = res?.document ?? null
    fetched.value = true
    pagination.value.page++
  }

  const fetchContent = async () => {
    if (content.value === null && !fetched.value) {
      await initContent()
    }

    return content
  }

  const fetchNextContent = async () => {
    if(reachedEnd.value) {
      return
    }

    if (!fetched.value || !content.value) {
      return fetchContent()
    }

    const res = await apiResolver.fetchRecently(pagination.value)
    if(!res || (res.document.lastly.length === 0 && res.document.lately.length === 0)) {
      reachedEnd.value = true
      return
    }

    if(res.document.lastly && res.document.lastly.length > 0) {
      content.value.lastly = content.value.lastly.concat(res.document.lastly)
    }
    if(res.document.lately && res.document.lately.length > 0) {
      content.value.lately = content.value.lately.concat(res.document.lately)
    }

    pagination.value.page++
  }

  const getContent = () => {
    return content.value
  }

  const hasFetched = () => {
    return fetched
  }

  return {
    fetched,
    hasFetched,
    fetchContent,
    fetchNextContent,
    getContent,
    initContent
  }
})
