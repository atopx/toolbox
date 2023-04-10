use std::collections::HashMap;
use std::fs;
use std::io::{BufRead, Write};
use std::ops::Not;

struct Meta {
    code: String,
    value: String,
    desc: String,
}

const TAB: &str = "    ";
const HEADER: &str = r#"/// 根据proto枚举自动生成, 构建代码查看 build.rs
use axum::{Json, http::StatusCode};
use axum::extract::Path;
use serde::Serialize;

#[derive(Serialize)]
pub struct EnumMate {
    code: &'static str,
    desc: &'static str,
    value: i32,
}
"#;


fn main() {
    // 一般不编译
    println!("cargo:rerun-if-changed=build.rs");

    let enum_proto = vec![
        ("../../protocol/proto/public/common.proto", "Source"),
        ("../../protocol/proto/public/common.proto", "RangeScope"),
        ("../../protocol/proto/public/common.proto", "BooleanScope"),
        ("../../protocol/proto/public/ecode.proto", "ECode"),
    ];

    let mut records: HashMap<String, Vec<Meta>> = HashMap::new();

    for (path, schema) in enum_proto {
        let f = fs::File::open(path).expect("proto file not found");

        let buffered = std::io::BufReader::new(f);
        let mut record = vec![];
        let mut state = false;
        for line in buffered.lines() {
            let line = line.unwrap_or_default();
            let line = line.trim().to_owned();

            if line.is_empty() {
                continue;
            }

            if state {
                if line.starts_with('}') {
                    break;
                }
                if line.starts_with("//") {
                    continue;
                }
                let (mut code, other) = line.as_str().split_once('=').unwrap();
                code = code.trim();
                let (mut value, other) = other.split_once(';').unwrap();
                value = value.trim();
                let (_, mut comment) = other.split_once("//").unwrap();
                comment = comment.trim();
                record.push(Meta {
                    code: code.to_string(),
                    value: value.to_string(),
                    desc: comment.to_string(),
                });
            } else {
                // 匹配schema
                if line.starts_with("enum") && line.ends_with('{') && line.contains(schema) {
                    state = true;
                }
            }
        }
        if record.is_empty().not() {
            records.insert(schema.to_lowercase(), record);
        }
    }

    let outfile = "src/meta/mod.rs";
    let mut out = fs::File::create(outfile).
        unwrap_or(fs::File::open(outfile).unwrap());


    let mut handler: String = String::from("pub(crate) async fn meta_handler(Path(name): Path<String>) -> (StatusCode, Json<Vec<EnumMate>>) {");
    handler = format!("{handler}\n{TAB}match name.as_str() {{\n");

    writeln!(out, "{HEADER}").unwrap();
    for (schema, record) in records {
        writeln!(out, "pub async fn {schema}_data() -> (StatusCode, Json<Vec<EnumMate>>) {{").unwrap();
        writeln!(out, "{TAB}let data = vec![").unwrap();
        for line in record {
            writeln!(out, "{TAB}{TAB}EnumMate {{ code: \"{}\", value: {}, desc: \"{}\" }},",
                     line.code, line.value, line.desc).unwrap();
        }
        writeln!(out, "{TAB}];").unwrap();
        writeln!(out, "{TAB}(StatusCode::OK, Json(data))\n}}\n").unwrap();
        handler = format!("{handler}{TAB}{TAB}\"{schema}\" => {schema}_data().await,\n");
    }
    writeln!(out, "{handler}{TAB}{TAB}_=> {{ (StatusCode::OK, Json(vec![])) }}\n{TAB}}}\n}}\n").unwrap();
}
