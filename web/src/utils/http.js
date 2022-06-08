//引入 axios
import axios from "axios";
import { ElMessage } from 'element-plus';

const http = axios.create({
  baseURL: '/',
  timeout: 50000
})

// 数据请求拦截
http.interceptors.request.use((config) => {
  return config;
}, (error) => {
  return Promise.reject(error);
});

// 返回响应数据拦截
http.interceptors.response.use((res) => {
  const data = res.data;
  if (res.status === 200) {
    return data;
  }
}, (error) => {
  if (error.response.status) {
    switch (error.response.status) {
      case 404:
        ElMessage({
          type: 'error',
          message: '请求路径找不到！',
          showClose: true
        });
        break;
      case 502:
        ElMessage({
          type: 'error',
          message: '服务器内部报错！',
          showClose: true
        });
        break;
      default:
        break;
    }
  }
  return Promise.reject(error);
});

export default http;
// 这是我原来的写法。
// post 请求方法
// export const post = (url, params) => {
//  return new Promise((resolve, reject) => {
//    http.post(url, params).then(res => {
//      resolve(res);
//    }).catch(error => {
//      reject(error);
//   })
//  });
//}
// get 请求方法
//export const get = (url) => {
//  return new Promise((resolve, reject) => {
//    http.get(url).then(res => {
//      resolve(res);
//    }).catch(error => {
//      reject(error);
//    })
//  });
//}