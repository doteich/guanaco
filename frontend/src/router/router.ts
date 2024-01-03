import { createRouter, createWebHashHistory } from 'vue-router'
import clientView from "../views/clientView.vue"
import browseView from "../views/browseView.vue"

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: clientView
    },
    {
      path: '/browse',
      name: 'browse',
      component: browseView
    },


  ]
})

export default router

