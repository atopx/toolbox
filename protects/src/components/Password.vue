<script setup lang="ts">
import { ref } from "vue";
import { invoke } from "@tauri-apps/api/tauri";

const password = ref("");

async function setPassword() {
  password.value = await invoke("set_password", { password: password.value });
}
</script>

<template>
  <div class="card">
    <lay-row space="10">
      <lay-col><lay-input type="password" placeholder="输入你的密码" v-model="password"></lay-input></lay-col>
      <lay-col><lay-button type="primary" @click="setPassword()">提交密码</lay-button></lay-col>
    </lay-row>
    <br />
    <br />
    <br />
    <p style="color: red;">提示: 如果是第一次使用, 自动设置为管理密码，请记好, 以后每次打开需要输入密码验证。</p>
  </div>

</template>
