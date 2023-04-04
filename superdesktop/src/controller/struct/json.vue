<script lang="ts" setup>
import MonacoEditor from '../../component/editor.vue'
import {ref} from "vue";

const input_code = ref("{}")
const output_code = ref("")

const language_value = ref('golang')

const language_options = [
    "golang",
    "rust",
    "java",
    "javascript",
    "typescript",
    "yaml",
    "toml",
]

function transTo() {
    output_code.value = input_code.value;
}

</script>

<template>
    <div class="control-btn">
        <el-row :gutter="20">
            <el-col :span="3">
                <el-select v-model="language_value" filterable placeholder="Select">
                    <el-option
                            v-for="option in language_options"
                            :key="option"
                            :label="option"
                            :value="option"
                    />
                </el-select>
            </el-col>
            <el-col :span="4">
                <el-button type="primary" @click="transTo">转换</el-button>
            </el-col>
        </el-row>
    </div>
    <div>
        <el-row :gutter="2">
            <el-col :span="12">
                <MonacoEditor height="860px" v-model="input_code" language="json" :readonly="false"/>
            </el-col>
            <el-col :span="12">
                <MonacoEditor height="860px" v-model="output_code" language="json" :readonly="true"/>
            </el-col>
        </el-row>
    </div>
</template>

<style scoped>
.control-btn {
    margin-bottom: 10px;
}
</style>

