<template>
  <el-menu default-active="2" class="el-menu-vertical-demo" @open="handleOpen" @close="handleClose" router>


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

    <el-sub-menu index="computer">
      <template #title>
        <el-icon>
          <Monitor />
        </el-icon>
        <span>主机管理</span>
      </template>

      <el-menu-item index="computerList">电脑管理</el-menu-item>
      <el-menu-item index="portList">端口管理</el-menu-item>
    </el-sub-menu>



    <el-sub-menu index="user">

      <template #title>
        <el-icon>
          <User />
        </el-icon>
        <span>角色管理</span>
      </template>

      <el-menu-item index="roleList">角色管理</el-menu-item>
      <el-menu-item index="userList">用户管理</el-menu-item>
      <el-menu-item index="permissionList">权限管理</el-menu-item>

    </el-sub-menu>

    <el-menu-item index="4">
      <el-icon>
        <setting />
      </el-icon>
      <template #title>基础设置</template>
    </el-menu-item>

  </el-menu>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import {
  User,
  Menu as IconMenu,
  Setting,
  Monitor,
} from '@element-plus/icons-vue'

const isCollapse = ref(true)
const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
</script>

