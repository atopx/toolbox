pub enum NumberShape {
    Default,
    Int32,
    Int64,
    Uint32,
    Uint64,
    Float32,
    Float64,
}

pub enum StringShape {
    Default,
    Char,
    Varchar,
    Text,
}


pub enum Shape {
    /// #### 特殊类型
    Default,
    /// 空, `nil` / `null` / `none`
    None,
    /// 范型，`interface` / `any` / `T`
    Any,

    /// #### 基础类型
    /// 布尔
    Boolean,
    /// 整数
    Number(NumberShape),
    /// 字符型
    String(StringShape),

    /// #### 复合类型
    /// 可选项, example: `number | null`
    Optional(Box<Shape>),
    /// 数组
    Array { t: Box<Shape> },
    /// 元组: 类型数组+长度
    Tuple(Vec<Shape>, usize),
    /// 结构体: 用数组来描述
    Struct(Vec<(String, Shape)>),
    /// hash map
    Map { ft: Box<Shape>, vt: Box<Shape> },
}

