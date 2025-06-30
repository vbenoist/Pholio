<template>
  <div v-if="fetched && items" class="container" :ref="containerRefName">
    <TileGroup
      v-for="(group, idx) in items"
      :key="`tilegroup-${idx}`"
      :title="formatGroupDate(group['group-by'])"
      :items="group.results"
    />
  </div>
  <div v-else-if="fetched">Aucune photo n'est disponible.</div>
  <div v-else>Loading...</div>
</template>

<script setup lang="ts">
import { computed,  } from 'vue'
import { storeToRefs } from 'pinia'
import TileGroup from '@/components/Content/TileGroup.vue'
import { usePerDateStore } from '@/stores/perDateContent'
import scrollHook from '@/composables/scrollHook'

const perDateContentStore = usePerDateStore()
const { fetched } = storeToRefs(perDateContentStore)
const { containerRefName, registerHook } = scrollHook()

const resolveContent = async () => {
  await perDateContentStore.fetchContent()
}

const formatGroupDate = (date: string): string => {
  const d = new Date(date)
  const formatted =  d.toLocaleString("fr-FR", {
    weekday: "long",
    year: "numeric",
    month: "long",
    day: "numeric",
  })

  return `Le ${formatted}`
}

const items = computed(() => perDateContentStore.getContent())

resolveContent()

registerHook(() => {
  perDateContentStore.fetchNextContent()
})
</script>

<style lang="scss" scoped>
@use '@/assets/colors.scss';

.container {
  width: 90%;
  margin: auto;
  padding-top: 20px;
  background-color: colors.$background-pm-color;
}

@media (min-width: 1024px) {
}
</style>
