#![cfg_attr(
    all(not(debug_assertions), target_os = "windows"),
    windows_subsystem = "windows"
)]

#[tauri::command]
fn dblog(name: &str) -> String {
    println!("INPUT:: {}", name);
    return String::from("select * from order;");
}

fn main() {
    tauri::Builder::default()
        .invoke_handler(tauri::generate_handler![dblog])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
