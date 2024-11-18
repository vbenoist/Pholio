<template>
  <!-- <header>
    <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" />

    <div class="wrapper">
      <HelloWorld msg="You did it!" />

      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/about">About</RouterLink>
      </nav>
    </div>
  </header> -->

  <BannerHeader />
  <BannerMenu class="banner-menu" />

  <RouterView />
</template>

<script setup lang="ts">
// import { RouterLink, RouterView } from 'vue-router'
// import HelloWorld from './components/HelloWorld.vue'

import BannerHeader from '@/components/Common/BannerHeader.vue'
import BannerMenu from '@/components/Common/BannerMenu.vue'
</script>

<style scoped>
.banner-menu {
  margin: 40px 0;
}

@media (min-width: 1024px) {
}
</style>
