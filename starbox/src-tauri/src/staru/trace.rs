use ssh2::Session;
use std::path::Path;
use std::{
    io::{BufRead, BufReader},
    net::TcpStream,
};

use serde_json::Value;

#[derive(serde::Serialize)]
pub enum LogType {
    UNKNOWN,
    DB,
    TRACK,
    RUNTIME,
}

impl LogType {
    pub fn to_str(&self) -> &str {
        return match self {
            Self::UNKNOWN => "UNKNOWN",
            Self::DB => "DB",
            Self::TRACK => "TRACK",
            Self::RUNTIME => "RUNTIME",
        };
    }

    pub fn to_string(&self) -> String {
        return String::from(self.to_str());
    }
}

#[derive(serde::Serialize)]
pub struct TraceLog {
    timestamp: u64,
    timestr: String,
    service: String,
    r#type: LogType,
    level: String,
    infc: String,
    message: String,
}

#[tauri::command]
pub fn tracelog(id: &str) -> Vec<TraceLog> {
    let mut result: Vec<TraceLog> = Vec::new();
    result.push(TraceLog {
        timestamp: 1,
        timestr: format!("2023-01-01 16:00:00"),
        service: format!("Service1"),
        r#type: LogType::UNKNOWN,
        level: format!("Level1"),
        infc: format!("Infc1"),
        message: format!("message: {id}"),
    });
    result.push(TraceLog {
        timestamp: 2,
        timestr: format!("2023-01-01 16:00:10"),
        service: format!("Service2"),
        r#type: LogType::DB,
        level: format!("Level2"),
        infc: format!("Infc13"),
        message: format!("message: {id}"),
    });
    result.push(TraceLog {
        timestamp: 3,
        timestr: format!("2023-01-01 16:00:20"),
        service: format!("Service2"),
        r#type:LogType::TRACK,
        level: format!("Level2"),
        infc: format!("Infc13"),
        message: format!("message: {id}"),
    });
    result.push(TraceLog {
        timestamp: 4,
        timestr: format!("2023-01-01 16:00:30"),
        service: format!("Service2"),
        r#type: LogType::RUNTIME,
        level: format!("Level2"),
        infc: format!("Infc13"),
        message: format!("message: {id}"),
    });
    result.push(TraceLog {
        timestamp: 5,
        timestr: format!("2023-01-01 16:00:50"),
        service: format!("Service2"),
        r#type: LogType::RUNTIME,
        level: format!("Level2"),
        infc: format!("Infc13"),
        message: format!("message: {id}"),
    });
    result.push(TraceLog {
        timestamp: 6,
        timestr: format!("2023-01-01 16:00:60"),
        service: format!("Service2"),
        r#type: LogType::RUNTIME,
        level: format!("Level2"),
        infc: format!("Infc13"),
        message: format!("message: {id}"),
    });
    result.push(TraceLog {
        timestamp: 7,
        timestr: format!("2023-01-01 16:00:70"),
        service: format!("Service2"),
        r#type: LogType::RUNTIME,
        level: format!("Level2"),
        infc: format!("Infc13"),
        message: format!("message: {id}"),
    });
    result.push(TraceLog {
        timestamp: 8,
        timestr: format!("2023-01-01 16:00:80"),
        service: format!("Service2"),
        r#type: LogType::RUNTIME,
        level: format!("Level2"),
        infc: format!("Infc13"),
        message: format!("message: {id}"),
    });
    return result;
}

fn fetch_log(trace_id: &str) {
    let tcp = TcpStream::connect("10.0.0.211:22").unwrap();
    let mut sess = Session::new().unwrap();
    sess.set_tcp_stream(tcp);
    sess.handshake().unwrap();
    let privatekey = Path::new("/Users/atopx/.ssh/id_rsa");
    sess.userauth_pubkey_file("staru", None, &privatekey, None)
        .unwrap();
    let mut channel = sess.channel_session().unwrap();
    channel
        .exec(format!("/bin/grep -h {trace_id} work/star_union*/logs/*.log").as_str())
        .unwrap();
    let reader = BufReader::new(&mut channel);

    for line in reader.lines() {
        match line {
            Ok(data) => {
                let value: Value =  serde_json::from_str(data.as_str()).unwrap();
                println!("{:?}", value);
            }
            Err(_) => {}
        }
    }
    channel.wait_close().unwrap();
}

#[cfg(test)]
mod tests {
    use super::fetch_log;

    #[test]
    fn test_fetch_log() {
        fetch_log("1678762915168580224");
    }
}
