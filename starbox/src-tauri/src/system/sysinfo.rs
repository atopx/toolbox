use std::process::{Command, Stdio};

use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct SystemInfo {
    #[serde(rename = "SPHardwareDataType")]
    pub sphardware_data_types: Vec<SphardwareDataType>,

    #[serde(rename = "SPStorageDataType")]
    pub spstorage_data_types: Vec<SpstorageDataType>,
}

#[derive(Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct SphardwareDataType {
    #[serde(rename = "_name")]
    pub name: String,
    #[serde(rename = "activation_lock_status")]
    pub activation_lock_status: String,
    #[serde(rename = "boot_rom_version")]
    pub boot_rom_version: String,
    #[serde(rename = "chip_type")]
    pub chip_type: String,
    #[serde(rename = "machine_model")]
    pub machine_model: String,
    #[serde(rename = "machine_name")]
    pub machine_name: String,
    #[serde(rename = "number_processors")]
    pub number_processors: String,
    #[serde(rename = "os_loader_version")]
    pub os_loader_version: String,
    #[serde(rename = "physical_memory")]
    pub physical_memory: String,
    #[serde(rename = "platform_UUID")]
    pub platform_uuid: String,
    #[serde(rename = "provisioning_UDID")]
    pub provisioning_udid: String,
    #[serde(rename = "serial_number")]
    pub serial_number: String,
}

#[derive(Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct SpstorageDataType {
    #[serde(rename = "_name")]
    pub name: String,
    #[serde(rename = "bsd_name")]
    pub bsd_name: String,
    #[serde(rename = "file_system")]
    pub file_system: String,
    #[serde(rename = "free_space_in_bytes")]
    pub free_space_in_bytes: i64,
    #[serde(rename = "ignore_ownership")]
    pub ignore_ownership: String,
    #[serde(rename = "mount_point")]
    pub mount_point: String,
    #[serde(rename = "size_in_bytes")]
    pub size_in_bytes: i64,
    #[serde(rename = "volume_uuid")]
    pub volume_uuid: String,
    pub writable: String,
}



impl SystemInfo {
    pub fn new() -> Self {
        let output = Command::new("system_profiler")
            .arg("SPHardwareDataType")
            .arg("SPStorageDataType")
            .arg("-json")
            .stdout(Stdio::piped())
            .output();

        match output {
            Ok(value) => {
                let r = String::from_utf8(value.stdout).unwrap();
                let info: SystemInfo = serde_json::from_str(r.as_str()).unwrap();
                return info;
            }
            Err(_) => {
                return SystemInfo::default();
            }
        };
    }

    fn default() -> Self {
        return SystemInfo {
            sphardware_data_types: Vec::new(),
            spstorage_data_types: Vec::new(),
        };
    }
}

#[derive(Serialize, Deserialize)]
pub struct SystemInfoResponse {
    hardward: SphardwareDataType,
    storage: SpstorageDataType,
}

impl SystemInfoResponse {
    fn new(hardward: &SphardwareDataType, storage: &SpstorageDataType) -> Self {
        return SystemInfoResponse {
            hardward: SphardwareDataType {
                name: hardward.name.to_string(),
                activation_lock_status: hardward.activation_lock_status.to_string(),
                boot_rom_version: hardward.boot_rom_version.to_string(),
                chip_type: hardward.chip_type.to_string(),
                machine_model: hardward.machine_model.to_string(),
                machine_name: hardward.machine_name.to_string(),
                number_processors: hardward.number_processors.to_string(),
                os_loader_version: hardward.os_loader_version.to_string(),
                physical_memory: hardward.physical_memory.to_string(),
                platform_uuid: hardward.platform_uuid.to_string(),
                provisioning_udid: hardward.provisioning_udid.to_string(),
                serial_number: hardward.serial_number.to_string(),
            },
            storage: SpstorageDataType {
                name: storage.name.to_string(),
                bsd_name: storage.bsd_name.to_string(),
                file_system: storage.file_system.to_string(),
                free_space_in_bytes: storage.free_space_in_bytes.to_be(),
                ignore_ownership: storage.ignore_ownership.to_string(),
                mount_point: storage.mount_point.to_string(),
                size_in_bytes: storage.size_in_bytes.to_be(),
                volume_uuid: storage.volume_uuid.to_string(),
                writable: storage.writable.to_string(),
            },
        };
    }
}

#[tauri::command]
pub fn get_info() -> SystemInfoResponse {
    let info = SystemInfo::new();
    let hardward = info.sphardware_data_types.last().unwrap();
    let storage = info.spstorage_data_types.last().unwrap();
    return SystemInfoResponse::new(hardward, storage);
}
