<script setup lang="ts">
import {ref} from "vue";
import {invoke} from "@tauri-apps/api/tauri";

const greetMsg = ref("");
const name = ref("");

const downloader = ref({
	status: false,
	percentage: 0,
	message: '准备中...',
});

const colors = [
	{ color: '#f56c6c', percentage: 20 },
	{ color: '#e6a23c', percentage: 50 },
	{ color: '#5cb87a', percentage: 80 },
	{ color: '#6f7ad3', percentage: 99 },
	{ color: '#1989fa', percentage: 100 },
]

async function greet() {
	// Learn more about Tauri commands at https://tauri.app/v1/guides/features/command
	greetMsg.value = await invoke("greet", {name: name.value});
}

const Sleep = (ms: number | undefined)=> {
	return new Promise(resolve=>setTimeout(resolve, ms))
}

async function test() {

	downloader.value.status = true

	const Sleep = (ms: number | undefined)=> {
		return new Promise(resolve=>setTimeout(resolve, ms))
	}

	for (let i = 0; i <= 100; i++) {
		downloader.value.message = "loading..." + i;
		downloader.value.percentage = i
		await Sleep(100);
	}

}

</script>

<template>

	<el-button text @click="test()">
		click to open the Dialog
	</el-button>

	<el-dialog
		v-model="downloader.status"
		title="下载器"
		width="30%"
		align-center
	>

		<el-progress type="dashboard" :percentage="downloader.percentage" stroke-width="8" width="320" :color="colors">
			<template #default="{ percentage }">
				<span class="percentage-value">{{ downloader.percentage }}%</span>
				<span class="percentage-label">{{ downloader.message }}</span>
			</template>
		</el-progress>

		<template #footer>
      <span class="dialog-footer">
        <el-button @click="downloader.status = false">取消</el-button>
      </span>
		</template>
	</el-dialog>
</template>


<style scoped>

.percentage-value {
	display: block;
	margin-top: 10px;
	font-size: 50px;
}
.percentage-label {
	display: block;
	margin-top: 10px;
	font-size: 12px;
}
.demo-progress .el-progress--line {
	margin-bottom: 15px;
	width: 350px;
}
.demo-progress .el-progress--circle {
	margin-right: 15px;
}

</style>