import '@tabler/core'
import '@tabler/core/dist/css/tabler.min.css'
import '@tabler/icons-vue'
import 'bootstrap-icons/font/bootstrap-icons.css'
import './assets/main.css'
import 'vue3-loading-overlay/dist/vue3-loading-overlay.css'
import Vue3Toastify, { type ToastContainerOptions } from 'vue3-toastify'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import VueApexCharts from 'vue3-apexcharts'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(VueApexCharts)
app.use(Vue3Toastify, {
  autoClose: 3000
} as ToastContainerOptions)

app.mount('#app')
