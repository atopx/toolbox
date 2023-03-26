use domain::public::{ECode, ReplyHeader};

pub fn reply(id: i64, code: ECode) -> Option<ReplyHeader> {
    return Some(ReplyHeader {
        trace_id: id,
        code: code.into(),
        message: code.as_str_name().parse().unwrap(),
    });
}

pub fn err_reply(id: i64, code: ECode, msg: String) -> Option<ReplyHeader> {
    return Some(ReplyHeader {
        trace_id: id,
        code: code.into(),
        message: msg,
    });
}
