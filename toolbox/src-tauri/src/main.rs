#![cfg_attr(
    all(not(debug_assertions), target_os = "windows"),
    windows_subsystem = "windows"
    )]
    
    #[tauri::command]
    fn dblog(value: &str) -> String {
        let temp = value.split("] ").last().unwrap();

        let mut result = String::from(temp.split("\",").next().unwrap());
        result = result.replace("\\n", " ");
        result = result.replace("\\t", " ");
        if !result.ends_with(";") {
            result.push(';');
        }
        return result;
    }
    
    fn main() {
        tauri::Builder::default()
            .invoke_handler(tauri::generate_handler![dblog])
            .run(tauri::generate_context!())
            .expect("error while running tauri application");
    }