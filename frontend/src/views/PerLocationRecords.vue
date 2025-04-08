<template>
  <div v-if="hasFetched && items" class="container">
    <TileGroup
      v-for="(group, idx) in items"
      :key="`tilegroup-${idx}`"
      :title="group['group-by']"
      :items="group.results"
    />
  </div>
  <div v-else-if="hasFetched">Aucune photo n'est disponible.</div>
  <div v-else>Loading...</div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import TileGroup from '@/components/Content/TileGroup.vue'
import { usePerLocationStore } from '@/stores/perLocationContent'

const perLocationContentStore = usePerLocationStore()
const hasFetched = perLocationContentStore.hasFetched()

const resolveContent = async () => {
  await perLocationContentStore.fetchContent()
}

const items = computed(() => perLocationContentStore.getContent())

resolveContent()
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
