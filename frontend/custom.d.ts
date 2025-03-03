import 'pinia'
import type { ApiResolver } from "./src/plugins/apiResolver"

declare module 'pinia' {
  export interface PiniaCustomProperties {
    api: ApiResolver
  }
}
