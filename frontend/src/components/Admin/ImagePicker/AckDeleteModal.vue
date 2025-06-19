<template>
  <dialog ref="delete-modal-dial" class="delete-modal" @close="onCloseModal">
    <div class="delete-modal__container">
      <span>Supprimer cette photo ?</span>
      <div class="delete-modal__container__buttons">
        <button @click="manuallyCloseModal">Annuler</button>
        <button @click="ackDelete">Supprimer</button>
      </div>
    </div>
  </dialog>
</template>

<script setup lang="ts">
import type { Ref } from 'vue'
import { defineModel, useTemplateRef, watch } from 'vue'

const emit = defineEmits<{
  close: [],
  delete: []
}>()
const dialogRef: Ref<HTMLDialogElement | null> = useTemplateRef('delete-modal-dial')
const isModalOpen = defineModel<boolean>('is-open', { default: false })

const openModal = () => {
  if (!dialogRef.value) return
  dialogRef.value.showModal()
}

const onCloseModal = () => {
  isModalOpen.value = false
  emit('close')
}

const manuallyCloseModal = () => {
  if (!dialogRef.value) return
  dialogRef.value.close()
}

const ackDelete = () => {
  isModalOpen.value = false
  emit('delete')
}

watch(isModalOpen, () => {
  if (!dialogRef.value) return
  if (isModalOpen.value && !dialogRef.value.open) openModal()
})
</script>

<style scoped lang="scss">
dialog::backdrop {
  backdrop-filter: blur(3px);
}

dialog[open] {
  animation: modalFadeIn 0.5s ease normal;
}

@keyframes modalFadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

.delete-modal {
  margin: auto;
  padding: 10px;
  background-color: #aaa;
  border: 2px solid;
  border-radius: 5px;

  &__container {
    width: 300px;

    &__buttons {
      margin: 20px 0 10px 0;
      display: flex;
      flex-flow: row nowrap;
      justify-content: space-between;
    }
  }

  button {
    border: none;
    padding: 10px 20px;
    text-align: center;
    text-decoration: none;
    display: inline-block;
    border-radius: 5px;
    cursor: pointer;
  }
}

</style>
