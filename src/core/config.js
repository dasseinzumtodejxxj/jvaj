/**
 * 网站配置文件
 */

const config = {
  appName: 'gxva',
  appLogo: 'https://www.gin--vue-admin.com/img/logo.png',
  showViteLogo: true
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    const chalk = require('chalk')
    console.log(
      chalk.green(
     )
    )
    console.log(
      chalk.green(
        `> 当前版本:v2.5.7`
      )
    )
    console.log(
      chalk.green(
        `> 加群方式:微信：gsz19831210`
      )
    )
    console.log(
      chalk.green(
        ``
      )
    )
    console.log(
      chalk.green(
        `> 插件市场:暂无`
      )
    )
    console.log(
      chalk.green(
        `> 默认自动化文档地址:http://127.0.0.1:${env.VITE_SERVER_PORT}/swagger/index.html`
      )
    )
    console.log(
      chalk.green(
        `> 默认前端文件运行地址:http://127.0.0.1:${env.VITE_CLI_PORT}`
      )
    )
    console.log(
      chalk.green(
        `> 如果项目让您获得了收益，希望您能请团队喝杯可乐`
      )
    )
    console.log('\n')
  }
}

export default config
