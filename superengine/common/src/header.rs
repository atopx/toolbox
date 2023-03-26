use domain::public::{ECode, ReplyHeader};

pub fn reply(code: ECode) -> Option<ReplyHeader> {
    return Some(ReplyHeader {
        code: code.into(),
        message: code.as_str_name().parse().unwrap(),
    });
}
