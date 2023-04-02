<script setup lang="ts">
import {ref} from "vue";
import Permission from "../controller/staru/permission.vue";
import DbLog from "../controller/staru/dblog.vue";
import Trace from "../controller/staru/trace.vue";

const activeTab = ref('trace')
</script>

<template>
    <el-tabs v-model="activeTab" class="tabs">
        <el-tab-pane label="日志追踪" name="trace">
            <Trace/>
        </el-tab-pane>
        <el-tab-pane label="日志解析" name="parse">
            <DbLog/>
        </el-tab-pane>
        <el-tab-pane label="接口权限" name="permission">
            <Permission/>
        </el-tab-pane>
        <el-tab-pane label="Task" name="fourth">Task</el-tab-pane>
    </el-tabs>
</template>