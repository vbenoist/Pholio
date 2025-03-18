<template>
  <div class="container">
    <div class="container__content">
      <span @click="routerGo('RECENT')">Récents</span>
      <span @click="routerGo('DATE')">Par Date</span>
      <span @click="routerGo('LOCATION')">Par Lieu</span>
      <span @click="routerGo('PINED')">Mis en avant</span>
    </div>

    <hr class="container__delimiter" />
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import type { RoutesNames } from '@/router/index'

const router = useRouter()

const routerGo = (route: RoutesNames) => {
  router.push({ name: route })
}
</script>

<style lang="scss" scoped>
@use '@/assets/colors.scss';

.container {
  display: flex;
  flex-flow: column;
  justify-content: center;

  &__content {
    display: flex;
    flex-flow: row;
    justify-content: center;

    span {
      color: #fff;

      &:hover {
        cursor: pointer;
      }

      &::before,
      &::after {
        content: '';
        margin: 0 15px;
      }
    }
  }

  &__delimiter {
    border: 1px solid colors.$foreground-pm-color;
    margin: 16px auto 0;
    width: 500px;
  }
}
</style>
