<template>
  <div>
    <ImagePreview
      v-if="records.length > 0"
      v-model="photos"
      orientation="vertical"
      prevent-delete
      custom-delete
      @delete-photo="onPhotoDeletion"
    >
      <template #preview-extend="{ photo }">
        <PhotoCard
          :modelValue="retreiveRecord(photo)"
          @update:modelValue="onRecordUpdate"
        />
      </template>
    </ImagePreview>
  </div>
</template>

<script setup lang="ts">
import { computed, defineModel, inject } from 'vue'
import { UploadableFile } from '@/models/uploadableFile'
import type { DetailedRecord, DraftRecord } from '@/models/record'
import PhotoCard from '@/components/Admin/ImageForm/PhotoCard.vue'
import ImagePreview from '@/components/Admin/ImagePicker/ImagePreview.vue'
import type { ApiResolver } from '@/plugins/apiResolver'

const apiResolver = inject('$apiResolver') as ApiResolver
const records = defineModel<Array<DetailedRecord>>({default: []})

const photos = computed<UploadableFile[]>(() => {
  return records.value.map(dr => dr.file)
})

const retreiveRecord = (p: UploadableFile): DetailedRecord => {
  return records.value.find(dr => dr.file.id === p.id) as DetailedRecord
}

const onRecordUpdate = (detRecord: DetailedRecord | DraftRecord) => {
  const res = detRecord as DetailedRecord
  const toUpdate = records.value.findIndex((df) => df.id === res.id)
  if (toUpdate === -1) return
  records.value[toUpdate] = res
}

const onPhotoDeletion = (photo: UploadableFile) => {
  const idx = records.value.findIndex((dr) => dr.file.id === photo.id)
  if (idx === -1) return

  apiResolver.deleteRecord(records.value[idx].id)
    .then(() => {
      records.value.splice(idx, 1)
    })
}

</script>

<style scoped lang="scss"></style>
