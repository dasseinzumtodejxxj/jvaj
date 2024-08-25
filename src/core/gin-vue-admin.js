/*
 * gin-vue-admin web框架组
 *
 * */
// 加载网站配置文件夹
import { register } from './global'

export default {
  install: (app) => {
    register(app)
    console.log(`
       欢迎使用 Gin-Xj-Vue-Admin
       当前版本:v2.5.7
       加群方式:微信：gsz19831210
       默认自动化文档地址:http://127.0.0.1:${import.meta.env.VITE_SERVER_PORT}/swagger/index.html
       默认前端文件运行地址:http://127.0.0.1:${import.meta.env.VITE_CLI_PORT}
       如果项目让您获得了收益，希望您能请团队喝杯可乐:
    `)
  }
}
