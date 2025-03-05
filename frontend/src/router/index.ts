import { createRouter, createWebHistory } from 'vue-router'
import Recently from '@/views/Recently.vue'
import WelcomeItem from '@/components/WelcomeItem.vue'
import TheWelcome from '@/components/TheWelcome.vue'
import HelloWorld from '@/components/HelloWorld.vue'

export type RoutesNames = 'RECENT' | 'DATE' | 'LOCATION' | 'PINED'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    // {
    //   path: '/',
    //   name: 'home',
    //   component: HomeView,
    // },
    // {
    //   path: '/about',
    //   name: 'about',
    //   // route level code-splitting
    //   // this generates a separate chunk (About.[hash].js) for this route
    //   // which is lazy-loaded when the route is visited.
    //   component: () => import('../views/AboutView.vue'),
    // },
    {
      path: '/',
      name: 'RECENT',
      component: Recently,
    },
    {
      path: '/per-date',
      name: 'DATE',
      component: HelloWorld,
    },
    {
      path: '/per-location',
      name: 'LOCATION',
      component: WelcomeItem,
    },
    {
      path: '/pin',
      name: 'PINED',
      component: TheWelcome,
    },
    {
      path: '/admin/add',
      name: 'ADMINADD',
      component: () => import('@/views/Admin/AddImages.vue'),
    },
  ],
})

export default router
