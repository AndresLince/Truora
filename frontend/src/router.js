import Vue from 'vue'
import Router from 'vue-router'
import Keys from './views/Keys.vue'
import KeyCreate from './views/KeyCreate.vue'
import KeyEncryptText from './views/KeyEncryptText.vue'


Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'keys',
      component: Keys
    },
    {
      path: '/keys',
      name: 'keys-route',      
      component: () => import('./views/Keys.vue')
    },
    {
      path: '/keys/create',
      name: 'keys-create',
      component: () => import('./views/KeyCreate.vue')
    },
    {
      path: '/keys/:id/encrypt',
      name: 'keys-encrypt',
      component: () => import('./views/KeyEncryptText.vue')
    },   
    
  ]
})
