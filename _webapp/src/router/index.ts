import httpclient from '@/http-client'
import HomeView from '@/views/HomeView.vue'
import LoginView from '@/views/LoginView.vue'
import RegisterView from '@/views/RegisterView.vue'
import { createRouter, createWebHistory } from 'vue-router'

import ContactUsView from '@/views/ContactUsView.vue'
import CrossStationView from '@/views/CrossStationView.vue'
import AdminView from '@/views/admin/AdminView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/cross-station',
      name: 'cross-station',
      component: CrossStationView
    },
    {
      path: '/contact-us',
      name: 'contact-us',
      component: ContactUsView
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
