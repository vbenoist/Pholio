import { onMounted, onUnmounted, ref, useTemplateRef } from 'vue'
import { debounce as ldDebounce } from 'lodash'

interface LoaderHook {
  (): void
}

export default () => {
  const containerRefName = 'scrollhookContainer'
  const elContainerRef = useTemplateRef(containerRefName)
  const registeredHook = ref<LoaderHook>(() => false)
  const hasRegisteredHook = ref<boolean>(false)

  const debouncedHandleScroll = ldDebounce(() => {
    handleScroll()
  }, 100)

  onMounted(() => {
    if(hasRegisteredHook.value) {
      window.addEventListener("scroll", debouncedHandleScroll)
    }
  })

  onUnmounted(() => {
    if(hasRegisteredHook.value) {
      window.removeEventListener("scroll", debouncedHandleScroll)
    }
  })

  const handleScroll = () => {
    if(!elContainerRef.value) return

    const element = elContainerRef.value as HTMLElement
    if (element.getBoundingClientRect().bottom < window.innerHeight) {
      registeredHook.value()
    }
  }

  const registerHook = (hook: LoaderHook) => {
    registeredHook.value = hook
    hasRegisteredHook.value = true
  }

  return { containerRefName, registerHook }
}
