<script setup lang="ts">
import {ref} from "vue";
import {invoke} from "@tauri-apps/api/tauri";
import {Aim} from '@element-plus/icons-vue';
import {load} from 'cheerio';


const message = ref({
	name: "",
	author: "",
	intro: "",
	lasted: "",
	source: "",
	chapters: new Array<{ index: number, title: string, source: string }>(),
});

const loading = ref(true)
const display = ref()

async function clean() {
	message.value.name = ""
	message.value.author = ""
	message.value.intro = ""
	message.value.lasted = ""
	message.value.chapters = new Array<{ index: number, title: string, source: string }>();
	display.value.style = "display: none";
}

const error = ref(
	{
		status: false,
		message: ''
	}
)


async function fetch() {
	await clean();
	display.value.style = "display: block";
	await invoke("novel_fetch", {link: message.value.source}).then((html) => {
		const $ = load(html as string);
		let info = $('div#info');
		message.value.name = info.children("h1").text();
		message.value.author = info.children("p:nth-child(2)").text().split("：")[1];
		message.value.lasted = info.children("p:nth-child(4)").text().split("：")[1];
		message.value.intro = $('div#intro').text();

		let chapters = new Array<string>();
		let chapterMap = new Map<string, string>();

		let origin = new URL(message.value.source).origin;

		$('div#list dd a').each(function (index, element) {
			let source = element.attribs["href"]
			// 章节去重
			if (!chapterMap.has(source)) {
				source = origin + source
				chapters.push(source);
				chapterMap.set(source, $(element).text());
			}
		})

		// 章节排序
		chapters.sort().forEach(function (source, index) {
			message.value.chapters.push({
				index: index + 1,
				title: chapterMap.get(source)!,
				source: source,
			})
		})
		loading.value = false
	}).catch((err) => {
		display.value.style = "display: none";
		error.value.message = err;
		error.value.status = true;
	});
}

</script>

<template>
	<el-input v-model="message.source" placeholder="目录地址" class="input-with-select" size="large">
		<template #append>
			<el-button :icon="Aim" @click="fetch()"></el-button>
		</template>
	</el-input>

	<el-dialog
		v-model="error.status"
		title="ERROR"
		width="30%"
		align-center
	>
		<span class="error-message"> {{ error.message }} </span>
		<template #footer>
      <span class="dialog-footer">
        <el-button @click="error.status = false">Cancel</el-button>
      </span>
		</template>
	</el-dialog>

	<div ref="display" id="info" v-loading="loading">
		<el-card shadow="hover" id="head">
			<el-row>
				<el-col :span="22">
					<span style="font-size: 20px; color: bisque">《{{ message.name }}》</span>
					<span style="font-size: 8px; margin-left:5px">{{ message.author }} - {{ message.lasted }}</span>
				</el-col>
				<el-col :span="1">
					<el-button type="primary" style="width: 80px; height: 30px">Download</el-button>
				</el-col>
			</el-row>

		</el-card>

		<el-table :data="message.chapters" id="table" height="550" stripe>
			<el-table-column prop="index" label="序号" width="120px" sortable/>
			<el-table-column prop="title" label="章节"/>
			<el-table-column prop="source" label="地址"/>
		</el-table>
	</div>
</template>


<style scoped>

#info {
	display: none;
}

#head {
	padding: 2px;
	margin: 20px 0 20px 0;
	background-color: darkslategray;
	color: azure;
}

.el-alert {
	margin: 20px 0 0;
	padding-left: 8px;
	display: inline-block;
}

.error-message {
	background-color: var(--el-alert-bg-color);
	color: var(--el-color-error);
}

#table {
	width: 100%;
}

</style>