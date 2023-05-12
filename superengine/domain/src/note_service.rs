#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct NoteLabel {
    /// 主键
    #[prost(int32, tag = "1")]
    pub id: i32,
    /// 笔记ID
    #[prost(int32, tag = "2")]
    pub note_id: i32,
    /// 标签ID
    #[prost(int32, tag = "3")]
    pub label_id: i32,
    /// 标签类型
    #[prost(enumeration = "LabelType", tag = "4")]
    pub label_type: i32,
    /// 删除时间 时间戳：秒
    #[prost(int64, tag = "1001")]
    pub delete_time: i64,
    /// 创建时间 时间戳：秒
    #[prost(int64, tag = "1002")]
    pub create_time: i64,
    /// 最后更新时间 时间戳：秒
    #[prost(int64, tag = "1003")]
    pub update_time: i64,
    /// 创建人
    #[prost(int32, tag = "1004")]
    pub creator: i32,
    /// 更新人
    #[prost(int32, tag = "1005")]
    pub updater: i32,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct NoteLabelFilter {
    #[prost(int32, repeated, tag = "1")]
    pub ids: ::prost::alloc::vec::Vec<i32>,
    #[prost(int32, repeated, tag = "2")]
    pub note_ids: ::prost::alloc::vec::Vec<i32>,
    #[prost(int32, repeated, tag = "3")]
    pub label_ids: ::prost::alloc::vec::Vec<i32>,
    #[prost(enumeration = "LabelType", repeated, tag = "4")]
    pub label_types: ::prost::alloc::vec::Vec<i32>,
    #[prost(message, optional, tag = "101")]
    pub delete_time_range: ::core::option::Option<super::public::BetweenInt64>,
    #[prost(message, optional, tag = "102")]
    pub create_time_range: ::core::option::Option<super::public::BetweenInt64>,
    #[prost(message, optional, tag = "103")]
    pub update_time_range: ::core::option::Option<super::public::BetweenInt64>,
    #[prost(int32, repeated, tag = "104")]
    pub creators: ::prost::alloc::vec::Vec<i32>,
    #[prost(int32, repeated, tag = "105")]
    pub updaters: ::prost::alloc::vec::Vec<i32>,
    #[prost(enumeration = "super::public::BooleanScope", tag = "500")]
    pub deleted: i32,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ListNoteLabelParams {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::Header>,
    #[prost(message, optional, tag = "2")]
    pub pager: ::core::option::Option<super::public::Pager>,
    #[prost(message, repeated, tag = "3")]
    pub sorts: ::prost::alloc::vec::Vec<super::public::Sort>,
    #[prost(message, optional, tag = "4")]
    pub filter: ::core::option::Option<NoteLabelFilter>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ListNoteLabelReply {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::ReplyHeader>,
    #[prost(message, optional, tag = "2")]
    pub pager: ::core::option::Option<super::public::Pager>,
    #[prost(message, repeated, tag = "3")]
    pub data: ::prost::alloc::vec::Vec<NoteLabel>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct OperateNoteLabelParams {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::Header>,
    #[prost(enumeration = "super::public::Operation", tag = "2")]
    pub operate: i32,
    #[prost(message, optional, tag = "3")]
    pub data: ::core::option::Option<NoteLabel>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct OperateNoteLabelReply {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::ReplyHeader>,
    #[prost(message, optional, tag = "2")]
    pub data: ::core::option::Option<NoteLabel>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct BatchOperateNoteLabelParams {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::Header>,
    #[prost(enumeration = "super::public::Operation", tag = "2")]
    pub operate: i32,
    #[prost(message, repeated, tag = "3")]
    pub data: ::prost::alloc::vec::Vec<NoteLabel>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct BatchOperateNoteLabelReply {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::ReplyHeader>,
}
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum LabelType {
    Tag = 0,
    Topic = 1,
}
impl LabelType {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            LabelType::Tag => "TAG",
            LabelType::Topic => "TOPIC",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "TAG" => Some(Self::Tag),
            "TOPIC" => Some(Self::Topic),
            _ => None,
        }
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Note {
    /// 主键
    #[prost(int32, tag = "1")]
    pub id: i32,
    /// 标题
    #[prost(string, tag = "2")]
    pub title: ::prost::alloc::string::String,
    /// 笔记内容
    #[prost(string, tag = "3")]
    pub content: ::prost::alloc::string::String,
    /// 删除时间 时间戳：秒
    #[prost(int64, tag = "1001")]
    pub delete_time: i64,
    /// 创建时间 时间戳：秒
    #[prost(int64, tag = "1002")]
    pub create_time: i64,
    /// 最后更新时间 时间戳：秒
    #[prost(int64, tag = "1003")]
    pub update_time: i64,
    /// 创建人
    #[prost(int32, tag = "1004")]
    pub creator: i32,
    /// 更新人
    #[prost(int32, tag = "1005")]
    pub updater: i32,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct NoteFilter {
    #[prost(int32, repeated, tag = "1")]
    pub ids: ::prost::alloc::vec::Vec<i32>,
    #[prost(message, optional, tag = "101")]
    pub delete_time_range: ::core::option::Option<super::public::BetweenInt64>,
    #[prost(message, optional, tag = "102")]
    pub create_time_range: ::core::option::Option<super::public::BetweenInt64>,
    #[prost(message, optional, tag = "103")]
    pub update_time_range: ::core::option::Option<super::public::BetweenInt64>,
    #[prost(int32, repeated, tag = "104")]
    pub creators: ::prost::alloc::vec::Vec<i32>,
    #[prost(int32, repeated, tag = "105")]
    pub updaters: ::prost::alloc::vec::Vec<i32>,
    #[prost(enumeration = "super::public::BooleanScope", tag = "500")]
    pub deleted: i32,
    #[prost(message, optional, tag = "201")]
    pub keywords: ::core::option::Option<note_filter::Keywords>,
}
/// Nested message and enum types in `NoteFilter`.
pub mod note_filter {
    #[allow(clippy::derive_partial_eq_without_eq)]
    #[derive(Clone, PartialEq, ::prost::Message)]
    pub struct Keywords {
        #[prost(string, tag = "1")]
        pub keyword: ::prost::alloc::string::String,
        #[prost(string, tag = "2")]
        pub title: ::prost::alloc::string::String,
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ListNoteParams {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::Header>,
    #[prost(message, optional, tag = "2")]
    pub pager: ::core::option::Option<super::public::Pager>,
    #[prost(message, repeated, tag = "3")]
    pub sorts: ::prost::alloc::vec::Vec<super::public::Sort>,
    #[prost(message, optional, tag = "4")]
    pub filter: ::core::option::Option<NoteFilter>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ListNoteReply {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::ReplyHeader>,
    #[prost(message, optional, tag = "2")]
    pub pager: ::core::option::Option<super::public::Pager>,
    #[prost(message, repeated, tag = "3")]
    pub data: ::prost::alloc::vec::Vec<Note>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct OperateNoteParams {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::Header>,
    #[prost(enumeration = "super::public::Operation", tag = "2")]
    pub operate: i32,
    #[prost(message, optional, tag = "3")]
    pub data: ::core::option::Option<Note>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct OperateNoteReply {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::ReplyHeader>,
    #[prost(message, optional, tag = "2")]
    pub data: ::core::option::Option<Note>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct BatchOperateNoteParams {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::Header>,
    #[prost(enumeration = "super::public::Operation", tag = "2")]
    pub operate: i32,
    #[prost(message, repeated, tag = "3")]
    pub data: ::prost::alloc::vec::Vec<Note>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct BatchOperateNoteReply {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::ReplyHeader>,
}
/// Generated server implementations.
pub mod note_service_server {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    /// Generated trait containing gRPC methods that should be implemented for use with NoteServiceServer.
    #[async_trait]
    pub trait NoteService: Send + Sync + 'static {
        /// 笔记
        async fn list_note(
            &self,
            request: tonic::Request<super::ListNoteParams>,
        ) -> std::result::Result<tonic::Response<super::ListNoteReply>, tonic::Status>;
        async fn operate_note(
            &self,
            request: tonic::Request<super::OperateNoteParams>,
        ) -> std::result::Result<
            tonic::Response<super::OperateNoteReply>,
            tonic::Status,
        >;
        async fn batch_operate_note(
            &self,
            request: tonic::Request<super::BatchOperateNoteParams>,
        ) -> std::result::Result<
            tonic::Response<super::BatchOperateNoteReply>,
            tonic::Status,
        >;
        /// 笔记标签
        async fn list_note_label(
            &self,
            request: tonic::Request<super::ListNoteLabelParams>,
        ) -> std::result::Result<
            tonic::Response<super::ListNoteLabelReply>,
            tonic::Status,
        >;
        async fn operate_note_label(
            &self,
            request: tonic::Request<super::OperateNoteLabelParams>,
        ) -> std::result::Result<
            tonic::Response<super::OperateNoteLabelReply>,
            tonic::Status,
        >;
        async fn batch_operate_note_label(
            &self,
            request: tonic::Request<super::BatchOperateNoteLabelParams>,
        ) -> std::result::Result<
            tonic::Response<super::BatchOperateNoteLabelReply>,
            tonic::Status,
        >;
    }
    #[derive(Debug)]
    pub struct NoteServiceServer<T: NoteService> {
        inner: _Inner<T>,
        accept_compression_encodings: EnabledCompressionEncodings,
        send_compression_encodings: EnabledCompressionEncodings,
        max_decoding_message_size: Option<usize>,
        max_encoding_message_size: Option<usize>,
    }
    struct _Inner<T>(Arc<T>);
    impl<T: NoteService> NoteServiceServer<T> {
        pub fn new(inner: T) -> Self {
            Self::from_arc(Arc::new(inner))
        }
        pub fn from_arc(inner: Arc<T>) -> Self {
            let inner = _Inner(inner);
            Self {
                inner,
                accept_compression_encodings: Default::default(),
                send_compression_encodings: Default::default(),
                max_decoding_message_size: None,
                max_encoding_message_size: None,
            }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> InterceptedService<Self, F>
        where
            F: tonic::service::Interceptor,
        {
            InterceptedService::new(Self::new(inner), interceptor)
        }
        /// Enable decompressing requests with the given encoding.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.accept_compression_encodings.enable(encoding);
            self
        }
        /// Compress responses with the given encoding, if the client supports it.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.send_compression_encodings.enable(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.max_decoding_message_size = Some(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.max_encoding_message_size = Some(limit);
            self
        }
    }
    impl<T, B> tonic::codegen::Service<http::Request<B>> for NoteServiceServer<T>
    where
        T: NoteService,
        B: Body + Send + 'static,
        B::Error: Into<StdError> + Send + 'static,
    {
        type Response = http::Response<tonic::body::BoxBody>;
        type Error = std::convert::Infallible;
        type Future = BoxFuture<Self::Response, Self::Error>;
        fn poll_ready(
            &mut self,
            _cx: &mut Context<'_>,
        ) -> Poll<std::result::Result<(), Self::Error>> {
            Poll::Ready(Ok(()))
        }
        fn call(&mut self, req: http::Request<B>) -> Self::Future {
            let inner = self.inner.clone();
            match req.uri().path() {
                "/note_service.NoteService/ListNote" => {
                    #[allow(non_camel_case_types)]
                    struct ListNoteSvc<T: NoteService>(pub Arc<T>);
                    impl<
                        T: NoteService,
                    > tonic::server::UnaryService<super::ListNoteParams>
                    for ListNoteSvc<T> {
                        type Response = super::ListNoteReply;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::ListNoteParams>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move { (*inner).list_note(request).await };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = ListNoteSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/note_service.NoteService/OperateNote" => {
                    #[allow(non_camel_case_types)]
                    struct OperateNoteSvc<T: NoteService>(pub Arc<T>);
                    impl<
                        T: NoteService,
                    > tonic::server::UnaryService<super::OperateNoteParams>
                    for OperateNoteSvc<T> {
                        type Response = super::OperateNoteReply;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::OperateNoteParams>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                (*inner).operate_note(request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = OperateNoteSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/note_service.NoteService/BatchOperateNote" => {
                    #[allow(non_camel_case_types)]
                    struct BatchOperateNoteSvc<T: NoteService>(pub Arc<T>);
                    impl<
                        T: NoteService,
                    > tonic::server::UnaryService<super::BatchOperateNoteParams>
                    for BatchOperateNoteSvc<T> {
                        type Response = super::BatchOperateNoteReply;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::BatchOperateNoteParams>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                (*inner).batch_operate_note(request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = BatchOperateNoteSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/note_service.NoteService/ListNoteLabel" => {
                    #[allow(non_camel_case_types)]
                    struct ListNoteLabelSvc<T: NoteService>(pub Arc<T>);
                    impl<
                        T: NoteService,
                    > tonic::server::UnaryService<super::ListNoteLabelParams>
                    for ListNoteLabelSvc<T> {
                        type Response = super::ListNoteLabelReply;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::ListNoteLabelParams>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                (*inner).list_note_label(request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = ListNoteLabelSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/note_service.NoteService/OperateNoteLabel" => {
                    #[allow(non_camel_case_types)]
                    struct OperateNoteLabelSvc<T: NoteService>(pub Arc<T>);
                    impl<
                        T: NoteService,
                    > tonic::server::UnaryService<super::OperateNoteLabelParams>
                    for OperateNoteLabelSvc<T> {
                        type Response = super::OperateNoteLabelReply;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::OperateNoteLabelParams>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                (*inner).operate_note_label(request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = OperateNoteLabelSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/note_service.NoteService/BatchOperateNoteLabel" => {
                    #[allow(non_camel_case_types)]
                    struct BatchOperateNoteLabelSvc<T: NoteService>(pub Arc<T>);
                    impl<
                        T: NoteService,
                    > tonic::server::UnaryService<super::BatchOperateNoteLabelParams>
                    for BatchOperateNoteLabelSvc<T> {
                        type Response = super::BatchOperateNoteLabelReply;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::BatchOperateNoteLabelParams>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                (*inner).batch_operate_note_label(request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let inner = inner.0;
                        let method = BatchOperateNoteLabelSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                _ => {
                    Box::pin(async move {
                        Ok(
                            http::Response::builder()
                                .status(200)
                                .header("grpc-status", "12")
                                .header("content-type", "application/grpc")
                                .body(empty_body())
                                .unwrap(),
                        )
                    })
                }
            }
        }
    }
    impl<T: NoteService> Clone for NoteServiceServer<T> {
        fn clone(&self) -> Self {
            let inner = self.inner.clone();
            Self {
                inner,
                accept_compression_encodings: self.accept_compression_encodings,
                send_compression_encodings: self.send_compression_encodings,
                max_decoding_message_size: self.max_decoding_message_size,
                max_encoding_message_size: self.max_encoding_message_size,
            }
        }
    }
    impl<T: NoteService> Clone for _Inner<T> {
        fn clone(&self) -> Self {
            Self(Arc::clone(&self.0))
        }
    }
    impl<T: std::fmt::Debug> std::fmt::Debug for _Inner<T> {
        fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
            write!(f, "{:?}", self.0)
        }
    }
    impl<T: NoteService> tonic::server::NamedService for NoteServiceServer<T> {
        const NAME: &'static str = "note_service.NoteService";
    }
}
