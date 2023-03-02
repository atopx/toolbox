<script setup lang="ts">

</script>


<template>
  <div class="toolbox">

    <el-container>
      <el-aside width="120px" class="side">
        <el-menu active-text-color="#ffd04b" background-color="#545c64" router text-color="#fff">
          <el-menu-item index="0">TODO 0</el-menu-item>
          <el-menu-item index="greet">百宝箱</el-menu-item>
          <el-menu-item index="novel">藏经阁</el-menu-item>
          <el-menu-item index="3">TODO 3</el-menu-item>
          <el-menu-item index="4">TODO 4</el-menu-item>
          <el-menu-item index="5">TODO 5</el-menu-item>
          <el-menu-item index="6">TODO 6</el-menu-item>
          <el-menu-item index="7">TODO 7</el-menu-item>
          <el-menu-item index="8">TODO 8</el-menu-item>
          <el-menu-item index="9">TODO 9</el-menu-item>
        </el-menu>
      </el-aside>

      <el-main><router-view></router-view></el-main>
    </el-container>
  </div>
</template>

<style scoped>
.side {
  background-color: #545c64;
  height: 780px;
}
</style>
