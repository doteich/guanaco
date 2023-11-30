import { createRouter, createWebHashHistory } from 'vue-router'
import clientView from "../views/clientView.vue"

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: clientView
    },

  ]
})

export default router

