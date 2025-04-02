import type { App } from 'vue'
import axios from 'axios'

export const axiosInstance = axios.create({
  baseURL: 'http://localhost:8081',
  withCredentials: true,
})

export default {
  install(app: App) {
    app.config.globalProperties.$axios = { ...axiosInstance }
  },
}
