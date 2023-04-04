use clap::{Error, Parser};

#[derive(Parser, Debug)]
#[command(name = "curl")]
pub struct Curl {
    url: String,

    #[arg(short = 'H', long = "header", action = clap::ArgAction::Append)]
    header: Vec<String>,

    #[arg(short = 'X', long = "method")]
    method: Option<String>,

    #[arg(short = 'd', long = "data-raw")]
    data: Option<String>,

    #[arg(short = 'u', long = "user")]
    auth: Option<String>,
}

impl Curl {
    pub fn with_input(command: &str) -> Result<Self, Error> {
        let mut record: Vec<&str> = Vec::new();
        for datum in command.split_terminator("\\\n").map(|x| x.trim()) {
            match datum.split_once(" ") {
                Some((key, val)) => {
                    record.push(key);
                    record.push(val.trim_matches('\''));
                }
                None => continue
            }
        };
        return Curl::try_parse_from(record);
    }
}

pub fn transform_curl(curl: Curl) -> Option<domain::public_service::CurlInfo> {
    let mut headers = vec![];
    for header in curl.header {
        let (key, value) = header.split_once(": ").unwrap_or_default();
        headers.push(domain::public::HttpHeader { key: key.to_string(), value: value.to_string() });
    }

    let info = domain::public_service::CurlInfo {
        url: match url::Url::parse(&curl.url) {
            Ok(v) => {
                let protocol = match domain::public::HttpProtocol::from_str_name(v.scheme().to_uppercase().as_str()) {
                    Some(p) => p,
                    None => domain::public::HttpProtocol::Http,
                };
                Some(domain::public_service::Url {
                    protocol: protocol as i32,
                    host: v.host_str().unwrap_or_default().to_string(),
                    port: v.port().unwrap_or_default() as u32,
                    name: v.username().to_string(),
                    pass: v.password().unwrap_or_default().to_string(),
                    uri: v.path().to_string(),
                    query_string: v.query().unwrap_or_default().to_string(),
                })
            }
            Err(_) => None,
        },
        method: match curl.method {
            Some(v) => match domain::public::HttpMethod::from_str_name(v.to_uppercase().as_str()) {
                Some(method) => method as i32,
                None => domain::public::HttpMethod::Get as i32,
            },
            None => {
                if curl.data == None {
                    domain::public::HttpMethod::Get as i32
                } else {
                    domain::public::HttpMethod::Post as i32
                }
            }
        },
        headers,
        body: match curl.data {
            Some(v) => v,
            None => String::new(),
        },
    };

    return Some(info);
}