#[derive(serde::Serialize)]
pub struct TraceLog {
    timestamp: u64,
    timestr: String,
    service: String,
    r#type: String,
    level: String,
    infc: String,
    message: String,
}

#[tauri::command]
pub fn tracelog(id: &str) -> Vec<TraceLog> {
    let mut result: Vec<TraceLog> = Vec::new();
    result.push(TraceLog{
        timestamp :1,
        timestr: format!("2023-01-01 16:00:00"),
        service: format!("Service1"),
        r#type: format!("Type1"),
        level: format!("Level1"),
        infc: format!("Infc1"),
        message: format!("message: {id}"),
    });
    result.push(TraceLog{
        timestamp :2,
        timestr: format!("2023-01-01 16:00:10"),
        service: format!("Service2"),
        r#type: format!("Type12"),
        level: format!("Level2"),
        infc: format!("Infc13"),
        message: format!("message: {id}"),
    });
    result.push(TraceLog{
        timestamp :3,
        timestr: format!("2023-01-01 16:00:20"),
        service: format!("Service2"),
        r#type: format!("Type12"),
        level: format!("Level2"),
        infc: format!("Infc13"),
        message: format!("message: {id}"),
    });
    result.push(TraceLog{
        timestamp :4,
        timestr: format!("2023-01-01 16:00:30"),
        service: format!("Service2"),
        r#type: format!("Type12"),
        level: format!("Level2"),
        infc: format!("Infc13"),
        message: format!("message: {id}"),
    });
    result.push(TraceLog{
        timestamp :5,
        timestr: format!("2023-01-01 16:00:50"),
        service: format!("Service2"),
        r#type: format!("Type12"),
        level: format!("Level2"),
        infc: format!("Infc13"),
        message: format!("message: {id}"),
    });
    result.push(TraceLog{
        timestamp :6,
        timestr: format!("2023-01-01 16:00:60"),
        service: format!("Service2"),
        r#type: format!("Type12"),
        level: format!("Level2"),
        infc: format!("Infc13"),
        message: format!("message: {id}"),
    });
    result.push(TraceLog{
        timestamp :7,
        timestr: format!("2023-01-01 16:00:70"),
        service: format!("Service2"),
        r#type: format!("Type12"),
        level: format!("Level2"),
        infc: format!("Infc13"),
        message: format!("message: {id}"),
    });
    result.push(TraceLog{
        timestamp :8,
        timestr: format!("2023-01-01 16:00:80"),
        service: format!("Service2"),
        r#type: format!("Type12"),
        level: format!("Level2"),
        infc: format!("Infc13"),
        message: format!("message: {id}"),
    });
    return result;
}
