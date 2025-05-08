// src/stores/postStore.ts
import { defineStore } from 'pinia';
import axios from 'axios';

// 定义帖子类型
type Post = {
  id: number;
  title: string;
  content: string;
};

// 定义帖子Store
export const usePostStore = defineStore('post', {
  state: () => ({
    posts: [] as Post[], // 存储帖子列表
  }),
  actions: {
    // 获取帖子列表（从API加载）
    async fetchPosts() {
      const res = await axios.get('/api/posts'); // 假设后端提供获取帖子列表的接口
      this.posts = res.data;
    },
    // 添加新帖子（发布成功后调用）
    async addPost(newPost: Post) {
      this.posts.unshift(newPost); // 将新帖子添加到列表头部（或根据需求调整顺序）
    },
  },
});