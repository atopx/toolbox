<script lang="ts" setup>
</script>

<template>
    <el-menu class="el-menu-demo" mode="horizontal" router style="height: 60px; position: fixed; z-index: 1; width: 100%;">

      <!-- logo区域 -->
      <el-menu-item index="login">Spuer</el-menu-item>
  
      <el-menu-item index="3">Info</el-menu-item>
      <el-menu-item index="4">Orders</el-menu-item>
  
      <!-- 多层级菜单 -->
      <el-sub-menu index="2">
        <template #title>Workspace</template>
        <el-menu-item index="2-1">item one</el-menu-item>
        <el-menu-item index="2-2">item two</el-menu-item>
        <el-menu-item index="2-3">item three</el-menu-item>
        <el-sub-menu index="2-4">
          <template #title>item four</template>
          <el-menu-item index="2-4-1">item one</el-menu-item>
          <el-menu-item index="2-4-2">item two</el-menu-item>
          <el-menu-item index="2-4-3">item three</el-menu-item>
        </el-sub-menu>
      </el-sub-menu>
    </el-menu>
</template>
