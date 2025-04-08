import type { Ref } from 'vue'
import { inject, ref } from 'vue'
import { defineStore } from 'pinia'
import { ApiResolver } from '@/plugins/apiResolver'
import type { ApiGetRecord } from '@/models/api/record'

export type RecentlyContent = {
  lastly: Array<ApiGetRecord>
  lately: Array<ApiGetRecord>
}

export const useRecentlyContentStore = defineStore('recentlyContent', () => {
  let content: RecentlyContent | null = null
  const fetched: Ref<boolean> = ref(false)
  const apiResolver = inject('$apiResolver') as ApiResolver

  async function initContent() {
    const res = await apiResolver.fetchRecently()
    content = res?.document ?? null
    fetched.value = true
  }

  async function fetchContent() {
    if (content === null && !fetched.value) {
      await initContent()
    }

    return content
  }

  function getContent() {
    return content
  }

  function hasFetched() {
    return fetched
  }

  return { content, hasFetched, fetchContent, getContent, initContent }
})
