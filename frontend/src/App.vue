<template>
  <header v-if="!isAdminRoute">
      <BannerHeader />
      <BannerMenu class="banner-menu" />
  </header>

  <RouterView />
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import BannerHeader from '@/components/Common/BannerHeader.vue'
import BannerMenu from '@/components/Common/BannerMenu.vue'

const route = useRoute()

const isAdminRoute = computed(() => {
  return route.fullPath.includes('admin')
})

</script>

<style scoped>
.banner-menu {
  margin: 40px 0;
}

@media (min-width: 1024px) {
}
</style>
