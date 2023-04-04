<script lang="ts" setup>
import MonacoEditor from '../../component/editor.vue'
import {onMounted, ref} from "vue";

const input_code = ref("{}")
const output_code = ref("")

</script>

<template>
    <el-row :gutter="2">

        <el-col :span="12">
            <MonacoEditor height="860px" v-model="input_code" language="json" :readonly="false"/>
        </el-col>

        <el-col :span="12">
            <MonacoEditor height="860px" v-model="output_code" language="json" :readonly="true"/>
        </el-col>

    </el-row>
</template>

