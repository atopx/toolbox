<script setup lang="ts">
import {ref} from "vue";
import {invoke} from "@tauri-apps/api/tauri";

const source = ref({
    id: "1678849282905208556",
    host: "10.0.0.211"
})
const logs = ref([{
    timestr: "",
    type: "",
    service: "",
    infc: "",
    level: "",
    func: "",
    timestamp: "",
    message: "",
    title: "",
}]);
const display = ref(false);


async function trace() {
    logs.value = await invoke("tracelog", {id: source.value.id, host: source.value.host});
    display.value = true;
}

</script>
<template>
    <el-form :inline="true" :model="source" class="demo-form-inline">
        <el-form-item label="Server">
            <el-select v-model="source.host" placeholder="server host">
                <el-option label="QA" value="10.0.0.211"/>
                <el-option label="DEBUG" value="10.0.0.206"/>
            </el-select>
        </el-form-item>
        <el-form-item label="TraceId">
            <el-input v-model="source.id" placeholder="trace id"/>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="trace">查询</el-button>
        </el-form-item>
    </el-form>

    <div v-show="display" class="tracelogs">
        <el-scrollbar height="560px">
            <el-collapse accordion>
                <el-timeline style="margin-right: 40px;">
                    <el-timeline-item placement="top" v-for="(item, index) in logs" :key="index"
                                      :timestamp="item.timestr">
                        <Timeline>
                            <el-tag class="title-tag">{{ item.type }}</el-tag>
                            <el-tag class="title-tag" type="warning">{{ item.level }}</el-tag>
                            <el-tag class="title-tag" type="success">{{ item.service }}</el-tag>
                        </Timeline>
                        <el-collapse-item :title="item.infc">{{ item.message }}</el-collapse-item>
                    </el-timeline-item>
                </el-timeline>
            </el-collapse>
        </el-scrollbar>
    </div>
</template>

<style scoped>
.tracelogs {
    margin-top: 30px;
}

.title-tag {
    margin-right: 6px;
}
</style>