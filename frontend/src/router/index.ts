import { createRouter, createWebHistory } from 'vue-router'

export type RoutesNames = 'RECENT' | 'DATE' | 'LOCATION' | 'PINED'
export type RoutesAdminNames = 'ADMIN-LOGIN' | 'ADMIN-ADD' | 'ADMIN-MANAGE' | 'ADMIN-SETUP'

export const routes = [
  {
    path: '/',
    name: 'RECENT',
    component: () => import('@/views/RecentlyRecords.vue'),
    meta: {
      layout: 'guest',
    },
  },
  {
    path: '/per-date',
    name: 'DATE',
    component: () => import('@/views/PerDateRecords.vue'),
    meta: {
      layout: 'guest',
    },
  },
  {
    path: '/per-location',
    name: 'LOCATION',
    component: () => import('@/views/PerLocationRecords.vue'),
    meta: {
      layout: 'guest',
    },
  },
  {
    path: '/pin',
    name: 'PINED',
    component: () => import('@/views/RecentlyRecords.vue'),
    meta: {
      layout: 'guest',
    },
  },
  {
    path: '/admin',
    redirect: '/admin/add'
  },
  {
    path: '/admin/login',
    name: 'ADMIN-LOGIN',
    component: () => import('@/views/Admin/LogIn.vue'),
    meta: {
      layout: 'default',
    },
  },
  {
    path: '/admin/add',
    name: 'ADMIN-ADD',
    component: () => import('@/views/Admin/AddImages.vue'),
    meta: {
      layout: 'admin',
      auth: true,
    },
  },
  {
    path: '/admin/manage',
    name: 'ADMIN-MANAGE',
    component: () => import('@/views/Admin/RecordManage.vue'),
    meta: {
      layout: 'admin',
      auth: true,
    },
  },  {
    path: '/admin/setup',
    name: 'ADMIN-SETUP',
    component: () => import('@/views/Admin/RecordManage.vue'),
    meta: {
      layout: 'admin',
      auth: true,
    },
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

export default router
