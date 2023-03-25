use domain::public::{ECode, ReplyHeader};

pub fn reply(code: ECode) -> Option<ReplyHeader> {
    return match code {
        ECode::Success => { Some(ReplyHeader { code: 0, message: String::new() }) }
        _ => {
            Some(ReplyHeader {
                code: i32::from(code),
                message: code.as_str_name().parse().unwrap(),
            })
        }
    };
}