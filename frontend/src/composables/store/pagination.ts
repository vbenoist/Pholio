import type { Ref } from 'vue'
import { ref } from 'vue'
import type { ApiResolverCallable } from '@/plugins/apiResolver'
import type { PaginatedResults } from '@/models/api/paginated'
import { PaginationQuery } from '@/models/api/paginated'

export default <T>(apiResolverClbk: ApiResolverCallable, content: Ref<T[]>) => {
  const fetched: Ref<boolean> = ref(false)
  const reachedEnd: Ref<boolean> = ref(false)
  const pagination: Ref<PaginationQuery> = ref(new PaginationQuery({ perPage: 3 }))

  const initContent = async () => {
    const res = await apiResolverClbk(pagination.value)
    if(!res) return

    content.value = (res as PaginatedResults<T>).documents ?? []
    fetched.value = true
    pagination.value.page++
  }

  const fetchContent = async () => {
    if (content.value.length === 0 && !fetched.value) {
      await initContent()
    }

    return content
  }

  const fetchNextContent = async() => {
    if(reachedEnd.value) {
      return
    }

    if (!fetched.value) {
      return fetchContent()
    }

    const res = await apiResolverClbk(pagination.value)
    if(!res || !res.documents) {
      reachedEnd.value = true
      return
    }

    content.value = content.value.concat(res.documents)
    pagination.value.page++
  }

  const getContent = () => {
    return content.value
  }

  const hasFetched = () => {
    return fetched
  }

  return {
    content,
    fetched,
    hasFetched,
    fetchContent,
    fetchNextContent,
    getContent,
    initContent
  }
}
