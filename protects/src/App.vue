<script setup lang="ts">
import Password from "./components/Password.vue";
</script>


<template>
  <lay-layout class="example">
    <lay-header>文件夹保护助手</lay-header>
    <lay-body>
      <lay-container fluid>
        <div class="body">
          <Password />
        </div>
      </lay-container>
    </lay-body>
  </lay-layout>
</template>
