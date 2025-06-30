import type { ApiGetRecord } from '@/models/api/record'

export type RecentlyRecord = {
  lastly: Array<ApiGetRecord>
  lately: Array<ApiGetRecord>
}
