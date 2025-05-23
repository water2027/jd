<script lang="ts" setup>
import { computed, ref, useTemplateRef } from 'vue'

import FormContainer from '@/components/FormContainer.vue'
import type { CustomFormData } from '@/composables/FormExam'
import { useFormExam } from '@/composables/FormExam'
import { showMsg } from '@/components/MessageBox.tsx'

import { register } from '@/api/user/register'
import { sendCode } from '@/api/user/send_code'

import apiBus from '@/utils/apiBus'

const rememberMe = useTemplateRef('rememberMe')

const passwordReg = /^\S*(?=\S{6,})(?=\S*\d)(?=\S*[A-Z])(?=\S*[a-z])(?=\S*[!@#$%^&*? ])\S*$/

const registerForm = ref<CustomFormData[]>([
  {
    id: 'username',
    value: '',
    label: '用户名',
    autocomplete: 'username',
  },
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
    reg: passwordReg,
    autocomplete: 'new-password',
  },
  {
    id: 'password2',
    value: '',
    label: '确认密码',
    reg:passwordReg,
    type: 'password',
  },
  {
    id: 'v_code',
    value: '',
    label: '验证码',
  },
])

const telephoneIsCorrect = computed(() => {
  const telephone = registerForm.value[1].value
  const reg = /^(?:(?:\+|00)86)?1\d{10}$/
  return reg.test(telephone)
})

const correct = useFormExam(registerForm)

const passwordIsCorrect = computed(() => {
  const password = registerForm.value[2].value
  const password2 = registerForm.value[3].value
  return password === password2 && passwordReg.test(password)
})

const sendCodeAction = async () => {
  const telephone = registerForm.value[1].value
  if (!telephoneIsCorrect.value) {
    return
  }
  await sendCode({ telephone })
}

const registerAction = async () => {
  if (!correct.value) {
    return
  }

  const name = registerForm.value[0].value
  const telephone = registerForm.value[1].value
  const password = registerForm.value[2].value
  const password2 = registerForm.value[3].value
  const vcode = registerForm.value[4].value

  try {    
    const resp = await register({ name, telephone, password, password2, vcode })
  
    if(rememberMe.value?.checked) {
      localStorage.setItem('telephone', telephone)
      localStorage.setItem('password', password)
    }
    console.log('register')
    apiBus.emit('API:LOGIN',resp)
  } catch (error) {
    console.error('Register failed:', error)
    showMsg(error as string)
    return
  }
}
</script>

<template>
  <FormContainer
    class="w-full sm:w-1/2 mt-5"
    @submit-form="registerAction"
    :form-data="registerForm"
    :disabled="!correct"
  >
  <span v-if="!passwordIsCorrect" class="text-red-500 text-sm">
    密码至少6位，包含大小写字母、数字和!@#$%^&*? 中的一个
    <br>
  </span>
  <span v-else class="text-green-500 text-sm block mx-a">
    密码符合要求
  </span>
  <div class="flex flex-row justify-between">
      <div class="mx-a">
        <input type="checkbox" ref="rememberMe">
          <label for="rememberMe">记住我</label>
          <span class="text-sm text-gray-500">下次自动登录</span>
        </input>
      </div>
      <RouterLink class="mx-a text-purple-800" to="/auth/login">
        有账号了
      </RouterLink>
    </div>
    <button
      type="button"
      @click.prevent="sendCodeAction"
      :disabled="!telephoneIsCorrect"
      class="w-full h-10 bg-[#eb6b26] text-white border-0 text-lg cursor-pointer mt-5 rounded-[20px] flex justify-center items-center hover:bg-[#ff7e3b] disabled:bg-zinc-600"
    >
      {{ telephoneIsCorrect ? '发送验证码' : '请填写正确邮箱' }}
    </button>
  </FormContainer>
</template>