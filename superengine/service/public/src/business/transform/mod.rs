mod interface;
mod shape;
mod error;


/// # `ITransform`转换器抽象: `source` -> `vec<shape>` -> `result`
pub trait ITransform {
    /// ### `parser` 从原始数据转换为中继数据
    fn parser(&self, value: &str) -> Result<Vec<shape::Shape>, error::TransformError>;

    /// ### `transfer` 从中继数据转换为目标数据
    fn transfer(&self, shapes: Vec<shape::Shape>) -> Result<String, error::TransformError>;
}

/// ### `transform` 构建转换器对象
fn transform(t: domain::public_service::TransType) -> Box<dyn ITransform> {
    return match t {
        domain::public_service::TransType::Json => Box::new(interface::json::Transform {}),
        domain::public_service::TransType::Yaml => Box::new(interface::yaml::Transform {}),
        domain::public_service::TransType::Toml => Box::new(interface::toml::Transform {}),
        domain::public_service::TransType::Sql => Box::new(interface::sql::Transform {}),
        domain::public_service::TransType::Xml => Box::new(interface::xml::Transform {}),
        domain::public_service::TransType::Protobuf => Box::new(interface::protobuf::Transform {}),
        domain::public_service::TransType::Golang => Box::new(interface::golang::Transform {}),
        domain::public_service::TransType::Rust => Box::new(interface::rust::Transform {}),
        domain::public_service::TransType::Java => Box::new(interface::java::Transform {}),
        domain::public_service::TransType::Python => Box::new(interface::python::Transform {}),
        domain::public_service::TransType::Typescript => Box::new(interface::typescript::Transform {}),
    };
}


/// # `goto` 执行数据转换
pub fn goto(data: domain::public_service::Transform) -> Result<String, error::TransformError> {
    let shapes = transform(data.from()).parser(data.value.as_str())?;
    return transform(data.to()).transfer(shapes);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_transfer() {
        let result = goto(domain::public_service::Transform {
            value: String::new(),
            from: domain::public_service::TransType::Python.into(),
            to: domain::public_service::TransType::Json.into(),
        });
        match result {
            Ok(v) => println!("{}", v),
            Err(e) => println!("{}", e.to_string()),
        }
    }
}
