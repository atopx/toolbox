<script setup lang="ts">
import commands from "../controller/system/commands.vue";
import sysinfo from "../controller/system/sysinfo.vue";

</script>

<template>
    <el-tabs class="tabs">
        
        <el-tab-pane label="监控" name="trace">
            <sysinfo />
        </el-tab-pane>
        
        <el-tab-pane label="进程" name="parse">
            2
        </el-tab-pane>
        
        <el-tab-pane label="端口" name="permission">
            3
        </el-tab-pane>
        
        <el-tab-pane label="命令" name="commands">
           <commands />
        </el-tab-pane>

        <el-tab-pane label="日志" name="syslog">
           <commands />
        </el-tab-pane>
    
    </el-tabs>
</template>