import { createRouter, createWebHistory } from 'vue-router';
import LoginView from '../views/LoginView.vue';
import UploadView from '../views/UploadView.vue';

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
    path: '/',
    redirect: '/login'
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
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
