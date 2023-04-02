<script lang="ts" setup>

import { ref } from "vue";
import { invoke } from "@tauri-apps/api/tauri";

const result = ref({
    hardward: {
        boot_rom_version: "8419.60.44",
        chip_type: "Apple M1 Pro",
        machine_model: "MacBookPro18,3",
        machine_name: "MacBook Pro",
        number_processors: "proc 10:8:2",
        os_loader_version: "7459.141.1",
        physical_memory: "32 GB",
        platform_UUID: "73A8052C-27F9-5B04-85D3-CFAF48E53698",
        provisioning_UDID: "00006000-001824CC14A3801E",
        serial_number: "RQC232YY9P"
    },
    storage: {
        _name: "Macintosh HD",
        bsd_name: "disk3s3s1",
        file_system: "APFS",
        free_space_in_bytes: 31745751126114304,
        ignore_ownership: "no",
        mount_point: "/",
        size_in_bytes: 40675981278576640,
        volume_uuid: "9295DBD6-9487-48F7-893C-B63CC59AAADB",
        writable: "no"
    }
});




async function refresh() {
    result.value = await invoke("get_info");
}

refresh()

</script>

<template>
        <el-descriptions title="系统" :column="2" border class="info">
        <el-descriptions-item label="设备型号" label-align="center" align="center" width="240px"> {{ result.hardward.chip_type
        }} </el-descriptions-item>

        <el-descriptions-item label="基带版本" label-align="center" align="center">{{ result.hardward.boot_rom_version
        }}</el-descriptions-item>

        <el-descriptions-item label="CPU核心" label-align="center" align="center"> {{ result.hardward.number_processors }}
        </el-descriptions-item>

        <el-descriptions-item label="内存空间" label-align="center" align="center"> {{ result.hardward.physical_memory }}
        </el-descriptions-item>

        <el-descriptions-item label="序列号" label-align="center" align="center">{{ result.hardward.serial_number
        }}</el-descriptions-item>

        <el-descriptions-item label="平台ID" label-align="center" align="center">{{ result.hardward.platform_UUID
        }}</el-descriptions-item>
    </el-descriptions>

    <el-descriptions title="磁盘" :column="2" border class="info">
        <el-descriptions-item label="磁盘名称" label-align="center" align="center" width="240px"> {{ result.storage._name }}
        </el-descriptions-item>

        <el-descriptions-item label="文件系统" label-align="center" align="center"> {{ result.storage.file_system }}
        </el-descriptions-item>

        <el-descriptions-item label="磁盘容量" label-align="center" align="center">{{ result.storage.size_in_bytes
        }}</el-descriptions-item>

        <el-descriptions-item label="可用容量" label-align="center" align="center">{{ result.storage.free_space_in_bytes
        }}</el-descriptions-item>

        <el-descriptions-item label="挂载路径" label-align="center"
            align="center">{{ result.storage.mount_point }}</el-descriptions-item>

        <el-descriptions-item label="磁盘ID" label-align="center" align="center"> {{ result.storage.volume_uuid }}
        </el-descriptions-item>
    </el-descriptions>

    <el-descriptions title="网络" :column="2" border class="info">
        <el-descriptions-item label="磁盘名称" label-align="center" align="center" width="240px"> {{ result.storage._name }}
        </el-descriptions-item>

        <el-descriptions-item label="文件系统" label-align="center" align="center"> {{ result.storage.file_system }}
        </el-descriptions-item>

        <el-descriptions-item label="磁盘容量" label-align="center" align="center">{{ result.storage.size_in_bytes
        }}</el-descriptions-item>

        <el-descriptions-item label="可用容量" label-align="center" align="center">{{ result.storage.free_space_in_bytes
        }}</el-descriptions-item>

        <el-descriptions-item label="挂载路径" label-align="center"
            align="center">{{ result.storage.mount_point }}</el-descriptions-item>

        <el-descriptions-item label="磁盘ID" label-align="center" align="center"> {{ result.storage.volume_uuid }}
        </el-descriptions-item>
    </el-descriptions>
    
</template>


<style scoped>
.info {
    margin-top: 15px;
    height: 200px;
}


</style>