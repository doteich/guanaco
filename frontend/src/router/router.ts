import { createRouter, createWebHashHistory } from 'vue-router'
import clientView from "../views/clientView.vue"
import browseView from "../views/browseView.vue"
import monitorView from "../views/monitorView.vue"

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

    {
      path: '/monitor',
      name: 'monitor',
      component: monitorView
    },


  ]
})

export default router

