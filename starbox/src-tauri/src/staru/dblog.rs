#[tauri::command]
pub fn dblog(value: &str) -> String {
    let temp = value.split("] ").last().unwrap();
    let mut result = String::from(temp.split("\",").next().unwrap());
    result = result.replace("\\n", " ");
    result = result.replace("\\t", " ");
    if !result.ends_with(";") {
        result.push(';');
    }
    return result;
}

