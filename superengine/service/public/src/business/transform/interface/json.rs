use crate::business::transform::shape;
use crate::business::transform::error::{parse_error, trans_error, TransformError};
use crate::business::transform::ITransform as Base;

pub struct Transform;

impl Base for Transform {
    // 从原始数据转换为中继数据
    fn parser(&self, value: &str) -> Result<Vec<shape::Shape>, TransformError> {
        return Err(parse_error("unimplemented"));
    }

    // 从中继数据转换为目标数据
    fn transfer(&self, shapes: Vec<shape::Shape>) -> Result<String, TransformError> {
        return Err(trans_error("unimplemented"));
    }
}