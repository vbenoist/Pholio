<template>
  <dialog
    ref="image-modal-dial"
    class="image-modal"
    @close="onCloseModal"
  >
    <div class="image-modal__container">
      <img
        v-if="photo"
        class="image-modal__container__img"
        :src="photo.url"
      />

      <div class="image-modal__container__close" @click="manuallyCloseModal">
        <v-icon class="image-modal__container__close__icon" name="io-close" scale="1.6" />
      </div>
    </div>
  </dialog>
</template>

<script setup lang="ts">
import type { Ref } from 'vue'
import { defineModel, useTemplateRef, watch } from 'vue'
import type { UploadableFile } from '@/models/uploadableFile'

const emit = defineEmits<{
  close: []
}>()
const dialogRef: Ref<HTMLDialogElement | null> = useTemplateRef('image-modal-dial')
const isModalOpen = defineModel<boolean>('is-open', { default: false })
const photo = defineModel<UploadableFile | null>('photo', { default: null })

const openModal = () => {
  if(!dialogRef.value) return
  dialogRef.value.showModal()
}

const onCloseModal = () => {
  isModalOpen.value = false
  emit('close')
}

const manuallyCloseModal = () => {
  if(!dialogRef.value) return
  dialogRef.value.close()
}

watch(isModalOpen, () => {
  if(!dialogRef.value) return
  if(isModalOpen.value && !dialogRef.value.open) openModal()
})

</script>

<style scoped lang="scss">
dialog::backdrop {
  backdrop-filter: blur(3px);
}

dialog[open] {
  animation: modalFadeIn .5s ease normal;
}

@keyframes modalFadeIn{
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

.image-modal {
  max-width: 96vw;
  height: 96vh;
  margin: auto;
  padding: 0;
  background-color: #3d352b;
  border: 2px solid;
  border-radius: 5px;

  &__container {
    display: flex;
    height: 100%;
    align-items: center;

    &__img {
      max-height: 100%;
    }

    &__close {
      position: absolute;
      top: 0;
      right: 0;
      padding: 6px;
      cursor: pointer;
      background-color: rgba(36, 31, 25, 0.6);
      border-radius: 0 0 0 25px;

      &__icon {
        color: white;
      }
    }
  }
}
</style>
