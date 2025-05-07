import { createRouter, createWebHistory } from 'vue-router';
import LoginView from '@/views/LoginView.vue';
import UploadView from '@/views/UploadView.vue';
import PostPage from '../views/PostPage.vue';

const routes = [
  {
    path: '/login',
    component: LoginView
  },
  {
    path: '/upload',
    component: UploadView,
    meta: { requiresAuth: true } // 需要登录才能访问
  },
  {
    path: '/post', // 添加 /post 路由规则
    component: PostPage
  },
  {
    path: '/',
    redirect: '/login'
  }
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'default',
      component: () => import('@/layouts/DefaultLayout.vue'),
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
      ]
    },
    ...routes // 合并前面定义的路由
  ]
});

// 路由守卫，验证用户是否登录
router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth && !localStorage.getItem('token')) {
    next('/login');
  } else {
    next();
  }
});

export default router;