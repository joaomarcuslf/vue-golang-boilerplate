import { mount } from '@vue/test-utils'
import NuxtLogo from '@/pages/index.vue'

describe('IndexPage', () => {
  test('is a Vue instance', () => {
    const wrapper = mount(NuxtLogo)
    expect(wrapper.vm).toBeTruthy()
  })
})
