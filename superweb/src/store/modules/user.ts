import { ref } from "vue"
import store from "@/store"
import { defineStore } from "pinia"
import { useTagsViewStore } from "./tags-view"
import { getToken, removeToken, setToken } from "@/utils/cache/cookies"
import { resetRouter } from "@/router"
import { loginApi, getUserInfoApi, refreshTokenApi } from "@/api/login"
import { ILoginRequestData, IRefreshRequestData, Token } from "@/api/login/types/login"

export const useUserStore = defineStore("user", () => {
    const token = ref<Token>(getToken())
    const username = ref<string>("")

    const tagsViewStore = useTagsViewStore()

    // get token
    const login = (loginData: ILoginRequestData) => {
        return new Promise((resolve, reject) => {
            loginApi(loginData)
                .then((res) => {
                    res.data.expire_time = res.data.expire_time * 1000
                    setToken(res.data)
                    token.value = res.data
                    resolve(true)
                })
                .catch((error) => {
                    reject(error)
                })
        })
    }
    // refresh token
    const refreshToken = (refreshData: IRefreshRequestData) => {
        return new Promise((resolve, reject) => {
            refreshTokenApi(refreshData)
                .then((res) => {
                    res.data.expire_time = res.data.expire_time * 1000
                    setToken(res.data)
                    token.value = res.data
                    resolve(true)
                })
                .catch((error) => {
                    reject(error)
                })
        })
    }

    // get user info
    const getInfo = () => {
        return new Promise((resolve, reject) => {
            getUserInfoApi()
                .then((res) => {
                    const data = res.data
                    username.value = data.name
                    resolve(res)
                })
                .catch((error) => {
                    reject(error)
                })
        })
    }
    // 登出
    const logout = () => {
        resetToken()
        resetRouter()
        _resetTagsView()
    }
    // 重置 Token
    const resetToken = () => {
        removeToken()
        token.value = {
            id: 0,
            user_id: 0,
            access_token: "",
            refresh_token: "",
            expire_time: 0,
            issued_time: 0,
            create_time: 0,
            update_time: 0
        }
    }
    // 重置 visited views 和 cached views
    const _resetTagsView = () => {
        tagsViewStore.delAllVisitedViews()
        tagsViewStore.delAllCachedViews()
    }

    return { token, username, login, getInfo, logout, resetToken, refreshToken }
})

// 在 setup 外使用
export function useUserStoreHook() {
    return useUserStore(store)
}
