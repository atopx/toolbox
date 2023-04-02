#[tauri::command]
pub fn permission(name: &str, path: &str, code: &str, parent: &str) -> String {
    let mut result: Vec<String> = Vec::new();
    result.push(format!("insert ignore into star_union.admin_task (path, anonymous, create_time, update_time) 
        VALUES ('{path}', 0, unix_timestamp(), unix_timestamp());"));
    result.push(format!("insert ignore into admin_operation (name, `desc`, parent_id, level, create_time, update_time) 
        select '{name}', '{code}', aop.id, aop.level+1, unix_timestamp(), unix_timestamp() from admin_operation aop 
        where `desc` = '{parent}';"));
    result.push(format!("insert ignore into admin_operation_task_relation (operation_id, task_id, create_time, update_time) 
        select o.id, t.id, unix_timestamp(), unix_timestamp() from admin_task t 
        join admin_operation o where o.`desc` = '{code}' and t.path = '{path}';"));
    result.push(format!("insert ignore into admin_operation_task_relation (operation_id, task_id, create_time, update_time) 
        select o.id, t.id, unix_timestamp(), unix_timestamp() from admin_task t 
        join admin_operation o where o.`desc` = 'admin' and t.path = '{path}';"));
    return result.join("\n\n");
}
