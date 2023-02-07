<script setup lang="ts">
import {ref} from "vue";
import {invoke} from "@tauri-apps/api/tauri";

const dblog_output = ref("");
const dblog_input = ref("");

async function dblog() {
  // Learn more about Tauri commands at https://tauri.app/v1/guides/features/command
  dblog_output.value = await invoke("dblog", {name: dblog_input.value});
}
</script>

<template>
  <div class="input-group flex-column">
    <h5>DBLOG</h5>
    <div>
      <textarea class="form-control" v-model="dblog_input" aria-label="db-log" rows="5"></textarea>
    </div>
    <div class="d-flex flex-row-reverse">
      <button type="button" class="btn btn-success" @click="dblog()" style="margin: 10px">解析</button>
    </div>
  </div>

  <div v-highlight v-html="dblog_output">
 <pre>
            <code>
            </code>
 </pre>
  </div>
</template>
