use domain::ecode::ECode;
use domain::public::ReplyHeader;

pub fn reply(id: i64) -> Option<ReplyHeader> {
    Some(ReplyHeader {
        trace_id: id,
        code: ECode::Success.into(),
        message: Default::default(),
    })
}

pub fn err_reply(id: i64, code: ECode, msg: String) -> Option<ReplyHeader> {
    Some(ReplyHeader {
        trace_id: id,
        code: code.into(),
        message: msg,
    })
}

