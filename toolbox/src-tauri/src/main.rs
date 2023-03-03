#![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]

use reqwest;

#[tauri::command]
fn greet(name: &str) -> String {
    format!("Hello, {}! You've been greeted from Rust!", name)
}


#[tauri::command]
fn novel_fetch(link: &str) -> Result<String, String> {
    let client = reqwest::blocking::Client::new();
    let resp = client.get(link)
        .header("user-agent", "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)")
        .send();

    return match resp {
        Ok(data) => {
            Ok(data.text().unwrap())
        }
        Err(e) => { Err(e.to_string()) }
    };
}

fn main() {
    tauri::Builder::default()
        .invoke_handler(tauri::generate_handler![
            greet,
            novel_fetch,
        ])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
