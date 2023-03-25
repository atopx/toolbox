use chrono::Local;

pub fn current_timestamp() -> i64 {
    return Local::now().timestamp();
}