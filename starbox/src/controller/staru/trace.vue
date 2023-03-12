<script setup lang="ts">
import { ref } from "vue";
import { invoke } from "@tauri-apps/api/tauri";
import { Aim } from '@element-plus/icons-vue'

const traceId = ref("");
const logs = ref([{
    timestr: "",
    type:"",
    service: "",
    infc: "",
    timestamp: "",
    message: ""
}]);
const display = ref();


async function trace() {
    logs.value = await invoke("tracelog", { id: traceId.value });
    display.value.style = "display:block"
}

</script>
<template>
    <el-input v-model="traceId" placeholder="输入TraceID">
        <template #append>
            <el-button type="primary" @click="trace()" :icon="Aim" />
        </template>
    </el-input>


    <div ref="display" class="tracelogs">
        <el-scrollbar height="560px">
            <el-collapse accordion>
                <el-timeline style="margin-right: 30px;">
                    <el-timeline-item placement="top" v-for="(item, index) in logs" :key="index" :timestamp="item.timestr">
                        <el-collapse-item :title="item.type + item.service + item.infc" :name="item.timestamp">{{ item.message }}</el-collapse-item>
                    </el-timeline-item>
                </el-timeline>
            </el-collapse>
    </el-scrollbar>
    </div>

</template>

<style scoped>
    .tracelogs {
        display: none;
        margin-top: 40px;
    }

</style>