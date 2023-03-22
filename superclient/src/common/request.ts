import axios, { AxiosInstance, AxiosRequestConfig } from 'axios';
import router from '~/router';

const http: AxiosInstance = axios.create({
    baseURL: '/api',
    timeout: 60 * 1000,
    headers: { 'Content-Type': 'application/json' },
    paramsSerializer: {
		serialize(params) {
			return JSON.stringify(params)
		},
	},
})


http.interceptors.request.use(
    function (config) {
        const token = window.localStorage.getItem('token');
        if (token) {
            config.headers.Authorization = token;
        }
        return config;
    },
    function (error) {
        // Do something with request error
        return Promise.reject(error);
    }
);

// Add a response interceptor
http.interceptors.response.use(
    function (response) {
        return response.data;
    },
    function (error) {
        if (error.status == 401) {
            router.push({ path: '/login' })
        }
        // Do something with response error
        return Promise.reject(error);
    }
);

export default http;