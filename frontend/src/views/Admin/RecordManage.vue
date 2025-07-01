<template>
  <div :ref="containerRefName">
    <span v-if="fetched">
      <ManagePublications v-model="apiDetailedRecord" />
    </span>
    <span v-else>Loading...</span>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { storeToRefs } from 'pinia'
import scrollHook from '@/composables/scrollHook'
import ManagePublications from '@/components/Admin/Manage/ManagePublications.vue'
import { useDetailedStore } from '@/stores/detailedContent'
const detailedApiContentStore = useDetailedStore()
const { fetched } = storeToRefs(detailedApiContentStore)
const { containerRefName, registerHook } = scrollHook()

const apiDetailedRecord = computed(() => detailedApiContentStore.getContent())

const resolveContent = async () => {
  await detailedApiContentStore.fetchContent()
}

registerHook(() => {
  detailedApiContentStore.fetchNextContent()
})

resolveContent()

</script>
