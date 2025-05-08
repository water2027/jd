<!-- jd_frontend/src/views/CreatePost.vue -->
<template>
  <div>
    <!-- 新增：基础测试文本 -->
    <p>这是编辑器页面，若显示则组件加载正常</p>


    <h1>发布新帖子</h1>
    <div id="toolbar-container"></div>
    <div id="editor-container" style="height: 500px;"></div>
    <button @click="submitPost">发布</button>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import '@wangeditor/editor/dist/css/style.css';
import { createEditor, createToolbar } from '@wangeditor/editor'; // 不再导入 UploadImageConfig
import axios from 'axios';

// 定义 editor 实例的 ref（类型为编辑器实例或 null）
const editorRef = ref<ReturnType<typeof createEditor> | null>(null);
const router = useRouter();

onMounted(() => {
  // 初始化编辑器
  const editor = createEditor({
    selector: '#editor-container',
    config: {
      MENU_CONF: {
        uploadImage: {
          server: '/api/media/upload-image',
          fieldName: 'custom-fileName', // 与服务端 save-files.js 中的 fieldName 一致（服务端代码中 fieldName 是文件字段名）
          maxFileSize: 5 * 1024 * 1024, // 5MB
          // 明确 onSuccess 参数的类型（根据服务端返回结构）
          onSuccess: (file: File, res: { errno: number; data: Array<{ url: string }> }) => {
            console.log('图片上传成功', res);
          }
        }
      }
    }
  });

  // 初始化工具栏
  const toolbar = createToolbar({
    editor,
    selector: '#toolbar-container'
  });

  // 保存 editor 实例到 ref
  editorRef.value = editor;
});

// 提交帖子
const submitPost = async () => {
  const editor = editorRef.value;
  if (!editor) {
    alert('编辑器未初始化');
    return;
  }

  const content = editor.getHtml();

  try {
    await axios.post('/api/post/create', { content });
    alert('帖子发布成功');
    router.push('/');
  } catch (error) {
    // 明确 error 类型（使用 Error 类型断言）
    alert('帖子发布失败：' + (error as Error).message);
  }
};
</script>

<style scoped>
/* 样式 */
</style>