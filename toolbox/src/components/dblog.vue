<script setup>
import { ref } from "vue";
import { invoke } from "@tauri-apps/api/tauri";
import { writeText, readText } from "@tauri-apps/api/clipboard";

const dblog_input = ref("");
const dblog_output = ref("");

async function dblog() {
  dblog_output.value = await invoke("dblog", {value: dblog_input.value});
}

async function load() {
  dblog_input.value = await readText();
}

async function copy() {
  await writeText(dblog_output.value);
}

async function clean() {
  dblog_input.value = ""
  dblog_output.value = ""
}

async function flow() {
  await clean();
  await load();
  await dblog();
  await copy();
}
</script>

<template>
  <div class="input-group flex-column">
    <lay-row space="25">
      <lay-col><lay-textarea v-model="dblog_input" allow-clear rows="13" style="resize: none;"></lay-textarea></lay-col>
      <lay-col md-offset="22"><lay-button type="primary" @click="flow()">运行</lay-button></lay-col>
      <lay-col><lay-textarea v-model="dblog_output" :disabled=true rows="13" style="resize: none;"></lay-textarea></lay-col>

    </lay-row>
    
  </div>
  
</template>

