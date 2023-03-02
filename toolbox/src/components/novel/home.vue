<script setup lang="ts">
import { ref } from "vue";
import { invoke } from "@tauri-apps/api/tauri";
import { Aim } from '@element-plus/icons-vue';
import { load } from 'cheerio';

const message = ref({
  name: "",
  author: "",
  intro: "",
  lasted: "",
  source: "",
  chapters: [],
});



async function fetch() {
  let html: string = await invoke("novel_fetch", { link: message.value.source })
  const $ = load(html);
  let info = $('div#info');
  message.value.name = info.children("h1").text();
  message.value.author = info.children("p:nth-child(2)").text();
  message.value.lasted = info.children("p:nth-child(4)").text();
  message.value.intro = $('div#intro').text();
}

</script>

<template>
  <el-input v-model="message.source" placeholder="目录地址" class="input-with-select" size="large">
    <template #append>
      <el-button :icon="Aim" @click="fetch()"></el-button>
    </template>
  </el-input>

  <el-scrollbar height="600px" style="margin-top: 40px; font-size: 14px;">
    <div class="info">
      <h1>{{ message.name }}</h1>
      <p>{{ message.author }}</p>
      <p>{{ message.lasted }}</p>
    </div>
    <div>{{ message.intro }}</div>
  </el-scrollbar>
</template>


<style scoped>
.tables {
  max-height: 750px;
  padding: 10px;
  border-color: black;
  border: 1px;
}

.info p {
  height: 25px;
  line-height: 25px;
  padding-top: 2px;
  width: 50%;
  margin: auto;
  overflow: hidden;
  float: left;
}
</style>