#![cfg_attr(
all(not(debug_assertions), target_os = "windows"),
windows_subsystem = "windows"
)]

#[tauri::command]
fn dblog(value: &str) -> String {
    let mut result = value.split("] ").last().unwrap();
    result = result.split("\",").next().unwrap();
    let mut sql = String::from(result);
    sql.push(';');
    return sql;
}

fn main() {
    tauri::Builder::default()
        .invoke_handler(tauri::generate_handler![dblog])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
