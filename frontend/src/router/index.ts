import { createRouter, createWebHistory } from 'vue-router'
import Recently from '@/views/Recently.vue'
import WelcomeItem from '@/components/WelcomeItem.vue'
import TheWelcome from '@/components/TheWelcome.vue'
import HelloWorld from '@/components/HelloWorld.vue'

export type RoutesNames = 'RECENT' | 'DATE' | 'LOCATION' | 'PINED'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'RECENT',
      component: Recently,
      meta: {
        layout: 'guest'
      }
    },
    {
      path: '/per-date',
      name: 'DATE',
      component: HelloWorld,
      meta: {
        layout: 'guest'
      }
    },
    {
      path: '/per-location',
      name: 'LOCATION',
      component: WelcomeItem,
      meta: {
        layout: 'guest'
      }
    },
    {
      path: '/pin',
      name: 'PINED',
      component: TheWelcome,
      meta: {
        layout: 'guest'
      }
    },
    {
      path: '/admin/add',
      name: 'ADMINADD',
      component: () => import('@/views/Admin/AddImages.vue'),
      meta: {
        layout: 'admin'
      }
    },
  ],
})

export default router
