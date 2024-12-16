import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import axios from 'axios'
import store from './store'
import router from './router'
import elementplus from 'element-plus'
import 'element-plus/dist/index.css'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'


const app = createApp(App)
const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)
app.use(pinia)
app.use(router)
app.use(store)
app.use(elementplus)
app.config.globalProperties.$axios = axios;

app.mount('#app')



