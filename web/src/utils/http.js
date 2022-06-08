import axios from "axios";
import { ElMessage } from 'element-plus';

const http = axios.create({
	baseURL: '/',
	timeout: 50000
})

http.interceptors.request.use((config) => {
	return config;
}, (error) => {
	return Promise.reject(error);
});

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