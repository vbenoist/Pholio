import type { App } from 'vue'
import type { AxiosError, AxiosInstance, AxiosResponse } from 'axios'
import type { PaginatedResult, PaginatedResults } from '@/models/api/paginated'
import type { RecentlyContent } from '@/stores/recentlyContent'
import type { ApiAddRecord } from '@/models/api/record'
import type { UploadableFile } from '@/models/uploadableFile'
import type { GroupbyRecord } from '@/models/api/groupby-record'

export class ApiResolver {
  readonly axios: AxiosInstance

  constructor(axiosInst: AxiosInstance) {
    this.axios = axiosInst
  }

  fetchRecently = async (): Promise<PaginatedResult<RecentlyContent> | null> => {
    return this.axios
      .get('/content/records/recently')
      .then((res: AxiosResponse) => {
        return res.data
      })
      .catch((e: AxiosError) => {
        console.log(e)
        return null
      })
  }

  fetchPerDate = async (): Promise<PaginatedResults<GroupbyRecord> | null> => {
    return this.axios
      .get('/content/records/per-date')
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
        console.log('addRecord', res.data)
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
