/**
 * @description 服务端路由
 * @author wangfupeng
 */

const router = require('koa-router')()
const posts = []; // 临时存储（实际需替换为数据库）
const saveFiles = require('./controller/save-file')
const saveVideoFile = require('./controller/saveVideo-file')

router.prefix('/api')

// 测试用
router.get('/test', async function (ctx, next) {
    ctx.body = 'See other routes in `src/router.js`'
})

// 上传图片
// 新增保存帖子接口
router.post('/api/create-post', async function (ctx, next) {
    const { content } = ctx.request.body
    const newPost = {
      id: Date.now().toString(), // 生成唯一 ID
      content,
      createTime: new Date().toISOString()
    }
    posts.push(newPost)
    ctx.body = { errno: 0, data: newPost }
  })

// 上传图片 - 用于测试超时
router.post('/upload-img-10s', async function (ctx, next) {
    const data = await saveFiles(ctx.req, 10 * 1000)
    ctx.body = data
})

// 上传图片 - 测试 failed
router.post('/upload-img-failed', async function (ctx, next) {
    ctx.body = { errno: 1, message: 'upload-img-failed - 发送失败' }
})

// 新增获取帖子列表接口（用于左侧导航显示）
router.get('/api/posts', async function (ctx, next) {
    ctx.body = { errno: 0, data: posts }
  })
module.exports = router
