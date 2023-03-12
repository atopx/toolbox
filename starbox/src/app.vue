<script setup lang="ts">

</script>

<template>
  <div class="common-layout">
    <el-container>
      <el-aside width="80px">
        <el-menu active-text-color="#ffd04b" background-color="#545c64" router text-color="#fff" style="height: 772px;">
          <el-menu-item index="system">系统</el-menu-item>
          <el-menu-item index="staru">星合</el-menu-item>
          <el-menu-item index="network">网络</el-menu-item>
          <el-menu-item index="3">TO2</el-menu-item>
          <el-menu-item index="4">TO2</el-menu-item>
          <el-menu-item index="5">TO2</el-menu-item>
          <el-menu-item index="novel">藏经</el-menu-item>
          <el-menu-item index="7">TO2</el-menu-item>
          <el-menu-item index="8">TO2</el-menu-item>
          <el-menu-item index="9">TO2</el-menu-item>
          <el-menu-item index="t1">TO2</el-menu-item>
          <el-menu-item index="t2">TO2</el-menu-item>
          <el-menu-item index="t3">TO2</el-menu-item>
        </el-menu>
      </el-aside>

      <el-main><router-view></router-view></el-main>

    </el-container>

  </div>
</template>

<style scoped>
.side {
  background-color: #545c64;
  height: 1000px;
}
</style>

