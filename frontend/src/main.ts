import './assets/main.scss'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

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

app.mount('#app')
