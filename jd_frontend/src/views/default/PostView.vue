<template>
  <div class="root">
    <!-- 左侧导航栏（保留原有功能） -->
    <div class="nav">
      <button @click="router.push('/create')" class="new-post-btn">发布新帖子</button>
      <div v-for="post in posts" :key="post.id" class="post-item">
        {{ post.title }}
        <button @click="viewPost(post.id)" class="view-btn">查看</button>
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
import { useRouter, RouterLink } from 'vue-router' // 导入 RouterLink
import { getPostsList, getMaxId, type Post, getPost } from '@/api/post/get'

const router = useRouter()
const posts = ref<Post[]>([])
const selectedPost = ref<Post | null>(null)

onMounted(async () => {
  try {
    const res = await getPostsList({ limit: 10, id: 10 })
    posts.value = res
  } catch (error) {
    console.error('获取帖子列表失败:', error)
  }
})

const viewPost = async (id: number) => {
  try {
    const res = await getPost({ id })
    if (res) {
      selectedPost.value = res
    } else {
      console.error('获取帖子内容失败')
    }
  } catch (error) {
    console.error('获取帖子内容失败:', error)
  }
}
</script>

<style scoped lang="scss">
.root {
  display: flex;
  width: 100%;
  height: 100%;

  .nav {
    flex: 1;
    border-right: 1px solid #e5e7eb;
    padding: 1rem;
    max-width: 200px;
    background: #f8f9fa;

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
  }

  .content {
    flex: 6;
    padding: 1rem;

    overflow-x: auto;
    overflow-y: auto;

    .login-container {
      text-align: right;
      margin-bottom: 1rem;
      padding: 0 0.5rem;
      .login-btn {
        color: #2563eb; /* 与导航栏按钮颜色一致 */
        text-decoration: none;
        font-size: 0.9rem;
        cursor: pointer;
      }
    }
  }
}
</style>
