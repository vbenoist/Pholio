import { createRouter, createWebHistory } from 'vue-router'

export type RoutesNames = 'RECENT' | 'DATE' | 'LOCATION' | 'PINED'

export const routes = [
  {
    path: '/',
    name: 'RECENT',
    component: () => import('@/views/Recently.vue'),
    meta: {
      layout: 'guest'
    }
  },
  {
    path: '/per-date',
    name: 'DATE',
    component: () => import('@/views/Recently.vue'),
    meta: {
      layout: 'guest'
    }
  },
  {
    path: '/per-location',
    name: 'LOCATION',
    component: () => import('@/views/Recently.vue'),
    meta: {
      layout: 'guest'
    }
  },
  {
    path: '/pin',
    name: 'PINED',
    component: () => import('@/views/Recently.vue'),
    meta: {
      layout: 'guest'
    }
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

export default router
