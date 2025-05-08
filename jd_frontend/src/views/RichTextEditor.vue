<template>
    <div id="editor-container"></div>
    <div id="toolbar-container"></div>
  </template>
  
  <script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import { createEditor, createToolbar } from '@wangeditor/editor-for-vue';
  import axios from 'axios';
  
  const editorRef = ref(null);
  const toolbarRef = ref(null);
  
  onMounted(() => {
    const editor = createEditor({
      selector: '#editor-container',
      config: {
        MENU_CONF: {
          uploadImage: {
            server: '/upload-img', // 这里需要替换为你实际的图片上传接口
            fieldName: 'image', // 上传图片的字段名
            maxFileSize: 5 * 1024 * 1024, // 最大文件大小，单位为字节
            meta: {
              token: localStorage.getItem('token') // 如果需要传递 token 等元数据，可以在这里设置
            }
          }
        }
      }
    });
    const toolbar = createToolbar({
      editor,
      selector: '#toolbar-container'
    });
    editorRef.value = editor;
    toolbarRef.value = toolbar;
  });
  
  // 提交表单的方法
  const submitForm = async () => {
    const content = editorRef.value.getHtml();
    try {
      const response = await axios.post('/api/create-post', { content });
      alert('帖子发布成功');
    } catch (error) {
      alert('帖子发布失败');
    }
  };
  </script>
  
  <style scoped>
  /* 你可以在这里添加编辑器的样式 */
  </style>