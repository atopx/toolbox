<script setup lang="ts">
import { ref } from "vue";
import home from "../components/novel/home.vue";


</script>

<template>
    <el-tabs>
        <el-tab-pane label="index" name="index"><home/></el-tab-pane>
        <el-tab-pane label="Config" name="second">Config</el-tab-pane>
        <el-tab-pane label="Role" name="third">Role</el-tab-pane>
        <el-tab-pane label="Task" name="fourth">Task</el-tab-pane>
    </el-tabs>
</template>
