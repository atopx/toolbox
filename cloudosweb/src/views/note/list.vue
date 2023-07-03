<script lang="ts" setup>
import { reactive, ref, watch } from "vue"
import { noteListApi, noteDeleteApi, noteInfoApi, noteTopicsApi } from "@/api/note"
import { noteListRequest } from "@/api/note/types"
import { type FormInstance, ElMessage, ElMessageBox } from "element-plus"
import { Search, Refresh, CirclePlus, RefreshRight } from "@element-plus/icons-vue"
import { usePagination } from "@/hooks/usePagination"
import { Note } from "@/api/note/types"
import router from "@/router"
import { OptionIface } from "types/api"

defineOptions({
    name: "NoteList"
})

const loading = ref<boolean>(false)
const previewDialog = ref({
    visible: false,
    title: "",
    content: ""
})

const topicOptions = ref<OptionIface[]>([])

noteTopicsApi().then((resp) => {
    resp.data.forEach((v) => {
        topicOptions.value.push({ value: v, label: v })
    })
})

const { paginationData, handleCurrentChange, handleSizeChange } = usePagination()
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

const tableData = ref<Note[]>([])
const searchFormRef = ref<FormInstance | null>(null)
const searchData = reactive({
    keyword: "",
    topic: "",
    timeRange: ""
})

// 刷新列表
const handleRefresh = () => {
    loading.value = true

    const params: noteListRequest = {
        pager: {
            index: paginationData.currentPage,
            size: paginationData.pageSize,
            count: paginationData.total,
            disable: false
        },
        keyword: searchData.keyword,
        topic: searchData.topic
    }

    if (searchData.timeRange.length === 2) {
        params.updateTimeRange = {
            left: (searchData.timeRange[0] as unknown as Date).getTime() / 1000,
            right: (searchData.timeRange[1] as unknown as Date).getTime() / 1000 + 86400
        }
    }

    noteListApi(params)
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

// 搜索
const handleSearch = () => {
    if (paginationData.currentPage === 1) {
        handleRefresh()
    }
    paginationData.currentPage = 1
}

// 重置搜索条件
const resetSearch = () => {
    searchFormRef.value?.resetFields()
    if (paginationData.currentPage === 1) {
        handleRefresh()
    }
    paginationData.currentPage = 1
}

// 预览
const contentPreview = (row: Note) => {
    noteInfoApi(row.id)
        .then((resp) => {
            previewDialog.value.visible = true
            previewDialog.value.title = resp.data.title
            previewDialog.value.content = resp.data.content
        })
        .finally(() => {
            loading.value = false
        })
}

/** 监听分页参数的变化 */
watch([() => paginationData.currentPage, () => paginationData.pageSize], handleRefresh, { immediate: true })
</script>

<template>
    <div class="app-container">
        <el-card v-loading="loading" shadow="never" class="search-wrapper">
            <el-form ref="searchFormRef" :inline="true" :model="searchData">
                <el-form-item prop="keyword" label="关键词">
                    <el-input v-model="searchData.keyword" placeholder="关键词" />
                </el-form-item>
                <el-form-item prop="topic" label="主题">
                    <el-select v-model="searchData.topic" filterable placeholder="选择主题">
                        <el-option
                            v-for="item in topicOptions"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item prop="timeRange" label="时间">
                    <el-date-picker
                        v-model="searchData.timeRange"
                        type="daterange"
                        range-separator="至"
                        start-placeholder="开始日期"
                        end-placeholder="结束日期"
                    />
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
                            <el-tag>{{ scope.row.topic }}</el-tag>
                        </template>
                    </el-table-column>
                    <el-table-column prop="createTime" label="创建时间" align="center" />
                    <el-table-column prop="updateTime" label="更新时间" align="center" />
                    <el-table-column fixed="right" label="操作" width="200" align="center">
                        <template #default="scope">
                            <el-button type="primary" text bg size="small" @click="contentPreview(scope.row)"
                                >预览</el-button
                            >
                            <el-button type="primary" text bg size="small" @click="handleSave(scope.row)"
                                >修改</el-button
                            >
                            <el-button type="danger" text bg size="small" @click="handleDelete(scope.row)"
                                >删除</el-button
                            >
                        </template>
                    </el-table-column>
                </el-table>
            </div>
            <el-dialog v-model="previewDialog.visible" :title="previewDialog.title" width="40%">
                <v-md-preview :text="previewDialog.content" />
                <template #footer>
                    <span class="dialog-footer">
                        <el-button type="primary" @click="previewDialog.visible = false"> 关闭 </el-button>
                    </span>
                </template>
            </el-dialog>
            <div class="pager-wrapper">
                <el-pagination
                    background
                    :layout="paginationData.layout"
                    :page-sizes="paginationData.pageSizes"
                    :total="paginationData.total"
                    :page-size="paginationData.pageSize"
                    :currentPage="paginationData.currentPage"
                    @size-change="handleSizeChange"
                    @current-change="handleCurrentChange"
                />
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
