<script lang="ts" setup>
import { reactive, ref } from "vue"
import { useRouter } from "vue-router"
import { useUserStore } from "@/store/modules/user"
import { User, Lock } from "@element-plus/icons-vue"
import ThemeSwitch from "@/components/ThemeSwitch/index.vue"
import { type FormInstance, FormRules } from "element-plus"
import { type ILoginRequestData } from "@/api/login/types/login"

const router = useRouter()
const loginFormRef = ref<FormInstance | null>(null)

/** 登录按钮 Loading */
const loading = ref(false)
/** 登录表单数据 */
const loginForm: ILoginRequestData = reactive({
    username: "",
    password: ""
})
/** 登录表单校验规则 */
const loginFormRules: FormRules = {
    username: [{ required: true, message: "请输入用户名", trigger: "blur" }],
    password: [
        { required: true, message: "请输入密码", trigger: "blur" },
        { min: 7, max: 16, message: "长度在 7 到 16 个字符", trigger: "blur" }
    ],
    code: [{ required: true, message: "请输入验证码", trigger: "blur" }]
}
/** 登录逻辑 */
const handleLogin = () => {
    loginFormRef.value?.validate((valid: boolean) => {
        if (valid) {
            loading.value = true
            useUserStore()
                .login({
                    username: loginForm.username,
                    password: loginForm.password
                })
                .then(() => {
                    router.push({ path: "/" })
                })
                .catch(() => {
                    loginForm.password = ""
                })
                .finally(() => {
                    loading.value = false
                })
        } else {
            return false
        }
    })
}

/** 初始化验证码 */
</script>

<template>
    <div class="login-container">
        <ThemeSwitch class="theme-switch" />
        <div class="login-card">
            <div class="title">
                <img src="@/assets/layout/logo-text-2.png" />
            </div>
            <div class="content">
                <el-form ref="loginFormRef" :model="loginForm" :rules="loginFormRules" @keyup.enter="handleLogin">
                    <el-form-item prop="username">
                        <el-input
                            v-model.trim="loginForm.username"
                            placeholder="用户名"
                            type="text"
                            tabindex="1"
                            :prefix-icon="User"
                            size="large"
                        />
                    </el-form-item>
                    <el-form-item prop="password">
                        <el-input
                            v-model.trim="loginForm.password"
                            placeholder="密码"
                            type="password"
                            tabindex="2"
                            :prefix-icon="Lock"
                            size="large"
                            show-password
                        />
                    </el-form-item>
                    <el-button :loading="loading" type="primary" size="large" @click.prevent="handleLogin">
                        登 录
                    </el-button>
                </el-form>
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.login-container {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
    min-height: 100%;
    .theme-switch {
        position: fixed;
        top: 5%;
        right: 5%;
        cursor: pointer;
    }
    .login-card {
        width: 480px;
        border-radius: 20px;
        box-shadow: 0 0 10px #dcdfe6;
        background-color: #fff;
        overflow: hidden;
        .title {
            display: flex;
            justify-content: center;
            align-items: center;
            height: 150px;
            img {
                height: 100%;
            }
        }
        .content {
            padding: 20px 50px 50px 50px;
            :deep(.el-input-group__append) {
                padding: 0;
                overflow: hidden;
                .el-image {
                    width: 100px;
                    height: 40px;
                    border-left: 0px;
                    user-select: none;
                    cursor: pointer;
                    text-align: center;
                }
            }
            .el-button {
                width: 100%;
                margin-top: 10px;
            }
        }
    }
}
</style>
