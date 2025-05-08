<template>
  <div>
    <FormContainer
      class="w-full sm:w-1/2 mt-5"
      @submit-form="loginAction"
      :form-data="loginForm"
      :disabled="!correct"
    >
      <div class="flex flex-row justify-between">
        <div class="mx-a">
          <input type="checkbox" ref="rememberMe">
            <label for="rememberMe">记住我</label>
            <span class="text-sm text-gray-500">下次自动登录</span>
          </input>
        </div>
        <RouterLink class="mx-a text-purple-800" to="/auth/register">
          注册个账号先
        </RouterLink>
      </div>
    </FormContainer>
    <!-- 回退到主页面的按钮 -->
    <RouterLink to="/" class="back-to-home-button">返回主页面</RouterLink>
  </div>
</template>

<script lang="ts" setup>
import { ref, useTemplateRef } from 'vue';
import { RouterLink } from 'vue-router';
import FormContainer from '@/components/FormContainer.vue';
import type { CustomFormData } from '@/composables/FormExam';
import { useFormExam } from '@/composables/FormExam';
import { showMsg } from '@/components/MessageBox.tsx';
import { login } from '@/api/user/login';
import apiBus from '@/utils/apiBus';

const rememberMe = useTemplateRef('rememberMe');

const loginForm = ref<CustomFormData[]>([
  {
    id: 'telephone',
    value: '',
    label: '手机号',
  },
  {
    id: 'password',
    value: '',
    label: '密码',
    type: 'password',
    autocomplete: 'current-password',
  },
]);

const correct = useFormExam(loginForm);

const loginAction = async () => {
  const telephone = loginForm.value[0].value;
  const password = loginForm.value[1].value;

  try {
    const resp = await login({ telephone, password });
  
    apiBus.emit('API:LOGIN', resp);
  
    if(rememberMe.value?.checked) {
      localStorage.setItem('telephone', telephone);
      localStorage.setItem('password', password);
    }
  } catch (error) {
    console.error('Login failed:', error);
    showMsg(error as string);
    return;
  }
};
</script>

<style scoped>
.back-to-home-button {
  display: block;
  margin-top: 20px;
  color: #007bff;
  text-decoration: none;
}

.back-to-home-button:hover {
  text-decoration: underline;
}
</style>