<template>
  <div class="container">
    <div class="container__content">
      <div class="container__content__title">
        <span class="container__content__title__text">{{ props.title }}</span>
        <hr class="container__content__title__separator" />
      </div>
      <div class="container__content__list">
        <div
          v-for="(itm, key) in props.items"
          :key="`img-last-add-${key}`"
          class="container__content__list__item"
        >
          <img
            class="container__content__list__item--img"
            :src="apiPathBuilder.buildRecordThumbUrl(itm)"
            :alt="itm.description ?? itm.location"
            loading="lazy"
          />
          <div
            :class="{
              'container__content__list__item--txt--normal': !props.overlayCaption,
              'container__content__list__item--txt--overlay': props.overlayCaption
            }"
          >
            <slot name="item-extend" :item="itm"></slot>
          </div>
        </div>
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
  overlayCaption: { type: Boolean, default: false },
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
        display: flex;
        flex-flow: column wrap;

        &--img {
          max-width: 286px;
          margin-right: 40px;
        }

        &--txt {
          font-size: 0.9em;

          &--normal {
            margin-bottom: 10px;
          }
          &--overlay {
            width: fit-content;
            padding: 2px 4px;
            margin: -20px 0 10px 0px;
            backdrop-filter: blur(8px);
          }
        }
      }
    }
  }
}

@media (min-width: 1024px) {
}
</style>
