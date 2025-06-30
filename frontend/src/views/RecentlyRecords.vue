<template>
  <div v-if="fetched && items" class="container" :ref="containerRefName">
    <TileGroup
      title="Derniers ajouts"
      :items="items.lastly"
      overlay-caption
    >
      <template #item-extend="{ item }">
        <span class="item__caption">
          <v-icon
            name="io-location-sharp"
            scale="0.8"
          />
          {{ item.location }}
        </span>
      </template>
    </TileGroup>
    <TileGroup
      title="Plus anciennement"
      :items="items.lately"
      overlay-caption
    >
      <template #item-extend="{ item }">
        <span class="item__caption">
          <v-icon
            name="io-location-sharp"
            scale="0.8"
          />
          {{ item.location }}
        </span>
      </template>
    </TileGroup>
  </div>
  <div v-else-if="fetched">Aucune photo n'est disponible.</div>
  <div v-else>Loading...</div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { storeToRefs } from 'pinia'
import TileGroup from '@/components/Content/TileGroup.vue'
import { useRecentlyContentStore } from '@/stores/recentlyContent'
import scrollHook from '@/composables/scrollHook'

const recentlyContentStore = useRecentlyContentStore()
const { fetched } = storeToRefs(recentlyContentStore)
const { containerRefName, registerHook } = scrollHook()

const resolveContent = async () => {
  await recentlyContentStore.fetchContent()
}

const items = computed(() => recentlyContentStore.getContent())

resolveContent()

registerHook(() => {
  recentlyContentStore.fetchNextContent()
})
</script>

<style lang="scss" scoped>
@use '@/assets/colors.scss';

.container {
  width: 90%;
  margin: auto;
  padding-top: 20px;
  background-color: colors.$background-pm-color;

  .item {
    &__caption {
      font-size: 0.9em;
    }
  }
}

@media (min-width: 1024px) {
}
</style>
