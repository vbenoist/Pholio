<template>
  <div>
    <header class="header">
      <h1 class="header__title">PHOLIO</h1>
      <nav class="header__nav">
        <ul class="header__nav__content">
          <li
            :class="{ active: routeName === 'ADMIN-MANAGE' }"
            @click="goToRoute('ADMIN-MANAGE')"
          >Publications</li>
          <li
            :class="{ active: routeName === 'ADMIN-ADD' }"
            @click="goToRoute('ADMIN-ADD')"
          >Médiathèque</li>
          <li :class="{ active: routeName === 'ADMIN-SETUP' }">Configuration</li>
        </ul>
      </nav>
    </header>

    <div class="container">
      <slot></slot>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import type { RoutesAdminNames } from '@/router'

const router = useRouter()
const route = useRoute()

const goToRoute = (routeName: RoutesAdminNames) => {
  router.push({ name: routeName })
}

const routeName = computed((): string => {
  return route.name ?? ""
})

</script>

<style lang="scss" scoped>
.header {
  width: 100%;
  height: 68px;
  margin-bottom: 30px;

  &__nav {
    width: 600px;
    height: 100%;
    padding-left: 0;
    display: flex;
    margin: auto;

    &__content {
      width: inherit;
      display: flex;
      flex-flow: row nowrap;
      justify-content: space-around;
      align-items: center;
      padding-left: 0;

      > li {
        list-style-type: none;
        display: inline-block;
        position: relative;

        &::after {
          transform-origin: bottom center;
          content: '';
          position: absolute;
          width: 100%;
          transform: scaleX(0);
          height: 1px;
          bottom: 0;
          left: 0;
          background-color: white;
          transition: transform 0.25s ease-out;
        }

        &:hover {
          cursor: pointer;
          color: white;

          &::after {
            transform-origin: bottom center;
            transform: scaleX(1);
          }
        }

        .active {
          color: red;
          &::after {
            transform-origin: bottom center;
            content: '';
            position: absolute;
            width: 100%;
            transform: scaleX(0);
            height: 1px;
            top: 0;
            left: 0;
            background-color: white;
            transition: transform 0.25s ease-out;
          }
        }
      }
    }
  }

  &__title {
    position: absolute;
    margin: 8px 16px;
    font-family: 'JostSemiBold';
    font-size: 30px;
    letter-spacing: 0.5cap;
    color: #ffffff;
  }
}

.active {
  color: white;
}

</style>
