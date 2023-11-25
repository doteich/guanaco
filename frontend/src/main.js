import {createApp} from 'vue'
import { createPinia } from 'pinia'
import PrimeVue from 'primevue/config';

import App from './App.vue'

import './style.css';
import 'primevue/resources/themes/lara-dark-purple/theme.css'

const pinia = createPinia()

const app = createApp(App)

app.use(pinia)
app.use(PrimeVue)


app.mount("#app")
