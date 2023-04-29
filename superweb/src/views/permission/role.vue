<script lang="ts" setup>
import { reactive, ref, watch } from "vue"
import { createRoleApi, listRoleApi, updateRoleApi, deleteRoleApi } from "@/api/permissions"
import { type FormInstance, type FormRules, ElMessage, ElMessageBox } from "element-plus"
import { Search, Refresh, CirclePlus, Delete, Download, RefreshRight } from "@element-plus/icons-vue"
import { usePagination } from "@/hooks/usePagination"
import { Role } from "@/api/permissions/types"

defineOptions({
    name: "AccessManage"
})

const loading = ref<boolean>(false)
const { paginationData, handleCurrentChange, handleSizeChange } = usePagination()

//#region 增
const dialogVisible = ref<boolean>(false)
const formRef = ref<FormInstance | null>(null)
const formData = reactive({
    name: "",
})
const formRules: FormRules = reactive({
    name: [{ required: true, trigger: "blur", message: "请输入角色名" }],
})



const handleCreate = () => {
    formRef.value?.validate((valid: boolean) => {
        if (valid) {
            if (currentUpdateId.value > 0) {
                updateRoleApi(currentUpdateId.value, formData.name).then(() => {
                    ElMessage.success("修改成功")
                    dialogVisible.value = false
                    getTableData()
                })
            } else {
                createRoleApi(formData.name).then(() => {
                    ElMessage.success("新增成功")
                    dialogVisible.value = false
                    getTableData()
                })
            }
        } else {
            return false
        }
    })
}
const resetForm = () => {
    currentUpdateId.value = 0
    formData.name = ""
}
//#endregion

//#region 删
const handleDelete = (row: Role) => {
    ElMessageBox.confirm(`正在删除角色：${row.name}，确认删除？`, "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
    }).then(() => {
        deleteRoleApi(row.id).then(() => {
            ElMessage.success("删除成功")
            getTableData()
        })
    })
}
//#endregion

//#region 改
const currentUpdateId = ref<number>(0)
const handleUpdate = (row: Role) => {
    currentUpdateId.value = row.id
    formData.name = row.name
    dialogVisible.value = true
}
//#endregion

//#region 查
const tableData = ref<Role[]>([])
const searchFormRef = ref<FormInstance | null>(null)
const searchData = reactive({
    keyword: "",
})

const roleNatureTagTypes = ["info", "danger", "success"];

const getTableData = () => {
    loading.value = true

    listRoleApi({
        page: {
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
        getTableData()
    }
    paginationData.currentPage = 1
}
const resetSearch = () => {
    searchFormRef.value?.resetFields()
    if (paginationData.currentPage === 1) {
        getTableData()
    }
    paginationData.currentPage = 1
}
const handleRefresh = () => {
    getTableData()
}
//#endregion

/** 监听分页参数的变化 */
watch([() => paginationData.currentPage, () => paginationData.pageSize], getTableData, { immediate: true })
</script>

<template>
    <div class="app-container">
        <el-card v-loading="loading" shadow="never" class="search-wrapper">
            <el-form ref="searchFormRef" :inline="true" :model="searchData">
                <el-form-item prop="keyword" label="角色名">
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
                    <el-button type="primary" :icon="CirclePlus" @click="dialogVisible = true">新增用户</el-button>
                    <el-button type="danger" :icon="Delete">批量删除</el-button>
                </div>
                <div>
                    <el-tooltip content="下载">
                        <el-button type="primary" :icon="Download" circle />
                    </el-tooltip>
                    <el-tooltip content="刷新表格">
                        <el-button type="primary" :icon="RefreshRight" circle @click="handleRefresh" />
                    </el-tooltip>
                </div>
            </div>
            <div class="table-wrapper">
                <el-table :data="tableData">
                    <el-table-column type="selection" width="50" align="center" />
                    <el-table-column prop="id" label="ID" width="50" align="center" />
                    <el-table-column prop="name" label="名称" align="center" />
                    <el-table-column label="类型" align="center">
                        <template #default="scope">
                            <el-tag :type="roleNatureTagTypes[scope.row.nature]">{{ scope.row.nature_desc }}</el-tag>
                        </template>
                    </el-table-column>
                    <el-table-column prop="create_time" label="创建时间" align="center" />
                    <el-table-column prop="update_time" label="更新时间" align="center" />
                    <el-table-column fixed="right" label="操作" width="150" align="center">
                        <template #default="scope">
                            <el-button type="primary" text bg size="small" @click="handleUpdate(scope.row)">修改</el-button>
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
        <!-- 新增/修改 -->
        <el-dialog v-model="dialogVisible" :title="currentUpdateId > 0 ? '修改用户' : '新增用户'" @close="resetForm" width="30%">
            <el-form ref="formRef" :model="formData" :rules="formRules" label-width="100px" label-position="left">
                <el-form-item prop="username" label="角色名">
                    <el-input v-model="formData.name" placeholder="请输入" />
                </el-form-item>
            </el-form>
            <template #footer>
                <el-button @click="dialogVisible = false">取消</el-button>
                <el-button type="primary" @click="handleCreate">确认</el-button>
            </template>
        </el-dialog>
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
