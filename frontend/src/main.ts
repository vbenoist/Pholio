import './assets/main.scss'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import { OhVueIcon, addIcons } from "oh-vue-icons"
import {
  CoFullscreen,
  IoTrashBinSharp
} from "oh-vue-icons/icons"

import App from './App.vue'
import router from './router'
import Axios from './plugins/axios'
import ApiResolver from './plugins/apiResolver'

const app = createApp(App)
const pinia = createPinia()

app.use(router)
app.use(Axios)
app.use(ApiResolver)
app.use(pinia)

addIcons(CoFullscreen, IoTrashBinSharp)
app.component("v-icon", OhVueIcon)

app.mount('#app')
