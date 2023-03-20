<template>
    <el-menu default-active="2" class="sidemenu" router>
        <el-sub-menu index="1">
            <template #title>
                <span>TODO</span>
            </template>
            <el-menu-item-group>
                <template #title><span>Group One</span></template>
                <el-menu-item index="1-1">item one</el-menu-item>
                <el-menu-item index="1-2">item two</el-menu-item>
            </el-menu-item-group>
            <el-menu-item-group title="Group Two">
                <el-menu-item index="1-3">item three</el-menu-item>
            </el-menu-item-group>
            <el-sub-menu index="1-4">
                <template #title><span>item four</span></template>
                <el-menu-item index="1-4-1">item one</el-menu-item>
            </el-sub-menu>
        </el-sub-menu>


        <el-sub-menu index="document">
            <template #title>
                <el-icon>
                    <Notebook />
                </el-icon>
                <span>文档管理</span>
            </template>
            <el-menu-item index="noteManage">笔记管理</el-menu-item>
            <el-menu-item index="blogManage">博客管理</el-menu-item>
        </el-sub-menu>


        <el-sub-menu index="server">
            <template #title>
                <el-icon>
                    <Monitor />
                </el-icon>
                <span>主机管理</span>
            </template>
            <el-menu-item index="/server/computerManage">电脑管理</el-menu-item>
            <el-menu-item index="/server/portManage">端口管理</el-menu-item>
        </el-sub-menu>


        <el-sub-menu index="user">
            <template #title>
                <el-icon>
                    <User />
                </el-icon>
                <span>角色管理</span>
            </template>
            <el-menu-item index="roleManage">角色管理</el-menu-item>
            <el-menu-item index="userManage">用户管理</el-menu-item>
            <el-menu-item index="permissionManage">权限管理</el-menu-item>
        </el-sub-menu>

        <el-menu-item index="setting">
            <el-icon>
                <setting />
            </el-icon>
            <template #title>基础设置</template>
        </el-menu-item>

    </el-menu>
</template>

<script lang="ts" setup>
import {
    User,
    Setting,
    Monitor,
    Notebook,
} from '@element-plus/icons-vue'

</script>


<style scoped>

.sidemenu{
    height: 100%;
}


</style>