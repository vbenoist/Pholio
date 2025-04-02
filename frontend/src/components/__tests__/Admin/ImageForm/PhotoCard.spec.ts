import { ref } from 'vue'
import { expect, test, vi } from 'vitest'
import { mount } from '@vue/test-utils'
import axios from 'axios'
import PhotoCard from '@/components/Admin/ImageForm/PhotoCard.vue'
import { ApiResolver } from '@/plugins/apiResolver'
import { mockDraftRecords } from '@/mock/static/records'
import type { DraftRecord } from '@/models/record'

global.URL.createObjectURL = vi.fn(() => 'details')
const apiResolver = new ApiResolver(axios)
const draftRecord = ref<DraftRecord>(mockDraftRecords(1)[0])

/* While draftRecord is not set, form should not be mounted */
test('should not mount', async () => {
  const wrapper = mount(PhotoCard, {
    props: {
      draftRecord: null,
    },
    global: {
      provide: {
        $apiResolver: apiResolver,
      },
    },
  })

  expect(wrapper.find('form').exists()).toBe(false)
})

/* With the newly set draftRecord prop, all component should now be mounted */
test('should mount', async () => {
  const wrapper = mount(PhotoCard, {
    props: {
      modelValue: draftRecord.value,
      'onUpdate:modelValue': (e) => wrapper.setProps({ modelValue: e }),
    },
    global: {
      provide: {
        $apiResolver: apiResolver,
      },
    },
  })

  expect(wrapper.find('form').exists()).toBe(true)
})

test('input content tests', async () => {
  const wrapper = mount(PhotoCard, {
    props: {
      modelValue: draftRecord.value,
      'onUpdate:modelValue': (e) => wrapper.setProps({ modelValue: e }),
    },
    global: {
      provide: {
        $apiResolver: apiResolver,
      },
    },
  })

  const inputs = wrapper.findAll('input')
  expect(inputs.length).toEqual(2)

  const inputDateVal = new Date(inputs[0].element.value)
  expect(inputDateVal.toLocaleDateString()).toEqual(draftRecord.value.date.toLocaleDateString())
})

test('input validation tests', async () => {
  const wrapper = mount(PhotoCard, {
    props: {
      modelValue: draftRecord.value,
      'onUpdate:modelValue': (e) => wrapper.setProps({ modelValue: e }),
    },
    global: {
      provide: {
        $apiResolver: apiResolver,
      },
    },
  })

  expect(wrapper.vm.submitIcon).toEqual('bi-cloud-arrow-up-fill')

  wrapper.vm.triggerValidation()
  expect(wrapper.vm.submitIcon).toEqual('md-error-outlined')

  const inputs = wrapper.findAll('input')
  await inputs[1].setValue('Les Rochilles')
  expect(wrapper.vm.submitIcon).toEqual('bi-cloud-arrow-up-fill')
})
