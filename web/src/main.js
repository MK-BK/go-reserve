import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import axios from 'axios'

const app = createApp(App)
app.config.globalProperties.$axios = axios

// app.axios.interceptors.request.use(function(config) {
//     if (sessionStorage.getItem('token')) {
//         config.headers['token'] = sessionStorage.getItem('token')
//     }
//     return config;
//     }, function(error) {
//     return Promise.reject(error);
// })

// app.axios.interceptors.response.use(function(response) {
//     if (response.config.url.indexOf('login') > -1) {
//         return response;
//     }
//     return response.data;
//     }, function(error) {
//     return Promise.reject(error)
// })

app.
    use(store).
    use(router).
    use(ElementPlus).
    mount('#app')

