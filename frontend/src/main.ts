import './assets/main.scss'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import { OhVueIcon, addIcons } from 'oh-vue-icons'
import {
  BiCalendarDateFill,
  BiCloudArrowUpFill,
  BiCloudCheckFill,
  BiCloudSlashFill,
  CoFullscreen,
  IoClose,
  IoLocationSharp,
  IoTrashBinSharp,
  MdErrorOutlined,
  MdPending,
  RiLoader4Line,
} from 'oh-vue-icons/icons'

import App from './App.vue'
import router from './router'
import Axios from './plugins/axios'
import AuthResolver from './plugins/auth'
import ApiResolver from './plugins/apiResolver'
import ApiPathBuild from './plugins/apiPathBuilder'

const app = createApp(App)
const pinia = createPinia()

app.use(router)
app.use(Axios)
app.use(AuthResolver)
app.use(ApiResolver)
app.use(ApiPathBuild)
app.use(pinia)

addIcons(
  BiCalendarDateFill,
  BiCloudArrowUpFill,
  BiCloudCheckFill,
  BiCloudSlashFill,
  CoFullscreen,
  IoClose,
  IoLocationSharp,
  IoTrashBinSharp,
  MdErrorOutlined,
  MdPending,
  RiLoader4Line,
)
app.component('v-icon', OhVueIcon)

app.mount('#app')
