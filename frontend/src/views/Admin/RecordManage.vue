<template>
  <div :ref="containerRefName">
    <span v-if="fetched">
      <ManagePublications v-model="records" />
    </span>
    <span v-else>Loading...</span>
  </div>
</template>

<script setup lang="ts">
import { computed, inject, ref, watch } from 'vue'
import { storeToRefs } from 'pinia'
import { cloneDeep } from 'lodash'
import scrollHook from '@/composables/scrollHook'
import ManagePublications from '@/components/Admin/Manage/ManagePublications.vue'
import { useDetailedStore } from '@/stores/detailedContent'
import type { ApiResolver } from '@/plugins/apiResolver'
import { type DetailedRecord as ApiDetailedRecord } from '@/models/api/detailed-record'
import { type DetailedRecord } from '@/models/record'
import { UploadableFile } from '@/models/uploadableFile'

const apiResolver = inject('$apiResolver') as ApiResolver
const detailedApiContentStore = useDetailedStore()
const { fetched } = storeToRefs(detailedApiContentStore)
const { containerRefName, registerHook } = scrollHook()

const apiDetailedRecord = computed(() => detailedApiContentStore.getContent())
const records = ref<Array<DetailedRecord>>([])
const workingRecords = ref<Array<string>>([])

const resolveContent = async () => {
  await detailedApiContentStore.fetchContent()
}

const onRecordsResolve = async () => {
  const prs: Promise<void>[] = []

  apiDetailedRecord.value.forEach (adrc => {
    const rcIdx = records.value.findIndex((rc) => rc.id === adrc.id)
    const wkIdx = workingRecords.value.findIndex((id) => id === adrc.id)

    if(rcIdx === -1 && wkIdx === -1) {
      workingRecords.value.push(adrc.id)
      prs.push(completeDetailedRecord(cloneDeep(adrc)))
    }
  })

  await Promise.all(prs)
}

const completeDetailedRecord = async(record: ApiDetailedRecord): Promise<void> => {
  workingRecords.value.push(record.id)

  /* Resolving record thumb */
  console.log("resolving: ", record.id)
  const imgThumb = await apiResolver.getLinkedThumb(record.id)
  if(!imgThumb) return

  records.value.push({
    ...record,
    date: new Date(record.date),
    status: 'PENDING',
    file: new UploadableFile(imgThumb)
  })
}

registerHook(() => {
  detailedApiContentStore.fetchNextContent()
})

watch(apiDetailedRecord, () => {
  onRecordsResolve()
})

resolveContent()

</script>
