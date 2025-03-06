<template>
	<div class="main-drop">
		<DropZone
      class="main-drop__area"
      @files-dropped="addFiles"
      @click="triggerInput"
      #default="{ dropZoneActive }"
    >
			<label class="main-drop__area__label" for="photo-drop-input">
				<span v-if="dropZoneActive">
					<span>Glissez-déposez vos photos ici</span>
					<span class="smaller">pour les ajouter</span>
				</span>
				<span v-else class="main-drop__area__label__container">
					<span>Glissez-déposez vos photos ici</span>
					<span class="smaller">
						ou <strong><em>cliquer ici</em></strong> pour choisir vos photos
					</span>
				</span>
			</label>
      <input
        class="main-drop__area__input"
        type="file"
        id="photo-drop-input"
        ref="photo-drop-input"
        multiple
        accept="image/png, image/jpeg"
        @click.stop
        @change="onInputChange"
      />
		</DropZone>
	</div>
</template>

<script setup lang="ts">
import { useTemplateRef } from 'vue'
import DropZone from '@/components/Admin/ImagePicker/DropZone.vue'
import fileManager from '@/composables/fileManager'
import type { UploadableFile } from '@/models/uploadableFile'

const filesModel = defineModel<Array<UploadableFile>>()
const photoInput = useTemplateRef('photo-drop-input')
const {
  addMergeFiles,
  convertFileListArray
} = fileManager()

const onInputChange = (e: Event) => {
  if (!e.target || !(e.target instanceof HTMLInputElement)) return
  if(!e.target.files) return

  filesModel.value = addMergeFiles(convertFileListArray(e.target.files))
  console.log(filesModel.value)
	e.target.value = ''
}

const addFiles = (e: Array<File>) => {
  filesModel.value = addMergeFiles(e)
  console.log(filesModel.value)
}

const triggerInput = () => {
  if(!photoInput.value) return
  (photoInput.value as HTMLElement).click()
}
</script>

<style lang="scss">
.main-drop {
  width: 100%;
  height: 440px;

  &__area {
    background-image: url("data:image/svg+xml,%3csvg width='100%25' height='100%25' xmlns='http://www.w3.org/2000/svg'%3e%3crect width='100%25' height='100%25' fill='none' rx='18' ry='18' stroke='red' stroke-width='4' stroke-dasharray='6%2c 20' stroke-dashoffset='0' stroke-linecap='square'/%3e%3c/svg%3e");
    border-radius: 18px;
    flex-flow: column nowrap;
    display: flex;
    height: 400px;
    width: 400px;
    margin: auto;
    justify-content: center;
    cursor: pointer;

    &__label {
      text-align: center;

      &__container {
        cursor: pointer;
        display: flex;
        flex-flow: column nowrap;
      }
    }

    &__input {
      visibility: hidden;
      position: absolute;
    }
  }
}
</style>
