
<template>
  <el-container class="app-container">
    <BaseHeader />
    
    <el-container class="body-container">
      <BaseSide />
      <el-main>
        <RouterView />
      </el-main>
    </el-container>
  
  </el-container>
</template>
