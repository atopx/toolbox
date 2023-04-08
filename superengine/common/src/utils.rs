use chrono::Local;

pub fn current_timestamp() -> i64 {
    Local::now().timestamp()
}