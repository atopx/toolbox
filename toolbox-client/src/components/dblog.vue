<script setup lang="ts">
import {ref} from "vue";
import {invoke} from "@tauri-apps/api/tauri";

const dblog_input = ref("");
const dblog_output = ref("");


async function dblog() {
  dblog_output.value = await invoke("dblog", {value: dblog_input.value});
}

async function load() {
  dblog_input.value = await navigator.clipboard.readText();
}

async function copy() {
  await navigator.clipboard.writeText(dblog_output.value);
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
    <h5>DBLOG</h5>
    <div>
      <textarea class="form-control" v-model="dblog_input" aria-label="db-log" rows="15"></textarea>
    </div>
    <div class="d-flex flex-row-reverse">
      <button type="button" class="btn btn-success" @click="dblog()" style="margin: 10px">解析</button>
      <button type="button" class="btn btn-success" @click="copy()" style="margin: 10px">复制</button>
      <button type="button" class="btn btn-success" @click="load()" style="margin: 10px">读取</button>
      <button type="button" class="btn btn-success" @click="flow()" style="margin: 10px">一键操作</button>
      <button type="button" class="btn btn-success" @click="clean()" style="margin: 10px">清空</button>
    </div>
  </div>

  <pre>{{ dblog_output }}</pre>
</template>


<style>
pre {
  white-space: pre-wrap; /*css-3*/
  white-space: -moz-pre-wrap; /*Mozilla,since1999*/
  white-space: -o-pre-wrap; /*Opera7*/
  word-wrap: break-word; /*InternetExplorer5.5+*/
}
</style>