<script lang="ts" setup>
import { nextTick, reactive, ref, watch } from "vue"
import { noteListApi, noteDeleteApi } from "@/api/note"
import { type FormInstance, ElMessage, ElMessageBox } from "element-plus"
import { Search, Refresh, CirclePlus, RefreshRight } from "@element-plus/icons-vue"
import { usePagination } from "@/hooks/usePagination"
import { Note } from "@/api/note/types"
import router from "@/router"
import { tr } from "element-plus/es/locale"

defineOptions({
    name: "NoteList"
})

const loading = ref<boolean>(false)
const { paginationData, handleCurrentChange, handleSizeChange } = usePagination()
//#region 删
const handleDelete = (row: Note) => {
    ElMessageBox.confirm(`< ${row.title} >`, "删除笔记", {
        confirmButtonText: "删除",
        cancelButtonText: "取消",
        type: "warning"
    }).then(() => {
        noteDeleteApi(row.id).then(() => {
            ElMessage.success("删除成功")
            handleRefresh()
        })
    })
}
//#endregion

//#region save
const currentUpdateId = ref<number>(0)
const handleSave = (row: Note | null) => {
    if (row != null) {
        currentUpdateId.value = row.id
    }
    router.push({
        path: "/note/editor",
        query: { id: currentUpdateId.value }
    })
}
//#endregion

//#region 查
const tableData = ref<Note[]>([])
const searchFormRef = ref<FormInstance | null>(null)
const searchData = reactive({ keyword: "" })

const handleRefresh = () => {
    loading.value = true

    noteListApi({
        pager: {
            index: paginationData.currentPage,
            size: paginationData.pageSize,
            disabled: false,
            count: paginationData.total
        },
        filter: {
            keyword: searchData.keyword
        }
    })
        .then((resp) => {
            paginationData.total = resp.data.pager.count
            tableData.value = resp.data.list
        })
        .catch(() => {
            tableData.value = []
        })
        .finally(() => {
            loading.value = false
        })
}
const handleSearch = () => {
    if (paginationData.currentPage === 1) {
        handleRefresh()
    }
    paginationData.currentPage = 1
}
const resetSearch = () => {
    searchFormRef.value?.resetFields()
    if (paginationData.currentPage === 1) {
        handleRefresh()
    }
    paginationData.currentPage = 1
}
//#endregion

//#region label
const topicInputRef = ref({
    visible: false,
    text: ''
})
const labelInputRef = ref({
    visible: false,
    text: ''
})


const handleDelLebel = (id: number) => {
    // TODO delete note label/topic
}

const handleNewTag = (isTopic: boolean) => {
    if (isTopic) {
        topicInputRef.value.visible = false
        if (topicInputRef.value.text) {
            // TODO create note topic
            topicInputRef.value.text = ''
        }
    } else {
        labelInputRef.value.visible = false
        if (labelInputRef.value.text) {
            // TODO create note label
            labelInputRef.value.text = ''
        }
    }
}

const handleShowTagInput = (isTopic: boolean) => {
    if (isTopic) {
        topicInputRef.value.visible = true
        nextTick(() => {
            // topicInputRef.value!.text!.focus()
        })
    } else {
        labelInputRef.value.visible = true
        nextTick(() => {
            // labelInputRef.value!.text!.focus()
        })
    }
}
//#endregion


/** 监听分页参数的变化 */
watch([() => paginationData.currentPage, () => paginationData.pageSize], handleRefresh, { immediate: true })
</script>

<template>
    <div class="app-container">
        <el-card v-loading="loading" shadow="never" class="search-wrapper">
            <el-form ref="searchFormRef" :inline="true" :model="searchData">
                <el-form-item prop="keyword" label="关键词">
                    <el-input v-model="searchData.keyword" placeholder="请输入" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" :icon="Search" @click="handleSearch">查询</el-button>
                    <el-button :icon="Refresh" @click="resetSearch">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card v-loading="loading" shadow="never">
            <div class="toolbar-wrapper">
                <div>
                    <el-button type="primary" :icon="CirclePlus" @click="handleSave(null)">新增笔记</el-button>
                </div>
                <div>
                    <el-tooltip content="刷新表格">
                        <el-button type="primary" :icon="RefreshRight" circle @click="handleRefresh" />
                    </el-tooltip>
                </div>
            </div>
            <div class="table-wrapper">
                <el-table :data="tableData">
                    <el-table-column width="50" align="center" />
                    <el-table-column prop="id" label="ID" width="50" align="center" />
                    <el-table-column prop="title" label="标题" align="center" />
                    <el-table-column label="主题" align="center">
                        <template #default="scope">
                            <el-tag v-for="item in scope.row.topics" effect="light" closable @close="handleDelLebel(item)">
                                {{ item }}
                            </el-tag>
                            <el-input v-if="topicInputRef.visible" ref="topicInputRef.text" v-model="topicInputRef.text" class="ml-1 w-20"
                                size="small" @keyup.enter="handleNewTag(true)" @blur="handleNewTag(true)" />
                            <el-button v-else class="button-new-tag ml-1" size="small" @click="handleShowTagInput(true)">+
                                New</el-button>
                        </template>

                    </el-table-column>
                    <el-table-column label="标签" align="center">
                        <template #default="scope">
                            <el-tag v-for="item in scope.row.labels" effect="light" closable @close="handleDelLebel(item)">
                                {{ item }}
                            </el-tag>
                            <el-input v-if="labelInputRef.visible" ref="labelInputRef.text" v-model="labelInputRef.text" class="ml-1 w-20"
                                size="small" @keyup.enter="handleNewTag(false)" @blur="handleNewTag(false)" />
                            <el-button v-else class="button-new-tag ml-1" size="small" @click="handleShowTagInput(false)">+
                                New</el-button>
                        </template>
                    </el-table-column>
                    <el-table-column prop="create_time" label="创建时间" align="center" />
                    <el-table-column prop="update_time" label="更新时间" align="center" />
                    <el-table-column fixed="right" label="操作" width="150" align="center">
                        <template #default="scope">
                            <el-button type="primary" text bg size="small" @click="handleSave(scope.row)">修改</el-button>
                            <el-button type="danger" text bg size="small" @click="handleDelete(scope.row)">删除</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </div>
            <div class="pager-wrapper">
                <el-pagination background :layout="paginationData.layout" :page-sizes="paginationData.pageSizes"
                    :total="paginationData.total" :page-size="paginationData.pageSize"
                    :currentPage="paginationData.currentPage" @size-change="handleSizeChange"
                    @current-change="handleCurrentChange" />
            </div>
        </el-card>
    </div>
</template>

<style lang="scss" scoped>
.search-wrapper {
    margin-bottom: 20px;

    :deep(.el-card__body) {
        padding-bottom: 2px;
    }
}

.toolbar-wrapper {
    display: flex;
    justify-content: space-between;
    margin-bottom: 20px;
}

.table-wrapper {
    margin-bottom: 20px;
}

.pager-wrapper {
    display: flex;
    justify-content: flex-end;
}
</style>
