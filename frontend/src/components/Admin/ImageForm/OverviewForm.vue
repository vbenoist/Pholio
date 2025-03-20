<template>
  <div>
    <ImagePreview v-if="records.length > 0" v-model="photos" orientation="vertical">
      <template #preview-extend="{ photo }">
        <PhotoCard :modelValue="retreiveDraftRecord(photo)" @update:modelValue="onDraftRecordUpdate" />
      </template>
    </ImagePreview>
  </div>
</template>

<script setup lang="ts">
import { defineModel, ref, toRaw, watch } from 'vue'
import { UploadableFile } from '@/models/uploadableFile'
import { DraftRecord } from '@/models/record'
import PhotoCard from '@/components/Admin/ImageForm/PhotoCard.vue'
import ImagePreview from '@/components/Admin/ImagePicker/ImagePreview.vue'

const photos = defineModel<Array<UploadableFile>>({ default: [] })
const records = ref<Array<DraftRecord>>([])

watch(photos, () => {
  updateRecords()
})

const retreiveDraftRecord = (photo: UploadableFile): DraftRecord => {
  return records.value.find(df => df.file.id === photo.id) ?? new DraftRecord(photo)
}

const onDraftRecordUpdate = (draftRecord: DraftRecord) => {
  const toUpdate = records.value.findIndex(df => df.draftId === draftRecord.draftId)
  if(toUpdate === -1) return
  records.value[toUpdate] = draftRecord
}

const updateRecords = () => {
  records.value = toRaw(photos.value).map(p => new DraftRecord(p))
}

updateRecords()

</script>

<style scoped lang="scss">

</style>
