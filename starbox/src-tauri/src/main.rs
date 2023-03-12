#![cfg_attr(
    all(not(debug_assertions), target_os = "windows"),
    windows_subsystem = "windows"
)]

mod staru;
mod system;
mod novel;
mod config;

use std::{thread, time};

#[tauri::command]
async fn event_sleep(t: u64) {
    let ten_millis = time::Duration::from_millis(t);
    thread::sleep(ten_millis);
}

fn main() {
    tauri::Builder::default()
        .invoke_handler(tauri::generate_handler![
            staru::trace::tracelog,
            staru::permission::permission,
            staru::dblog::dblog,
            system::sysinfo::get_info,

            event_sleep,
            novel::downer::novel_fetch,
            novel::downer::novel_compose,
            novel::downer::novel_save,
            novel::downer::novel_clean,
            // system::portscan::scanport_event,

            ])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
