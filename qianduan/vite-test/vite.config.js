import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
//解决跨域问题
export default defineConfig({
  plugins: [vue()],
  transplieDependencies: true,
  server: {
    //  开发服务器配置
    // host: '0.0.0.0',
    // port: 5173, // 你的开发服务器端口
    proxy: {
      // 匹配所有以 '/api' 开头的请求
      '/api': {
        target: 'http://localhost:8080', // 后端服务器地址
        changeOrigin: true, // 是否开启代理服务器，默认false，这里改为true
        rewrite: (path) => path.replace(/^\/api/, ''), // 重写路径，去除请求路径中的'/api'
      },
    },
  },
})
