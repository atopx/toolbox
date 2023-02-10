#![cfg_attr(
    all(not(debug_assertions), target_os = "windows"),
    windows_subsystem = "windows"
)]

// Learn more about Tauri commands at https://tauri.app/v1/guides/features/command
#[tauri::command]
fn set_password(password: &str) -> String {
    format!("Hello, {}! You've been greeted from Rust!", password)
}

fn main() {
    tauri::Builder::default()
        .invoke_handler(tauri::generate_handler![set_password])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
