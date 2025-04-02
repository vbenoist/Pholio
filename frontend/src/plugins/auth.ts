import type { App } from 'vue'
import type { AxiosInstance } from 'axios'
import type { User } from '@/models/api/user'
import type { RouteLocation } from 'vue-router'

const _defaultRedirName = 'RECENT'
const _authRedir = 'ADMIN-LOGIN'

export class Auth {
  readonly axios: AxiosInstance
  private loggedIn: boolean
  private overwrittenRoute: RouteLocation | null

  constructor(axiosInst: AxiosInstance) {
    this.axios = axiosInst
    this.loggedIn = false
    this.overwrittenRoute = null
  }

  getOverwrittenRoute = (): RouteLocation | null => {
    return this.overwrittenRoute
  }

  login = async (ids: User): Promise<boolean> => {
    return this.axios
      .post('/auth/login', ids)
      .then((res) => {
        if (res.data && res.data.token) {
          this.loggedIn = true
          return true
        }
        return false
      })
      .catch(() => false)
  }

  check = async (): Promise<boolean> => {
    return this.axios
      .get('/auth/check')
      .then(() => true)
      .catch(() => false)
  }

  routerHandler = async (to: RouteLocation, from: RouteLocation) => {
    /* checking if route has a 'auth' meta */
    if (!to.meta.auth && to.name !== _authRedir) {
      this.overwrittenRoute = null
      return true
    }

    /* If we don't have the flag (refreshed page for ex., checking from back) */
    if (!this.loggedIn) {
      this.loggedIn = await this.check()
    }

    if (this.loggedIn) {
      /* If already logged in and trying to reach login page, force go back */
      if (to.name === _authRedir) {
        return { name: from.name ?? _defaultRedirName, replace: true }
      }
      return true
    }

    if (to.name !== _authRedir) {
      // return { name: _authRedir, replace: true, params: { redirectName: from.name ?? '' }}
      this.overwrittenRoute = to
      return { name: _authRedir, replace: true }
    }
  }
}

export default {
  install(app: App) {
    const auth = new Auth(app.config.globalProperties.$axios)
    app.config.globalProperties.$auth = auth
    app.provide('$auth', auth)
    app.config.globalProperties.$router.beforeEach(auth.routerHandler)
  },
}
