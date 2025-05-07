<template>
  <div class="post-page">
    <!-- 导航栏 -->
    <div class="nav-bar">
      <ul>
        <li v-for="post in posts" :key="post.id" @click="selectPost(post.id)">
          {{ post.title }}
        </li>
      </ul>
    </div>
    <!-- 帖子内容显示区域 -->
    <div class="post-content">
      <div v-if="selectedPost">
        <h2>{{ selectedPost.title }}</h2>
        <p>{{ selectedPost.content }}</p>
      </div>
      <div v-else>
        <p>请选择一个帖子查看详情</p>
      </div>
    </div>
    <!-- 右上角登录提示 -->
    <div class="login-tip">
      <router-link to="/auth/login">请登录</router-link>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';

// 定义帖子类型
type Post = {
  id: number;
  title: string;
  content: string;
};

// 模拟帖子数据
const posts = ref<Post[]>([
  { id: 1, title: '保护鸟类的重要性', content: '鸟类是生态系统的重要组成部分...' },
  { id: 2, title: '常见鸟类的保护方法', content: '对于常见鸟类，我们可以采取以下保护措施...' },
  { id: 3, title: '鸟类栖息地的保护', content: '保护鸟类栖息地是保护鸟类的关键...' }
]);

// 当前选中的帖子，明确类型为 Post 或 null
const selectedPost = ref<Post | null>(null);

// 选择帖子的方法，明确 id 类型为 number
const selectPost = (id: number) => {
  const post = posts.value.find(post => post.id === id);
  // 处理 post 可能为 undefined 的情况
  selectedPost.value = post || null;
};
</script>

<style scoped>
.post-page {
  display: flex;
  height: 100vh;
  background-color: #f4f4f4; /* 添加背景颜色 */
}

.nav-bar {
  width: 16.66%; /* 六分之一宽度 */
  border-right: 1px solid #ccc;
  overflow-y: auto;
  padding: 10px;
  background-color: #fff; /* 添加背景颜色 */
  box-shadow: 2px 0 5px rgba(0, 0, 0, 0.1); /* 添加阴影 */
}

.nav-bar ul {
  list-style-type: none;
  padding: 0;
}

.nav-bar li {
  cursor: pointer;
  margin-bottom: 10px;
  padding: 10px;
  border-radius: 5px;
  transition: background-color 0.3s; /* 添加过渡效果 */
}

.nav-bar li:hover {
  background-color: #e0e0e0; /* 鼠标悬停效果 */
}

.post-content {
  width: 83.33%; /* 六分之五宽度 */
  padding: 20px;
  overflow-y: auto;
  background-color: #fff; /* 添加背景颜色 */
}

.login-tip {
  position: absolute;
  top: 10px;
  right: 10px;
  background-color: #fff; /* 添加背景颜色 */
  padding: 5px 10px;
  border-radius: 5px;
  box-shadow: 0 0 5px rgba(0, 0, 0, 0.1); /* 添加阴影 */
}

.login-tip a {
  color: #007bff;
  text-decoration: none;
}

.login-tip a:hover {
  text-decoration: underline;
}
</style>