<!-- jd_frontend/src/views/CreatePost.vue -->
<template>
  <div>
    <!-- 新增：基础测试文本 -->
    <p>这是编辑器页面，若显示则组件加载正常</p>

    <h1>发布新帖子</h1>
    <div id="toolbar-container"></div>
    <div id="editor-container" style="height: 500px"></div>
    <button @click="submitPost">发布</button>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, shallowRef } from 'vue'
import { useRouter } from 'vue-router'
import '@wangeditor/editor/dist/css/style.css'
import { createEditor, createToolbar } from '@wangeditor/editor' // 不再导入 UploadImageConfig
import { createPost } from '@/api/post/create'
import { showMsg } from '@/components/MessageBox'

// 定义 editor 实例的 ref（类型为编辑器实例或 null）
const editorRef = shallowRef<ReturnType<typeof createEditor> | null>(null)
const router = useRouter()

onMounted(() => {
  // 初始化编辑器
  const editor = createEditor({
    selector: '#editor-container',
    config: {
      MENU_CONF: {
        uploadImage: {
          server: 'http://localhost:8080/api/media/upload-image',
          maxFileSize: 5 * 1024 * 1024, // 5MB
          // 明确 onSuccess 参数的类型（根据服务端返回结构）
          onSuccess: (file: File, res: { errno: number; data: Array<{ url: string }> }) => {
            console.log('图片上传成功', res)
          },
        },
      },
    },
  })

  // 初始化工具栏
  const toolbar = createToolbar({
    editor,
    selector: '#toolbar-container',
  })

  // 保存 editor 实例到 ref
  editorRef.value = editor
})

// 提交帖子
const submitPost = async () => {
  const editor = editorRef.value
  if (!editor) {
    alert('编辑器未初始化')
    return
  }

  const content = editor.getHtml()

  try {
    await createPost({ content, title: '帖子标题' })
    showMsg('帖子发布成功')
    router.push('/')
  } catch (error) {
    // 明确 error 类型（使用 Error 类型断言）
    showMsg('帖子发布失败')
  }
}
</script>

<style scoped>
/* 样式 */
</style>
