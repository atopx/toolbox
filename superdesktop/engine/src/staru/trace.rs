use ssh2::Session;
use std::path::Path;
use std::{
    io::{BufRead, BufReader},
    net::TcpStream,
};

use chrono::DateTime;

use serde_json::Value;

#[derive(serde::Serialize, Debug)]
pub struct TraceLog {
    timestamp: i64,
    timestr: String,
    service: String,
    r#type: String,
    level: String,
    infc: String,
    func: String,
    message: String,
    title: String,
}

impl TraceLog {
    fn new(value: Value) -> Self {

        let timestr = match value["time"].as_str() {
            Some(v) => String::from(v),
            None => String::from("null"),
        };
        let service = match value["svr"].as_str() {
            Some(v) => String::from(v),
            None => String::from("null"),
        };
        let level = match value["level"].as_str() {
            Some(v) => String::from(v),
            None => String::from("track"),
        };
        let infc = match value["infc"].as_str() {
            Some(v) => String::from(v),
            None => String::from("null"),
        };
        let func = match value["func"].as_str() {
            Some(v) => String::from(v),
            None => String::from("null"),
        };
        let message = match value["message"].as_str() {
            Some(v) => String::from(v),
            None => String::from("null"),
        };
        let timestamp = match DateTime::parse_from_rfc3339(&timestr) {
            Ok(v) => v.timestamp(),
            Err(_) => 0,
        };


        let r#type: String;

        if level == "null" {
            r#type = String::from("track");
        } else if func.contains("MysqlLogger") {
            r#type = String::from("db");
        } else {
            r#type = String::from("runtime");
        }

        let title = format!("[{type}-{level}] {infc}");

        return Self {
            timestamp,
            timestr,
            service,
            r#type,
            level,
            infc,
            func,
            message,
            title,
        };
    }
}

#[tauri::command]
pub fn tracelog(host: &str, id: &str) -> Vec<TraceLog> {
    let tcp = TcpStream::connect(format!("{host}:22")).unwrap();
    let mut sess = Session::new().unwrap();
    sess.set_tcp_stream(tcp);
    sess.handshake().unwrap();
    let privatekey = Path::new("/Users/atopx/.ssh/id_rsa");
    sess.userauth_pubkey_file("staru", None, &privatekey, None)
        .unwrap();
    let mut channel = sess.channel_session().unwrap();
    channel
        .exec(format!("/bin/grep -h {id} work/star_union*/logs/*.log").as_str())
        .unwrap();
    let reader = BufReader::new(&mut channel);
    let mut result: Vec<TraceLog> = Vec::new();
    for line in reader.lines() {
        match line {
            Ok(data) => {
                let value: Value = serde_json::from_str(data.as_str()).unwrap();
                result.push(TraceLog::new(value));
            }
            Err(_) => {}
        }
    }
    channel.wait_close().unwrap();
    return result;
}

#[cfg(test)]
mod tests {
    use super::tracelog;

    #[test]
    fn test_fetch_log() {
        let result = tracelog("127.0.0.1", "1678849282905208556");
        println!("{:#?}", result);
    }
}
