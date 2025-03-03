import type { App } from "vue"
import type { AxiosError, AxiosInstance, AxiosResponse } from 'axios'
import type { RecentlyContent } from '@/stores/recentlyContent'

export class ApiResolver {
  readonly axios: AxiosInstance

  constructor(axiosInst: AxiosInstance) {
    this.axios = axiosInst
  }

  fetchRecently = async (): Promise<RecentlyContent | null> => {
    return this.axios.get("/content/recently")
      .then((res: AxiosResponse) => {
        return res.data
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
  }
}
