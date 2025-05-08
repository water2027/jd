<template>
  <div class="container">
    <!-- 左侧导航栏（保留原有功能） -->
    <div class="nav">
      <button @click="router.push('/create-post')" class="new-post-btn">
        发布新帖子
      </button>
      <div v-for="post in posts" :key="post.id" class="post-item">
        {{ post.createTime }}
        <button @click="viewPost(post)" class="view-btn">查看</button>
      </div>
    </div>

    <!-- 右侧内容区（新增登录按钮） -->
    <div class="content">
      <!-- 新增：右上角登录按钮（路径改为 /auth/login） -->
      <div class="login-container">
        <RouterLink to="/auth/login" class="login-btn">登录</RouterLink>
      </div>

      <!-- 保留原有内容显示逻辑 -->
      <div v-if="selectedPost" v-html="selectedPost.content"></div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useRouter, RouterLink } from 'vue-router' // 导入 RouterLink

// 保留原始帖子类型定义（与原代码一致）
interface Post {
  id: number;
  createTime: string;
  content: string;
}

const router = useRouter()
// 保留原始数据声明（与原代码一致）
const posts = ref<Post[]>([])
const selectedPost = ref<Post | null>(null)

// 保留原始数据获取逻辑（与原代码一致）
onMounted(async () => {
  try {
    const res = await axios.get<{ data: Post[] }>('/api/posts')
    posts.value = res.data.data
  } catch (error) {
    console.error('获取帖子列表失败:', error)
  }
})

// 保留原始查看逻辑（与原代码一致）
const viewPost = (post: Post) => {
  selectedPost.value = post
}
</script>

<style scoped>
/* 仅调整登录按钮样式（与页面风格一致） */
.login-container {
  text-align: right;
  margin-bottom: 1rem;
  padding: 0 0.5rem;
}

.login-btn {
  color: #2563eb; /* 与导航栏按钮颜色一致 */
  text-decoration: none;
  font-size: 0.9rem;
  cursor: pointer;
}

/* 以下为原有样式（与之前一致，未修改） */
.container {
  display: flex;
  min-height: 100vh;
}

.nav {
  flex: 1;
  max-width: 200px;
  border-right: 1px solid #e5e7eb;
  padding: 1rem;
  background: #f8f9fa;
}

.new-post-btn {
  width: 100%;
  padding: 0.5rem;
  background: #2563eb;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  margin-bottom: 1rem;
}

.post-item {
  padding: 0.5rem;
  border-radius: 4px;
  margin-bottom: 0.5rem;
  transition: background 0.2s;
}

.post-item:hover {
  background: #e5e7eb;
}

.view-btn {
  margin-left: 1rem;
  padding: 0.2rem 0.5rem;
  background: #64748b;
  color: white;
  border: none;
  border-radius: 3px;
  cursor: pointer;
}

.content {
  flex: 5;
  padding: 2rem;
}
</style>