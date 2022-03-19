import Vue from 'vue'
import Router from 'vue-router'
import Index from './pages/Index.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  // eslint-disable-next-line no-undef
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'index',
      component: Index
    },
    {
      path: '/announce',
      name: 'announce',
      component: () => import('./pages/Announce.vue')
    },
    {
      path: '/policy',
      name: 'policy',
      component: () => import('./pages/Policy.vue')
    },
    {
      path: '/feedback',
      name: 'feedback',
      component: () => import('./pages/Feedback.vue')
    }
  ]
})