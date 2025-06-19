<template>
  <div :class="['preview', `preview--${orientation}`]">
    <div
      v-for="(photo, pk) in photos"
      :key="`photo-prev-${orientation}-${pk}`"
      :class="['preview__block', `preview__block--${orientation}`]"
    >
      <div class="preview__container">
        <img class="preview__container__img" :src="photo.url" />
        <div class="preview__container__overlay">
          <div class="preview__container__overlay__actions" @click="displayFullsize(photo)">
            <v-icon name="co-fullscreen" scale="1.3" />
            <v-icon name="io-trash-bin-sharp" scale="1.3" @click.stop="removePhoto(photo)" />
          </div>
        </div>
      </div>

      <slot name="preview-extend" :photo="photo"></slot>
    </div>

    <ImageModal v-model:is-open="isPreviewModalOpen" v-model:photo="selectedPhoto" @close="onCloseModal" />
    <AckDeleteModal v-model:is-open="isDeleteModalOpen" @delete="onAckRemovePhoto" @close="onCloseModal" />
  </div>
</template>

<script setup lang="ts">
import { defineModel, defineProps, ref } from 'vue'
import type { UploadableFile } from '@/models/uploadableFile'
import ImageModal from '@/components/Admin/ImagePicker/ImageModal.vue'
import AckDeleteModal from '@/components/Admin/ImagePicker/AckDeleteModal.vue'

type OrientationType = 'horizontal' | 'vertical'

const emit = defineEmits(['deletePhoto'])

const { orientation = 'horizontal' } = defineProps<{
  orientation?: OrientationType
}>()

const photos = defineModel<Array<UploadableFile>>()
const preventDelete = defineModel<boolean>('prevent-delete', { default: false })
const customDelete = defineModel<boolean>( 'custom-delete', { default: false })

const isPreviewModalOpen = ref<boolean>(false)
const isDeleteModalOpen = ref<boolean>(false)
const selectedPhoto = ref<UploadableFile | null>(null)

const displayFullsize = (photo: UploadableFile) => {
  isPreviewModalOpen.value = true
  selectedPhoto.value = photo
}

const removePhoto = (photo: UploadableFile, ackDelete: boolean = false) => {
  if (preventDelete.value && !ackDelete) {
    isDeleteModalOpen.value = true
    selectedPhoto.value = photo
    console.log('prevent delete')
    return
  }

  if (customDelete.value) {
    console.log('custom delete')
    emit('deletePhoto', photo)
    return
  }

  const idx = photos.value?.findIndex((p) => p.id === photo.id) ?? -1
  if (idx === -1) return
  photos.value?.splice(idx, 1)
}

const onAckRemovePhoto = () => {
  if(!selectedPhoto.value) return
  removePhoto(selectedPhoto.value, true)
  isDeleteModalOpen.value = false
}

const onCloseModal = () => {
  selectedPhoto.value = null
}
</script>

<style scopped lang="scss">
$blur-bg: rgba(0, 0, 0, 0.4);

.preview {
  display: flex;
  /* https://dev.to/janeori/css-fix-when-justify-content-space-evenly-overflows-un-center-the-content-4l50 */
  justify-content: safe center;
  align-items: center;
  overflow: scroll;

  &--vertical {
    flex-flow: column nowrap;
  }

  &--horizontal {
    flex-flow: row nowrap;
    margin: 10px 0;
  }

  /* The overrall block englobing container & slot */
  &__block {
    display: flex;
    margin: 10px 0;

    &--vertical {
      flex-flow: row nowrap;
      border: 1px;
      border-style: none solid none none;
      box-shadow: 10px 0px 10px -5px #3c3c3c;
    }

    &--horizontal {
      flex-flow: column nowrap;
    }
  }

  /* Img container with overlay */
  &__container {
    margin: auto 10px;
    position: relative;
    width: fit-content;
    height: fit-content;

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
      transition: 0.3s all;
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
      transition: 0.3s all;
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
