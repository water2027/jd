import { createRouter, createWebHistory } from 'vue-router'
import apiBus from '@/utils/apiBus'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'default',
      component: () => import('@/layouts/DefaultLayout.vue'),
    },
    {
      path: '/auth',
      name: 'auth',
      component: () => import('@/layouts/AuthLayout.vue'),
    },
  ],
})

const logout = () => {
  router.push('/auth/login')
}

apiBus.on('API:UN_AUTH', () => {
  logout()
})

apiBus.on('API:LOGOUT', () => {
  logout()
})

apiBus.on('API:LOGIN', (req) => {
  const { stop } = req
  if (stop) return
  router.push('/')
})

export default router
