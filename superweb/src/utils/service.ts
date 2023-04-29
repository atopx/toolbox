import axios, { type AxiosInstance, type AxiosRequestConfig } from "axios"
import { useUserStoreHook } from "@/store/modules/user"
import { ElMessage } from "element-plus"
import { get } from "lodash-es"
import { ignore_apis } from "@/api/ignore"

/** 创建请求实例 */
function createService() {
    // 创建一个 Axios 实例
    const service = axios.create()
    // 请求拦截
    service.interceptors.request.use(
        (config) => {
            if (ignore_apis.includes(config.url || "")) {
                return config
            }
            if (!useUserStoreHook().token.access_token) {
                Promise.reject()
            }
            if (Date.now() > useUserStoreHook().token.expire_time) {
                useUserStoreHook().refreshToken({ refresh_token: useUserStoreHook().token.refresh_token })
            }
            config.headers.Authorization = "Bearer " + useUserStoreHook().token.access_token
            return config
        },
        // 发送失败
        (error) => Promise.reject(error)
    )
    // 响应拦截（可根据具体业务作出相应的调整）
    service.interceptors.response.use(
        (response) => {
            if (response.data.header.code > 0) {
                ElMessage.error(response.data.header.message)
                return Promise.reject(new Error("Error"))
            }
            return response.data
        },
        (error) => {
            // Status 是 HTTP 状态码
            const status = get(error, "response.status")
            switch (status) {
                case 400:
                    error.message = "请求错误"
                    break
                case 401:
                    // Token 过期时，直接退出登录并强制刷新页面（会重定向到登录页）
                    useUserStoreHook().logout()
                    location.reload()
                    break
                case 403:
                    error.message = "拒绝访问"
                    break
                case 404:
                    error.message = "请求地址出错"
                    break
                case 408:
                    error.message = "请求超时"
                    break
                case 500:
                    error.message = "服务器内部错误"
                    break
                case 501:
                    error.message = "服务未实现"
                    break
                case 502:
                    error.message = "网关错误"
                    break
                case 503:
                    error.message = "服务不可用"
                    break
                case 504:
                    error.message = "网关超时"
                    break
                case 505:
                    error.message = "HTTP 版本不受支持"
                    break
                default:
                    break
            }
            ElMessage.error(error.message)
            return Promise.reject(error)
        }
    )
    return service
}

/** 创建请求方法 */
function createRequestFunction(service: AxiosInstance) {
    return function <T>(config: AxiosRequestConfig): Promise<T> {
        const configDefault = {
            headers: { "Content-Type": get(config, "headers.Content-Type", "application/json") },
            timeout: 60000,
            baseURL: import.meta.env.VITE_BASE_API,
            data: {}
        }
        return service(Object.assign(configDefault, config))
    }
}

/** 用于网络请求的实例 */
export const service = createService()
/** 用于网络请求的方法 */
export const request = createRequestFunction(service)
