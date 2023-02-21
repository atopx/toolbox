<script setup>
import Dblog from "./components/dblog.vue";
import Track from "./components/track.vue";
import linuxcmds from "./components/linuxcmds.vue";

import { ref } from 'vue';

const currentTableId = ref("1")
const currentOpenedSide = ref("1")
const currentSelectedSide = ref("1")


const changeSelectedKey = (val) => {
  currentSelectedSide.value = val;
}

const changeOpenKeys = (val) => {
  currentOpenedSide.value = val;
}


</script>

<template>
  <lay-layout>
    <!-- left -->

    <lay-side class="layside">
          <lay-menu :selected-key="currentSelectedSide" @change-selected-Key="changeSelectedKey" class="sidemenu"
            @change-open-keys="changeOpenKeys" v-model:openKeys="currentOpenedSide" :tree="true">
            <lay-menu-item id="1">日志处理</lay-menu-item>
            <lay-menu-item id="2">首页2</lay-menu-item>
            <lay-menu-item id="3">首页3</lay-menu-item>
            <lay-menu-item id="4">首页4</lay-menu-item>
            <lay-menu-item id="5">首页5</lay-menu-item>
            <lay-menu-item id="6">首页6</lay-menu-item>
            <lay-menu-item id="7">首页7</lay-menu-item>
            <lay-menu-item id="8">首页8</lay-menu-item>
            <lay-menu-item id="9">首页9</lay-menu-item>
          </lay-menu>
      </lay-side>


    <lay-layout>
      <!-- content -->
      <lay-body>
        
        <lay-tab type="brief" v-model="currentTableId">
          <lay-tab-item title="DBLOG" id="1">
            <div style="padding:20px">
              <Dblog />
            </div>
          </lay-tab-item>


          <lay-tab-item title="TRACK" id="2">
            <div style="padding:20px">
              <Track />
            </div>
          </lay-tab-item>


          <lay-tab-item title="Linux" id="3">
            <div style="padding:20px">
              <linuxcmds />
            </div>
          </lay-tab-item>


          <lay-tab-item title="选项4" id="4">
            <div style="padding:20px">选项4</div>
          </lay-tab-item>


          <lay-tab-item title="选项5" id="5">
            <div style="padding:20px">选项5</div>
          </lay-tab-item>


        </lay-tab>

      </lay-body>


    </lay-layout>
  </lay-layout>
</template>
