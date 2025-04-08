import type { ApiGetRecord } from "@/models/api/record"

export type GroupbyRecord = {
	'group-by': string,
	'results': ApiGetRecord[]
}
