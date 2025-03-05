<template>
	<div
    :data-active="active"
    @dragenter.prevent="setActive"
    @dragover.prevent="setActive"
    @dragleave.prevent="setInactive"
    @drop.prevent="onDrop"
  >
		<slot :dropZoneActive="active"></slot>
	</div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import fileManager from '@/composables/fileManager'

const emit = defineEmits(['files-dropped'])
const active = ref(false)
const { convertFileListArray } = fileManager()

let inActiveTimeout: number | undefined = undefined

// setActive and setInactive use timeouts, so that when you drag an item over a child element,
// the dragleave event that is fired won't cause a flicker. A few ms should be plenty of
// time to wait for the next dragenter event to clear the timeout and set it back to active.
const setActive = () => {
	active.value = true
	clearTimeout(inActiveTimeout)
}

const setInactive = () => {
	inActiveTimeout = setTimeout(() => {
		active.value = false
	}, 50)
}

const onDrop = (e: DragEvent) => {
	setInactive()
  const files = e.dataTransfer?.files
  if(!files || files.length === 0) return

  const filesArr = convertFileListArray(files)

  const filtered = filesArr.filter(f =>
    f.type === "image/png" || f.type === "image/jpeg"
  )
  if(filtered.length === 0) return

  emit('files-dropped', filtered)
}

const preventDefaults = (e: Event) => {
	e.preventDefault()
}

const events = ['dragenter', 'dragover', 'dragleave', 'drop']

onMounted(() => {
	events.forEach((eventName) => {
		document.body.addEventListener(eventName, preventDefaults)
	})
})

onUnmounted(() => {
	events.forEach((eventName) => {
		document.body.removeEventListener(eventName, preventDefaults)
	})
})
</script>
