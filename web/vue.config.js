const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    host: 'localhost',
    port: 8080,
    proxy:{
      '/api':{
        target: 'http://localhost:10086',
        changeOrigin: true,
        pathRewrite: {
          '^/api': '/' 
        }
      }
    }
  }
})
