import { expect, test } from 'vitest'
import { flushPromises, mount } from "@vue/test-utils"
import { createRouter, createWebHistory } from 'vue-router'
import BannerMenu from '@/components/Common/BannerMenu.vue'
import type { RoutesNames } from '@/router'
import { routes } from "@/router"

const router = createRouter({
  history: createWebHistory(),
  routes
})

test('checking menu list', async () => {
  expect(BannerMenu).toBeTruthy()

  const wrapper = mount(BannerMenu, {})

  expect(wrapper.text()).toContain("Récents")
  expect(wrapper.text()).toContain("Par Date")
  expect(wrapper.text()).toContain("Par Lieu")
  expect(wrapper.text()).toContain("Mis en avant")
})


test('checking menu nav actions', async () => {
  router.push('/')
  await router.isReady()

  const wrapper = mount(BannerMenu, {
    global: {
      plugins: [router]
    }
  })

  const menus = wrapper.findAll('span')
  expect(menus.length).toEqual(4)

  const navOrder: Array<RoutesNames> = ['RECENT', 'DATE', 'LOCATION', 'PINED']

  for(let i=0; i<menus.length; i++) {
    await menus[i].trigger('click')
    await flushPromises()
    expect(wrapper.vm.$route.name).toEqual(navOrder[i])
  }
})
