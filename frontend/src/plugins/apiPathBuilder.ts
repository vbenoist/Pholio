import type { App } from "vue"
import type { AxiosInstance } from 'axios'
import type { ApiGetRecord } from '@/models/api/record'

export class ApiPathBuilder {
  readonly axios: AxiosInstance

  constructor(axiosInst: AxiosInstance) {
    this.axios = axiosInst
  }

  buildRecordThumbUrl = (record: ApiGetRecord): string => {
    return `${this.axios.defaults.baseURL}/content/record/${record.id}/image/thumb`
  }
}

export default {
  install(app: App) {
    const apiPathBuilder = new ApiPathBuilder(app.config.globalProperties.$axios)
    app.config.globalProperties.$apiPathBuilder = apiPathBuilder
    app.provide('$apiPathBuilder', apiPathBuilder)
  }
}
