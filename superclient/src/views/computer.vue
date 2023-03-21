<script lang="ts" setup>

import { ref, onMounted } from 'vue';
import { Computer } from '~/apis/computers';
import { Page } from '~/common/types';
import { ArrowDown } from '@element-plus/icons-vue'

const payload = ref({
    page: <Page>{
        index: 1,
        size: 10,
        disabled: false,
    },
    filter: {}
})


const computers = ref();

const createFormData = ref({
    name: '',
    address: '',
    lan_hostname: '',
    wan_hostname: '',
    username: '',
    password: '',
    default_ports: [],

    display: false,
})

const updateFormData = ref({
    id: 0,
    name: "",
    address: "",
    lan_hostname: "",
    wan_hostname: "",
    username: "",
    password: "",

    display: false,
})


onMounted(async () => {
    await listComputer();
})

async function listComputer() {
    await Computer.list(payload.value).then((resp => {
        computers.value = resp.data.search;
    })).catch((err) => {
        console.log(err);
    })
}

async function createComputer() {
    let computer = createFormData.value;
    await Computer.create({
        name: computer.name,
        address: computer.address,
        lan_hostname: computer.lan_hostname,
        wan_hostname: computer.wan_hostname,
        username: computer.username,
        password: computer.password,
        default_ports: computer.default_ports,
    })
}

async function updateComputer() {
    let computer = updateFormData.value;
    await Computer.update({
        id: computer.id,
        name: computer.name,
        address: computer.address,
        lan_hostname: computer.lan_hostname,
        wan_hostname: computer.wan_hostname,
        username: computer.username,
        password: computer.password,
    });
}

async function deleteComputer(id: number) {
    await Computer.delete({ id: id });
}

async function operateComputer(id: number, operate: number) {
    await Computer.operate({ id: id, operate: operate });
}

</script>

<template>

    <!-- TODO 搜索表单 -->


    <!-- TODO 新建表单弹窗 -->


    <!-- TODO 更新表单弹窗 -->


    <!-- TODO 删除确认弹窗 -->


    <!-- 数据列表 -->
    <el-table :data="computers" style="width: 100%" :header-cell-style="{ 'text-align': 'center' }"
        :cell-style="{ 'text-align': 'center' }">
        <el-table-column prop="id" label="ID" width="50" />
        <el-table-column prop="name" label="主机名" width="120" />
        <el-table-column prop="address" label="MAC" width="200" />
        <!-- <el-table-column prop="lan_hostname" label="内网地址" /> -->
        <!-- <el-table-column prop="wan_hostname" label="外网地址" /> -->
        <el-table-column prop="username" label="用户名" />
        <!-- <el-table-column prop="password" label="密码" /> -->
        <el-table-column prop="create_time" label="创建时间" />
        <el-table-column prop="update_time" label="更新时间" />
        <el-table-column fixed="right" label="操作" width="160px">
            <template #default="{ item }">
                <el-dropdown>

                    <el-button type="primary">
                        管理<el-icon class="el-icon--right"><arrow-down /></el-icon>
                    </el-button>

                    <template #dropdown>
                        <el-dropdown-menu>
                            <!-- TODO 修改弹窗 -->
                            <el-dropdown-item>修改</el-dropdown-item>
                            <el-dropdown-item @click="deleteComputer(item.id)">删除</el-dropdown-item>
                            <el-dropdown-item @click="operateComputer(item.id, 1)">开机</el-dropdown-item>
                            <el-dropdown-item @click="operateComputer(item.id, 0)">关机</el-dropdown-item>
                            <el-dropdown-item @click="operateComputer(item.id, 2)">扫描</el-dropdown-item>
                        </el-dropdown-menu>
                    </template>
                </el-dropdown>
            </template>
        </el-table-column>

    </el-table>
</template>


<style scoped>
.example-showcase .el-dropdown-link {
    cursor: pointer;
    color: var(--el-color-primary);
    display: flex;
    align-items: center;
}
</style>