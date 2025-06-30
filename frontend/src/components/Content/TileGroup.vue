<template>
  <div class="container">
    <div class="container__content">
      <div class="container__content__title">
        <span class="container__content__title__text">{{ props.title }}</span>
        <hr class="container__content__title__separator" />
      </div>
      <div class="container__content__list">
        <img
          v-for="(itm, key) in props.items"
          :key="`img-last-add-${key}`"
          class="container__content__list__item"
          :src="apiPathBuilder.buildRecordThumbUrl(itm)"
          :alt="itm.description ?? itm.location"
          loading="lazy"
        />
        ette
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { inject } from 'vue'
import type { ApiGetRecord } from '@/models/api/record'
import { ApiPathBuilder } from '@/plugins/apiPathBuilder'

const apiPathBuilder = inject('$apiPathBuilder') as ApiPathBuilder

const props = defineProps({
  title: String,
  items: Array<ApiGetRecord>,
})
</script>

<style lang="scss" scoped>
@use '@/assets/colors.scss';

.container {
  &__content {
    &__title {
      margin-bottom: 20px;
      width: fit-content;

      &__separator {
        width: 120%;
        border: 1px solid colors.$foreground-pm-color;
      }
    }

    &__list {
      display: flex;
      flex-flow: row wrap;
      justify-content: flex-start;

      &__item {
        max-width: 286px;
        margin: 0px 40px 20px 0px;
      }
    }
  }
}

@media (min-width: 1024px) {
}
</style>
