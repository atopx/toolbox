use std::collections::HashMap;
use std::fs;
use std::io::{BufRead, Write};
use std::ops::Not;


struct Meta {
    code: String,
    value: String,
    desc: String,
}

pub fn build() {
    let rs_outfile = "service/src/public/enums.rs";
    let golang_outfile = "../superserver/common/enum/keys.go";
    let enum_proto = vec![
        ("../protocol/proto/public/ecode.proto", "ECode"),
        ("../protocol/proto/auth/role.proto", "RoleNature"),
        ("../protocol/proto/auth/access.proto", "AccessStatus"),
        ("../protocol/proto/auth/user.proto", "UserStatus"),
    ];

    let tab: &str = "    ";
    let comment = "/// 根据proto枚举自动生成, 构建代码查看 build.rs";
    let header = "use domain::public_service::{Enum, Enums, ListEnumReply};";
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
            records.insert(schema.parse().unwrap(), record);
        }
    }

    let mut rs_out = fs::File::create(rs_outfile).
        unwrap_or(fs::File::open(rs_outfile).unwrap());

    let mut golang_out = fs::File::create(golang_outfile).
        unwrap_or(fs::File::open(golang_outfile).unwrap());

    writeln!(rs_out, "{comment}").unwrap();
    writeln!(rs_out, "{header}\n\n").unwrap();

    writeln!(golang_out, "package enum\n\ntype Enum string\n\nconst (").unwrap();

    let mut handler = format!("pub fn list_enum(trace_id: i64) -> ListEnumReply {{\n{tab}ListEnumReply {{\n");
    handler = format!("{handler}{tab}{tab}header: common::header::reply(trace_id),\n{tab}{tab}data: vec![\n");

    for (schema, record) in records {
        let fn_name = format!("get_{}()", schema.to_lowercase());
        writeln!(rs_out, "fn {fn_name} -> Enums {{").unwrap();
        writeln!(rs_out, "{tab}let data = vec![").unwrap();
        for line in record {
            writeln!(rs_out, "{tab}{tab}Enum {{ code: \"{}\".to_string(), value: {}, desc: \"{}\".to_string() }},",
                     line.code, line.value, line.desc).unwrap();
        }
        writeln!(rs_out, "{tab}];").unwrap();
        writeln!(rs_out, "{tab}Enums {{ key: \"{schema}\".to_string(), data }}\n}}\n").unwrap();

        writeln!(golang_out, "{tab}{schema} Enum = \"{schema}\"").unwrap();

        handler = format!("{handler}{tab}{tab}{tab}{fn_name},\n");
    }
    writeln!(rs_out, "{handler}{tab}{tab}],\n{tab}}}\n}}").unwrap();
    writeln!(golang_out, ")").unwrap();
}