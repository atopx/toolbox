import axios from 'axios';
import router from '~/router';

let config = {
    baseURL: '/api',
    timeout: 60 * 1000,
    withCredentials: true,
};

const http = axios.create(config);



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
        return response;
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