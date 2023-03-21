<script lang="ts" setup>
import { reactive, ref } from 'vue'
import { FormInstance, FormRules } from 'element-plus';
import http from '../common/requestt';

const loginFormRef = ref<FormInstance>();

const loginForm = ref({
    username: '',
    password: '',
});


const validateUsername = (rule: any, value: any, callback: any) => {
    if (value.length < 5 || value.length > 24) {
        callback(new Error('无效的登录账号'))
    } else {
        callback()
    }
}

const validatePassword = (rule: any, value: any, callback: any) => {
    if (value.length < 6 || value.length > 24) {
        callback(new Error('无效的登录密码'))
    } else {
        callback()
    }
}

const loginRule = reactive<FormRules>({
    username: [{ required: true, trigger: 'blur', validator: validateUsername }],
    password: [{ required: true, trigger: 'blur', validator: validatePassword }]
})

const loading = ref(false);

async function login() {
    await http.post("/user/login", loginForm.value).then((resp) => {
        // set token to local storage
        window.localStorage.setItem("token", resp.data.data.token);
        console.log("success: ", resp.data.data.token);
    }).catch((err) => {
        console.log("err: " + err);
    })
}

</script>

<template>
    <div class="login-container">
        <el-form ref="loginFormRef" :model="loginForm" :rules="loginRule" class="login-form" autocomplete="on"
            label-position="left">

            <div class="title-container">
                <h3 class="title">Super Login</h3>
            </div>

            <el-form-item prop="username">
                <span class="svg-container">
                    <svg-icon icon-class="user" />
                </span>
                <el-input ref="username" v-model="loginForm.username" placeholder="登录账号" name="username" type="text"
                    tabindex="1" autocomplete="on" />
            </el-form-item>

            <el-form-item prop="password">
                <span class="svg-container">
                    <svg-icon icon-class="password" />
                </span>
                <el-input ref="password" v-model="loginForm.password" type="password" placeholder="登录密码" name="password"
                    autocomplete="on" />
            </el-form-item>

            <el-button :loading="loading" type="primary" style="width:100%;margin-bottom:30px;" @click="login()">登录</el-button>

        </el-form>

    </div>
</template>


<style lang="scss" scoped>
$bg: #2d3a4b;
$dark_gray: #889aa4;
$light_gray: #eee;

.login-container {
    min-height: 100%;
    width: 100%;
    background-color: $bg;
    overflow: hidden;

    .login-form {
        position: relative;
        width: 500px;
        max-width: 100%;
        padding: 160px 35px 0;
        margin: 0 auto;
        overflow: hidden;
    }

    .svg-container {
        padding: 6px 5px 6px 15px;
        color: $dark_gray;
        vertical-align: middle;
        width: 30px;
        display: inline-block;
    }

    .title-container {
        position: relative;

        .title {
            font-size: 26px;
            color: $light_gray;
            margin: 0px auto 40px auto;
            text-align: center;
            font-weight: bold;
        }
    }

}</style>