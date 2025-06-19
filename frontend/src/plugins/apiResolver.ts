import type { App } from 'vue'
import type { AxiosError, AxiosInstance, AxiosResponse } from 'axios'
import type { PaginatedResult, PaginatedResults } from '@/models/api/paginated'
import type { RecentlyContent } from '@/stores/recentlyContent'
import type { ApiAddRecord } from '@/models/api/record'
import { PaginationQuery } from '@/models/api/paginated'
import type { UploadableFile } from '@/models/uploadableFile'
import type { GroupbyRecord } from '@/models/api/groupby-record'
import type { DetailedRecord } from '@/models/api/detailed-record'

export type ApiResolverCallable = (pageParams: Partial<PaginationQuery> | null) => Promise<PaginatedResults<GroupbyRecord | DetailedRecord> | null>

export class ApiResolver {
  readonly axios: AxiosInstance

  constructor(axiosInst: AxiosInstance) {
    this.axios = axiosInst
  }

  fetchRecently = async (pageParams: Partial<PaginationQuery> | null): Promise<PaginatedResult<RecentlyContent> | null> => {
    const params = new PaginationQuery(pageParams)

    return this.axios
      .get('/content/records/recently', { params })
      .then((res: AxiosResponse) => {
        return res.data
      })
      .catch((e: AxiosError) => {
        console.log(e)
        return null
      })
  }

  fetchDetailed = async (pageParams: Partial<PaginationQuery> | null): Promise<PaginatedResults<DetailedRecord> | null> => {
    const params = new PaginationQuery(pageParams)

    return this.axios
      .get('/content/records/detailed', { params })
      .then((res: AxiosResponse) => {
        return res.data
      })
      .catch((e: AxiosError) => {
        console.log(e)
        return null
      })
  }

  fetchPerDate = async (pageParams: Partial<PaginationQuery> | null): Promise<PaginatedResults<GroupbyRecord> | null> => {
    const params = new PaginationQuery(pageParams)

    return this.axios
      .get('/content/records/per-date', { params })
      .then((res: AxiosResponse) => {
        return res.data
      })
      .catch((e: AxiosError) => {
        console.log(e)
        return null
      })
  }

  fetchPerLocation = async (pageParams: Partial<PaginationQuery> | null): Promise<PaginatedResults<GroupbyRecord> | null> => {
    const params = new PaginationQuery(pageParams)

    return this.axios
      .get('/content/records/per-location', { params })
      .then((res: AxiosResponse) => {
        return res.data
      })
      .catch((e: AxiosError) => {
        console.log(e)
        return null
      })
  }

  addRecord = async (record: ApiAddRecord): Promise<string | null> => {
    return this.axios
      .post('/content/record', record)
      .then((res: AxiosResponse) => {
        return res.data
      })
      .catch((e: AxiosError) => {
        console.log(e)
        return null
      })
  }

  linkImageRecord = async (recordId: string, upFile: UploadableFile): Promise<boolean> => {
    const form = new FormData()
    form.append('file', upFile.file)

    return this.axios
      .post(`/content/record/${recordId}/image`, form)
      .then(() => {
        return true
      })
      .catch((e: AxiosError) => {
        console.log(e)
        return false
      })
  }

  updateDraftRecord = async (record: ApiAddRecord): Promise<boolean> => {
    return this.axios
      .put(`/content/record/${record.draftId}`, record)
      .then(() => {
        return true
      })
      .catch((e: AxiosError) => {
        console.log(e)
        return false
      })
  }

  updateRecord = async (record: DetailedRecord): Promise<boolean> => {
    return this.axios
      .put(`/content/record/${record.id}`, record)
      .then(() => {
        return true
      })
      .catch((e: AxiosError) => {
        console.log(e)
        return false
      })
  }

  deleteRecord = async (recordId: string): Promise<boolean> => {
    return this.axios
      .delete(`/content/record/${recordId}`)
      .then(() => {
        return true
      })
      .catch((e: AxiosError) => {
        console.log(e)
        return false
      })
  }

  getLinkedThumb = async(recordId: string): Promise<File | null> => {
    return this.axios
    .get(`/content/record/${recordId}/image/thumb`, { responseType: 'blob' })
    .then((res: AxiosResponse) => {
      return new File([res.data  as BlobPart], `${recordId}-thumb`, {
        type: (res.headers['Content-Type'] as string) ?? 'image/jpeg'
      })
    })
    .catch((e: AxiosError) => {
      console.log(e)
      return null
    })
  }
}

export default {
  install(app: App) {
    const apiResolver = new ApiResolver(app.config.globalProperties.$axios)
    /* Making available as this.$apiResolver when in the vue instance */
    app.config.globalProperties.$apiResolver = apiResolver

    /* Making available outside vue instance, using inject (https://vuejs.org/api/application.html#app-provide) */
    app.provide('$apiResolver', apiResolver)
  },
}
