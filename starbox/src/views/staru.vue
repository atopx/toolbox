<script setup lang="ts">
import { ref } from "vue";
import dblog from "../controller/staru/dblog.vue";
import permission from "../controller/staru/permission.vue";
import trace from "../controller/staru/trace.vue";

const activeTabname = ref('trace')
</script>

<template>
  <el-tabs v-model="activeTabname" class="tabs">
    <el-tab-pane label="日志追踪" name="trace">
      <trace />
    </el-tab-pane>
    <el-tab-pane label="日志解析" name="parse">
      <dblog />
    </el-tab-pane>
    <el-tab-pane label="接口权限" name="permission">
      <permission />
    </el-tab-pane>
    <el-tab-pane label="Task" name="fourth">Task</el-tab-pane>
  </el-tabs>
</template>