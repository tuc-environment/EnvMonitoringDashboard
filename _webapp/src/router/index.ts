import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import NewHomeViewVue from '@/views/NewHomeView.vue'
import LoginView from '@/views/LoginView.vue'
import RegisterView from '@/views/RegisterView.vue'
import httpclient from '@/httpclient'

import AdminView from '@/views/admin/AdminView.vue'
import CrossStationView from '@/views/CrossStationView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/new',
      name: 'new-homw',
      component: NewHomeViewVue
    },
    {
      path: '/cross-station',
      name: 'cross-station-tab',
      component: CrossStationView
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView
    },
    {
      path: '/register',
      name: 'register',
      component: RegisterView
    },
    {
      path: '/admin',
      name: 'admin',
      meta: {
        requiresAuth: true
      },
      component: AdminView
    }
  ]
})

router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth && !httpclient.isAuthorized()) {
    next({ path: '/login' })
  } else {
    next()
  }
})

export default router
