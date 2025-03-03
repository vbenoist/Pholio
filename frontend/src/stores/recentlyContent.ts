import type { Ref } from 'vue'
import { inject, ref } from 'vue'
import { defineStore } from 'pinia'
import { ApiResolver } from "@/plugins/apiResolver"

export type RecentlyContent = {
  lastly: Array<Item>
  lately: Array<Item>
}

export type Item = {
  nativImgSrc: string
  midImgSrc: string
  thumbImgSrc: string
  description: string
  location: string
  date: Date
}

export const useRecentlyContentStore = defineStore('recentlyContent', () => {
  let content: RecentlyContent | null = null
  const fetched: Ref<boolean> = ref(false)
  const apiResolver = inject('$apiResolver') as ApiResolver

  async function initContent() {
    content = await apiResolver.fetchRecently()
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
