<template>
  <div>
    <header>
      <BannerHeader />
      <BannerMenu class="banner-menu" />
    </header>

    <div class="container">
      <slot></slot>
    </div>
  </div>
</template>

<script setup lang="ts">

import BannerHeader from '@/components/Common/BannerHeader.vue'
import BannerMenu from '@/components/Common/BannerMenu.vue'

</script>

<style>
.banner-menu {
  margin: 40px 0;
}

@media (min-width: 1024px) {
}
</style>
