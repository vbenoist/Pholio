import { expect, test } from 'vitest'
import { mount } from "@vue/test-utils"
import TileGroup from '@/components/Content/TileGroup.vue'
import { mockRecords } from "@/mock/static/records"

test('mount & renders title correctly', async () => {
  expect(TileGroup).toBeTruthy()

  const wrapper = mount(TileGroup, {
    props: {
      title: 'Derniers ajouts',
      items: []
    }
  })

  expect(wrapper.text()).toContain("Derniers ajouts")
})

test('renders images', async () => {
  const mockData = mockRecords(10)
  expect(TileGroup).toBeTruthy()

  const wrapper = mount(TileGroup, {
    props: {
      title: 'Derniers ajouts',
      items: mockData
    }
  })

  const imgs = wrapper.findAll('img')

  for(let i=0; i<mockData.length; i++) {
    expect(imgs[i].attributes('alt')).toEqual(mockData[i].description)
  }

  expect(wrapper.findAll('img').length).toEqual(10)
})

// test('orders & keys images', async () => {
//   expect(TileGroup).toBeTruthy()

//   const wrapper = mount(TileGroup, {
//     props: {
//       title: 'Derniers ajouts',
//       items: mockData
//     }
//   })

//   const imgs = wrapper.findAll('img')

//   for(let i=0; i<mockData.length; i++) {
//     console.log("test key", imgs[i].attributes('key'))
//     expect(imgs[i].attributes('key')).toEqual(`img-last-add-${i}`)
//   }
// })
