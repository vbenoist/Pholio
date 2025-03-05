<template>
  <div>
    <ImagePicker v-model="photos" />
    <div class="image__preview">
      <ImagePreview v-model="photos" />
    </div>

  </div>
</template>

<script setup lang="ts">
import type { Ref } from 'vue'
import type { UploadableFile } from '@/composables/fileManager'
import { ref } from 'vue'
import ImagePicker from "@/components/Admin/ImagePicker.vue"
import ImagePreview from "@/components/Admin/ImagePicker/ImagePreview.vue"

const photos: Ref<Array<UploadableFile>> = ref([])


</script>

<style scoped lang="scss">

.image__preview {
  margin: auto;
  width: 80vw;
}

</style>
