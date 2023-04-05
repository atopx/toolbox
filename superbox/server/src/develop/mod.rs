use std::str::FromStr;
use strum::{Display, EnumIter, EnumString, IntoEnumIterator};

#[derive(EnumString, Display, EnumIter, PartialEq, Debug)]
pub enum Action {
    JsonToGo,
    JsonToRust,
    JsonToJava,
    JsonToYaml,
    JsonToToml,
    JsonToSQL,
    LogToSQL,
}

impl Action {
    fn deal(&self, src: &str) -> String {
        let r = match self {
            Action::JsonToGo => format!("Unimplemented method {}", self),
            Action::JsonToRust => format!("Unimplemented method {}", self),
            Action::JsonToJava => format!("Unimplemented method {}", self),
            Action::JsonToYaml => format!("Unimplemented method {}", self),
            Action::JsonToToml => format!("Unimplemented method {}", self),
            Action::JsonToSQL => format!("Unimplemented method {}", self),
            Action::LogToSQL => self.log_to_sql(src),
        };
        return r;
    }

    fn log_to_sql(&self, src: &str) -> String {
        let temp = src.split("] ").last().unwrap();
        let mut result = String::from(temp.split("\",").next().unwrap());
        result = result.replace("\\n", " ");
        result = result.replace("\\t", " ");
        if !(result.is_empty() || result.ends_with(";")) {
            result.push(';');
        }
        return result;
    }
}


#[tauri::command]
pub fn list_develop_action() -> Vec<String> {
    return Action::iter().map(|s| s.to_string()).collect();
}

#[tauri::command]
pub fn exec_develop_action(action: &str, src: &str) -> String {
    let result = match Action::from_str(action) {
        Ok(v) => v.deal(src),
        Err(_) => String::from("missing a valid action"),
    };
    return result;
}
