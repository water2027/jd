<template>
  <div class="flex flex-col items-center justify-center w-full px-10% overflow-auto">
    <h1 class="text-center">发布新帖子</h1>
    <input class="w-50% p-2% mb-2" type="text" v-model="title" placeholder="帖子标题" />
    <div class="w-full" id="toolbar-container"></div>
    <div class="w-full h-xl" id="editor-container"></div>
    <button class="p-2% rounded-2xl" @click="submitPost">发布</button>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, shallowRef, ref } from 'vue'
import { useRouter } from 'vue-router'
import '@wangeditor/editor/dist/css/style.css'
import { createEditor, createToolbar } from '@wangeditor/editor' 
import { createPost } from '@/api/post/create'
import { showMsg } from '@/components/MessageBox'

const title = ref('') // 帖子标题

// 定义 editor 实例的 ref（类型为编辑器实例或 null）
const editorRef = shallowRef<ReturnType<typeof createEditor> | null>(null)
const toolbarRef = shallowRef<ReturnType<typeof createToolbar> | null>(null)
const router = useRouter()

onMounted(() => {
  // 初始化编辑器
  const editor = createEditor({
    selector: '#editor-container',
    config: {
      MENU_CONF: {
        uploadImage: {
          server: `${import.meta.env.VITE_API_URL}/media/upload`,
          maxFileSize: 5 * 1024 * 1024, // 5MB
        },
        uploadVideo: {
          server: `${import.meta.env.VITE_API_URL}/media/upload`,
          maxFileSize: 10 * 1024 * 1024, // 10MB
          onSuccess(file: File, res: any) {
            const { data } = res
            // 不知道为什么必须要手动插入，否则上传后什么都不会发生，跟图片不一样
            for (const video of data) {
              editor.insertNode({
                // @ts-ignore
                type: 'video',
                src: video.url,
                poster: video.poster,
                children: [{ text: '' }],
              })
            }
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
  toolbarRef.value = toolbar
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
    await createPost({ content, title: title.value })
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
