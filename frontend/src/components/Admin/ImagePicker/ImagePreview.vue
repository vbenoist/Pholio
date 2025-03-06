<template>
  <div class="preview">
    <div
      v-for="(photo, pk) in photos"
      :key="`photo-prev-${pk}`"
      class="preview__container"
    >
      <img class="preview__container__img" :src="photo.url" />
      <div class="preview__container__overlay">
        <div
          class="preview__container__overlay__actions"
          @click="displayFullsize(photo)"
        >
          <v-icon name="co-fullscreen" scale="1.3" />
          <v-icon name="io-trash-bin-sharp" scale="1.3" @click.stop="removePhoto(photo)" />
        </div>
      </div>
    </div>

    <ImageModal
      v-model:is-open="isModalOpen"
      v-model:photo="selectedPhoto"
      @close="onCloseModal"
    />
  </div>
</template>

<script setup lang="ts">
import { defineModel, ref } from 'vue'
import type { UploadableFile } from '@/composables/fileManager'
import ImageModal from "@/components/Admin/ImagePicker/ImageModal.vue"

const photos = defineModel<Array<UploadableFile>>()
const isModalOpen = ref<boolean>(false)
const selectedPhoto = ref<UploadableFile | null>(null)

const displayFullsize = (photo: UploadableFile) => {
  console.log("displayFullsize", photo)
  isModalOpen.value = true
  selectedPhoto.value = photo
}

const removePhoto = (photo: UploadableFile) => {
  console.log("removePhoto", photo)
  const idx = photos.value?.findIndex(p => p.id === photo.id) ?? -1

  if(idx === -1) return
  photos.value?.splice(idx, 1)
}

const onCloseModal = () => {
  selectedPhoto.value = null
}

</script>

<style scopped lang="scss">
$blur-bg: rgba(0, 0, 0, 0.4);

.preview {
  display: flex;
  flex-flow: row nowrap;
  /* https://dev.to/janeori/css-fix-when-justify-content-space-evenly-overflows-un-center-the-content-4l50 */
  justify-content: safe center;
  align-items: center;
  overflow: scroll;

  &__container {
    margin: 0 10px;
    position: relative;

    &:hover {
      .preview__container__img {
        -webkit-filter: blur(2px);
        filter: blur(2px);
      }

      .preview__container__overlay {
        opacity: 1;
      }
    }

    &__img {
      max-width: 250px;
      max-height: 125px;
      transition: .3s all;
      border-radius: 3px;
    }

    &__overlay {
      position: absolute;
      top: 0;
      left: 0;
      bottom: 0;
      width: 100%;
      background: $blur-bg;
      color: #fff;
      opacity: 0;
      transition: .3s all;
      border-radius: 3px;
      cursor: pointer;

      &__actions {
        display: flex;
        justify-content: space-evenly;
        align-items: center;
        height: 100%;
      }
    }
  }
}
</style>
