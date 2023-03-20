<script lang="ts" setup>

import { ref } from 'vue'
import http from '../common/common';


interface Computer {
    id: string,
    name: string,
    address: string,
    lan_hostname: string,
    wan_hostname: string,
    username: string,
    password: string,
    create_time: string,
    update_time: string,
}

const payload = ref({
    page: {
        "disabled": false,
        "index": 1,
        "size": 10
    },
})
const computers = ref<Array<Computer>>()

// TODO 不知道为什么: Uncaught (in promise) TypeError: Cannot read properties of null (reading 'parentNode')
let resp = await http.post("/computer/search", payload);
if (resp.status == 200) {
    console.log(resp.data);
} else {
    console.log("error");
}


</script>

<template>
    <el-table :data="computers" style="width: 100%">
        <el-table-column prop="id" label="ID" width="180" />
        <el-table-column prop="name" label="名称" width="180" />
        <el-table-column prop="address" label="MAC地址" />
        <el-table-column prop="lan_hostname" label="内网" />
        <el-table-column prop="wan_hostname" label="外网" />
        <el-table-column prop="username" label="用户名" />
        <el-table-column prop="password" label="密码" />
        <el-table-column prop="create_time" label="创建时间" />
        <el-table-column prop="update_time" label="更新时间" />
    </el-table>
</template>