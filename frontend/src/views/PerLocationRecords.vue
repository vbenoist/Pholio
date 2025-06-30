<template>
  <div v-if="fetched && items.length > 0" class="container" :ref="containerRefName">
    <TileGroup
      v-for="(group, idx) in items"
      :key="`tilegroup-${idx}`"
      :title="group['group-by']"
      :items="group.results"
      overlay-caption
    >
      <template #item-extend="{ item }">
       <span class="item__caption">
          <v-icon
            name="bi-calendar-date-fill"
            scale="0.8"
          />
          {{ formatDate(item.date) }}
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
import { usePerLocationStore } from '@/stores/perLocationContent'
import scrollHook from '@/composables/scrollHook'

const perLocationContentStore = usePerLocationStore()
const { fetched } = storeToRefs(perLocationContentStore)
const { containerRefName, registerHook } = scrollHook()

const resolveContent = async () => {
  await perLocationContentStore.fetchContent()
}

const formatDate = (date: Date): string => {
  return date.toLocaleDateString()
}

const items = computed(() => perLocationContentStore.getContent())

resolveContent()

registerHook(() => {
  perLocationContentStore.fetchNextContent()
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
