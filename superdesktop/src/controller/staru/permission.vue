<script setup lang="ts">
import { reactive, ref } from 'vue'
import type { FormInstance } from 'element-plus'
import { invoke } from '@tauri-apps/api/tauri';

const formRef = ref<FormInstance>()
const formData = reactive({
  name: '',
  path: '',
  code: '',
  parent: '',
})

const display = ref();
const result = ref('');

async function permission() {
  result.value = await invoke("permission", {
    name: formData.name,
    path: formData.path,
    code: formData.code,
    parent: formData.parent,
  });
  display.value.style = "display: block";
}

const submit = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate((valid) => {
    if (valid) {
      permission()
    } else {
      console.log('error submit!')
      return false
    }
  })
}

const reset = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.resetFields()
}

</script>
<template>
  <el-main>
    <el-form ref="formRef" :model="formData" status-icon label-width="120px">
      <el-form-item label="Name" prop="name" required>
        <el-input v-model="formData.name" />
      </el-form-item>

      <el-form-item label="Path" prop="path" required>
        <el-input v-model="formData.path" />
      </el-form-item>

      <el-form-item label="Code" prop="code" required>
        <el-input v-model="formData.code" />
      </el-form-item>

      <el-form-item label="ParentCode" prop="parent" required>
        <el-input v-model="formData.parent" />
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="submit(formRef)">生成</el-button>
        <el-button @click="reset(formRef)">清空</el-button>
      </el-form-item>
    </el-form>


    <div ref="display" class="result">
      <highlightjs language='sql' :code="result" />
    </div>


  </el-main>
</template>
<style>
.result {
  display: none;
  width: 990px;
  max-height: 300px;
  margin-top: 10px;
}
</style>