use std::fmt::{Display, Formatter};

pub enum TransformError {
    ParseError(String),
    TransError(String),
}

impl Display for TransformError {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        return match self {
            Self::ParseError(msg) => f.write_str(msg.as_str()),
            Self::TransError(msg) => f.write_str(msg.as_str()),
        };
    }
}


pub fn parse_error(msg: &str) -> TransformError {
    return TransformError::ParseError(format!("transform parse error: {msg}"));
}

pub fn trans_error(msg: &str) -> TransformError {
    return TransformError::TransError(format!("transform trans error: {msg}"));
}