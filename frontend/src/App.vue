<template>
  <component :is="layouts[routeLayout]">
    <slot>
      <RouterView />
    </slot>
  </component>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import LayoutAdmin from '@/layouts/LayoutAdmin.vue'
import LayoutGuest from '@/layouts/LayoutGuest.vue'

const route = useRoute()

const layouts = {
  LayoutAdmin,
  LayoutGuest
}

const routeLayout = computed<keyof typeof layouts>(() => {
  if(!route.meta.layout) return 'LayoutGuest'

  const meta = String(route.meta.layout)
  const formatted = meta[0].toUpperCase() + meta.slice(1)
  return `Layout${formatted}`
})
</script>
