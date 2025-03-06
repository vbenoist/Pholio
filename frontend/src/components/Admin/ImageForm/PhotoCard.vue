<template>
  <div v-if="draftRecord" class="photo-card">
      <div class="photo-card__form">
        <span>
          Date: {{ draftRecord.date }}
        </span>
        <span>
          Description: {{ draftRecord.description }}
        </span>
        <span>
          Lieu: <input type="text" name="Lieu" v-model="draftRecord.location" />
        </span>
      </div>
  </div>
</template>

<script setup lang="ts">
import { defineModel } from 'vue'
import type { DraftRecord } from '@/models/record'

const draftRecord = defineModel<DraftRecord | null>({ default: null })



</script>

<style scoped lang="scss">

.photo-card {
  display: flex;
  height: 125px;
  margin: 20px 0;

  &__img {
    max-width: 250px;
    max-height: 125px;
  }

  &__form {
    display: flex;
    flex-flow: column nowrap;
  }
}

</style>
