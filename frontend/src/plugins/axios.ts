import type { App } from "vue"
import axios from "axios"

const axiosInstance = axios.create({
  baseURL: "http://127.0.0.1:8081"
})

export default {
  install(app: App) {
    app.config.globalProperties.$axios = { ...axiosInstance }
  }
}
