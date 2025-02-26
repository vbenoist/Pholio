<template>
  <div v-if="hasFetched && items" class="container">
    <TileGroup title="Derniers ajouts" :items="items.lastly" />
    <TileGroup title="Plus anciennement" :items="items.lately" />
  </div>
  <div v-else>Loading...</div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import TileGroup from '@/components/Content/TileGroup.vue'
import { useRecentlyContentStore } from '@/stores/recentlyContent'

const recentlyContentStore = useRecentlyContentStore()
const hasFetched = recentlyContentStore.hasFetched()

const resolveContent = async () => {
  await recentlyContentStore.fetchContent()
}

const items = computed(() => recentlyContentStore.getContent())

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
