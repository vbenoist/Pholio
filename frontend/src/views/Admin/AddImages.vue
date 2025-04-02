<template>
  <div>
    <div v-if="step === 0">
      <ImagePicker v-model="photos" />
      <div class="image__preview">
        <ImagePreview v-model="photos" />
      </div>

      <div class="next-step">
        <button v-if="photos.length > 0" @click="goToNextStep">Etape suivante ></button>
      </div>
    </div>

    <div v-else-if="step === 1">
      <div class="test-overview">
        <OverviewForm v-model="photos" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { UploadableFile } from '@/models/uploadableFile'
import { ref } from 'vue'
import ImagePicker from '@/components/Admin/ImagePicker.vue'
import ImagePreview from '@/components/Admin/ImagePicker/ImagePreview.vue'
import OverviewForm from '@/components/Admin/ImageForm/OverviewForm.vue'

const photos = ref<Array<UploadableFile>>([])
const step = ref<number>(0)

const goToNextStep = () => {
  step.value++
}
</script>

<style scoped lang="scss">
.image__preview {
  margin: auto;
  width: 80vw;
}

.test-overview {
  margin-top: 20px;
}

.next-step {
  width: 100%;
  text-align: center;
  margin: 20px 0 20px 160px;

  button {
    cursor: pointer;
    color: var(--color-text);
    background-color: black;
    border: none;
    border-radius: 5px;
    padding: 10px;
  }
}
</style>
