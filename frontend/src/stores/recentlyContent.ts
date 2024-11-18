import type { Ref } from 'vue'
import { ref } from 'vue'
import { defineStore } from 'pinia'

export type RecentlyContent = {
  lastly: Array<Item>
  lately: Array<Item>
}

export type Item = {
  nativImgSrc: string
  midImgSrc: string
  thumbImgSrc: string
  description: string
  location: string
  date: Date
}

export const useRecentlyContentStore = defineStore('recentlyContent', () => {
  let content: RecentlyContent | null = null
  const fetched: Ref<boolean> = ref(false)

  async function initContent() {
    // API Call
    // MOCK instead

    content = {
      lastly: [
        {
          nativImgSrc: '/src/assets/img/IMG_2813.jpg',
          midImgSrc: '/src/assets/img/IMG_2813.jpg',
          thumbImgSrc: '/src/assets/img/IMG_2813.jpg',
          description: "Lac de Crouserocs au petit matin, à la fin du mois d'octobre",
          location: 'Lac de Crouserocs - Macif des Cerces',
          date: new Date('2024-10-30'),
        },
        {
          nativImgSrc: '/src/assets/img/IMG_2813.jpg',
          midImgSrc: '/src/assets/img/IMG_2813.jpg',
          thumbImgSrc: '/src/assets/img/IMG_2813.jpg',
          description: "Lac de Crouserocs au petit matin, à la fin du mois d'octobre",
          location: 'Lac de Crouserocs - Macif des Cerces',
          date: new Date('2024-10-30'),
        },
      ],
      lately: [
        {
          nativImgSrc: '/src/assets/img/IMG_2813.jpg',
          midImgSrc: '/src/assets/img/IMG_2813.jpg',
          thumbImgSrc: '/src/assets/img/IMG_2813.jpg',
          description: "Lac de Crouserocs au petit matin, à la fin du mois d'octobre",
          location: 'Lac de Crouserocs - Macif des Cerces',
          date: new Date('2024-10-30'),
        },
      ],
    }

    fetched.value = true
  }

  async function fetchContent() {
    if (content === null && !fetched.value) {
      await initContent()
    }

    return content
  }

  function getContent() {
    return content
  }

  function hasFetched() {
    return fetched
  }

  return { content, hasFetched, fetchContent, getContent, initContent }
})
