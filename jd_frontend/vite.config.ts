import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import UnoCSS from 'unocss/vite'

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue(), vueJsx(), UnoCSS()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
})
