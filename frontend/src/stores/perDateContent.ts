import type { Ref } from 'vue'
import { inject, ref } from 'vue'
import { defineStore } from 'pinia'
import { ApiResolver } from '@/plugins/apiResolver'
import type { GroupbyRecord } from '@/models/api/groupby-record'

export const usePerDateStore = defineStore('perDateContent', () => {
  let content: Array<GroupbyRecord> = []
  const fetched: Ref<boolean> = ref(false)
  const apiResolver = inject('$apiResolver') as ApiResolver

  async function initContent() {
    const res = await apiResolver.fetchPerDate()
    content = res?.documents ?? []
    fetched.value = true
  }

  async function fetchContent() {
    if (content.length === 0 && !fetched.value) {
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
