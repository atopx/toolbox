#![cfg_attr(
all(not(debug_assertions), target_os = "windows"),
windows_subsystem = "windows"
)]

#[tauri::command]
fn dblog(value: &str) -> String {
    let mut temp = value.split("] ").last().unwrap();
    temp = result.split("\",").next().unwrap();

    let mut result = String::from(temp);
    result = result.replace("\n", " ");
    result = result.replace("\t", " ");
    result.push(';');

    println!("{}", sql);
    return sql;
}

fn main() {
    tauri::Builder::default()
        .invoke_handler(tauri::generate_handler![dblog])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
