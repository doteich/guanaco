import {createApp} from 'vue'
import { createPinia } from 'pinia'
import router from "./router/router"

import PrimeVue from 'primevue/config';
import ToastService from 'primevue/toastservice';

import App from './App.vue'


import './style.css';
import 'primevue/resources/themes/lara-dark-teal/theme.css'
import 'primeicons/primeicons.css'


const pinia = createPinia()

const app = createApp(App)

app.use(pinia)
app.use(PrimeVue)
app.use(router)
app.use(ToastService)


app.mount("#app")
