import { createRouter, createWebHistory } from 'vue-router'
import CreatePost from '@/views/CreatePost.vue' // 发布页面组件
import PostList from '@/views/Postpage.vue'     // 帖子列表页面（含左侧导航）

const routes = [
  { path: '/', name: 'PostList', component: PostList },
  { path: '/create', name: 'CreatePost', component: CreatePost } // 发布页面路由
]

const router = createRouter({ history: createWebHistory(), routes })
export default router