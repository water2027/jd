import { defineStore } from 'pinia';

export const useUserStore = defineStore('user', {
  state: () => ({
    isLoggedIn: false,
    userInfo: null
  }),
  actions: {
    login(token: string) {
      localStorage.setItem('token', token);
      this.isLoggedIn = true;
      // 可以在这里设置用户信息
    },
    logout() {
      localStorage.removeItem('token');
      this.isLoggedIn = false;
      this.userInfo = null;
    }
  }
});