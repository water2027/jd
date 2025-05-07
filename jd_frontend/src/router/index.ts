import { createRouter, createWebHistory } from 'vue-router'
import apiBus from '@/utils/apiBus'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'default',
      component: () => import('@/layouts/DefaultLayout.vue'),
      children: [
        {
          path: '/upload',
          component: () => import('@/views/UploadView.vue'),
          meta: { requiresAuth: true }, // 需要登录才能访问
        },
        {
          path: '/post', // 添加 /post 路由规则
          component: () => import('@/views/PostPage.vue'),
        },
      ],
    },
    {
      path: '/auth/',
      name: 'auth',
      component: () => import('@/layouts/AuthLayout.vue'),
      children: [
        {
          path: 'login',
          name: 'login',
          component: () => import('@/views/auth/LoginView.vue'),
        },
        {
          path: 'register',
          name: 'register',
          component: () => import('@/views/auth/RegisterView.vue'),
        },
      ],
    },
  ],
})

// 路由守卫，验证用户是否登录
router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth && !localStorage.getItem('token')) {
    next('/login')
  } else {
    next()
  }
})

const logout = () => {
  router.push('/auth/login')
}

apiBus.on('API:LOGIN', (req) => {
  const { stop } = req
  if (stop) return
  router.push('/')
})

apiBus.on('API:UN_AUTH', () => {
  logout()
})

apiBus.on('API:LOGOUT', () => {
  logout()
})

export default router
