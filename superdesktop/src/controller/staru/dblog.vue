<script setup lang="ts">
import { ref } from "vue";
import { invoke } from "@tauri-apps/api/tauri";
import { writeText, readText } from "@tauri-apps/api/clipboard";

const dblog_input = ref("");
const dblog_output = ref("");
const display = ref();

async function dblog() {
  dblog_output.value = await invoke("dblog", { value: dblog_input.value });
  display.value.style = "display: block";
}

async function load() {
  dblog_input.value = await readText() as string;
}

async function copy() {
  await writeText(dblog_output.value);
}

async function clean() {
  dblog_input.value = ""
  dblog_output.value = ""
  display.value.style = "display: none";
}

async function flow() {
  await clean();
  await load();
  await dblog();
  await copy();
}
</script>

<template>
  <el-main>
    <el-row type="flex" justify="end">
      <el-input v-model="dblog_input" :rows="12" type="textarea" resize="none" placeholder="Please input"
        class="dblog_input" />
      <el-button type="primary" text auto-insert-space @click="flow();">VIP</el-button>
      <el-button type="primary" text auto-insert-space @click="load();">粘贴</el-button>
      <el-button type="primary" text auto-insert-space @click="dblog();">解析</el-button>
      <el-button type="primary" text auto-insert-space @click="copy();">复制</el-button>
      <el-button type="primary" text auto-insert-space @click="clean();">清空</el-button>
    </el-row>

    <div class="dblog_output" ref="display">
      <highlightjs language='sql' :code="dblog_output" />
    </div>

  </el-main>
</template>


<style scoped>
.dblog_output {
  display: none;
  width: 990px;
  max-height: 300px;
  margin-top: 10px;
}
</style>
