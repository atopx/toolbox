#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Book {
    #[prost(int32, tag = "1")]
    pub id: i32,
    #[prost(string, tag = "2")]
    pub src: ::prost::alloc::string::String,
    #[prost(string, tag = "3")]
    pub name: ::prost::alloc::string::String,
    #[prost(string, tag = "4")]
    pub author: ::prost::alloc::string::String,
    #[prost(string, tag = "5")]
    pub label: ::prost::alloc::string::String,
    #[prost(string, tag = "6")]
    pub cover: ::prost::alloc::string::String,
    #[prost(string, tag = "7")]
    pub intro: ::prost::alloc::string::String,
    #[prost(enumeration = "super::public::TaskStatus", tag = "8")]
    pub state: i32,
    #[prost(int64, tag = "9")]
    pub last_modify: i64,
    /// 删除时间 时间戳：秒
    #[prost(int64, tag = "1001")]
    pub delete_time: i64,
    /// 创建时间 时间戳：秒
    #[prost(int64, tag = "1002")]
    pub create_time: i64,
    /// 最后更新时间 时间戳：秒
    #[prost(int64, tag = "1003")]
    pub update_time: i64,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct BookFilter {
    #[prost(int32, repeated, tag = "1")]
    pub ids: ::prost::alloc::vec::Vec<i32>,
    #[prost(enumeration = "super::public::TaskStatus", repeated, tag = "2")]
    pub states: ::prost::alloc::vec::Vec<i32>,
    #[prost(string, repeated, tag = "3")]
    pub srcs: ::prost::alloc::vec::Vec<::prost::alloc::string::String>,
    #[prost(string, repeated, tag = "4")]
    pub names: ::prost::alloc::vec::Vec<::prost::alloc::string::String>,
    #[prost(string, repeated, tag = "5")]
    pub authors: ::prost::alloc::vec::Vec<::prost::alloc::string::String>,
    #[prost(string, repeated, tag = "6")]
    pub labels: ::prost::alloc::vec::Vec<::prost::alloc::string::String>,
    #[prost(message, optional, tag = "20")]
    pub last_modify: ::core::option::Option<super::public::BetweenInt64>,
    #[prost(message, optional, tag = "101")]
    pub delete_time_range: ::core::option::Option<super::public::BetweenInt64>,
    #[prost(message, optional, tag = "102")]
    pub create_time_range: ::core::option::Option<super::public::BetweenInt64>,
    #[prost(message, optional, tag = "103")]
    pub update_time_range: ::core::option::Option<super::public::BetweenInt64>,
    #[prost(enumeration = "super::public::BooleanScope", tag = "500")]
    pub deleted: i32,
    #[prost(message, optional, tag = "201")]
    pub keywords: ::core::option::Option<book_filter::Keywords>,
}
/// Nested message and enum types in `BookFilter`.
pub mod book_filter {
    #[allow(clippy::derive_partial_eq_without_eq)]
    #[derive(Clone, PartialEq, ::prost::Message)]
    pub struct Keywords {
        #[prost(string, tag = "1")]
        pub keyword: ::prost::alloc::string::String,
        #[prost(string, tag = "2")]
        pub name: ::prost::alloc::string::String,
        #[prost(string, tag = "3")]
        pub author: ::prost::alloc::string::String,
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ListBookParams {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::Header>,
    #[prost(message, optional, tag = "2")]
    pub pager: ::core::option::Option<super::public::Pager>,
    #[prost(message, repeated, tag = "3")]
    pub sorts: ::prost::alloc::vec::Vec<super::public::Sort>,
    #[prost(message, optional, tag = "4")]
    pub filter: ::core::option::Option<BookFilter>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ListBookReply {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::ReplyHeader>,
    #[prost(message, optional, tag = "2")]
    pub pager: ::core::option::Option<super::public::Pager>,
    #[prost(message, repeated, tag = "3")]
    pub data: ::prost::alloc::vec::Vec<Book>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Chapter {
    #[prost(int32, tag = "1")]
    pub id: i32,
    #[prost(int32, tag = "2")]
    pub book_id: i32,
    #[prost(int32, tag = "3")]
    pub code: i32,
    #[prost(string, tag = "4")]
    pub src: ::prost::alloc::string::String,
    #[prost(string, tag = "5")]
    pub title: ::prost::alloc::string::String,
    #[prost(string, tag = "6")]
    pub content: ::prost::alloc::string::String,
    #[prost(enumeration = "super::public::TaskStatus", tag = "7")]
    pub state: i32,
    /// 删除时间 时间戳：秒
    #[prost(int64, tag = "1001")]
    pub delete_time: i64,
    /// 创建时间 时间戳：秒
    #[prost(int64, tag = "1002")]
    pub create_time: i64,
    /// 最后更新时间 时间戳：秒
    #[prost(int64, tag = "1003")]
    pub update_time: i64,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ChapterFilter {
    #[prost(int32, repeated, tag = "1")]
    pub ids: ::prost::alloc::vec::Vec<i32>,
    #[prost(int32, repeated, tag = "2")]
    pub book_ids: ::prost::alloc::vec::Vec<i32>,
    #[prost(string, repeated, tag = "3")]
    pub srcs: ::prost::alloc::vec::Vec<::prost::alloc::string::String>,
    #[prost(enumeration = "super::public::TaskStatus", repeated, tag = "4")]
    pub states: ::prost::alloc::vec::Vec<i32>,
    #[prost(message, optional, tag = "101")]
    pub delete_time_range: ::core::option::Option<super::public::BetweenInt64>,
    #[prost(message, optional, tag = "102")]
    pub create_time_range: ::core::option::Option<super::public::BetweenInt64>,
    #[prost(message, optional, tag = "103")]
    pub update_time_range: ::core::option::Option<super::public::BetweenInt64>,
    #[prost(enumeration = "super::public::BooleanScope", tag = "500")]
    pub deleted: i32,
    #[prost(message, optional, tag = "201")]
    pub keywords: ::core::option::Option<chapter_filter::Keywords>,
}
/// Nested message and enum types in `ChapterFilter`.
pub mod chapter_filter {
    #[allow(clippy::derive_partial_eq_without_eq)]
    #[derive(Clone, PartialEq, ::prost::Message)]
    pub struct Keywords {
        #[prost(string, tag = "1")]
        pub keyword: ::prost::alloc::string::String,
    }
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ListChapterParams {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::Header>,
    #[prost(message, optional, tag = "2")]
    pub pager: ::core::option::Option<super::public::Pager>,
    #[prost(message, repeated, tag = "3")]
    pub sorts: ::prost::alloc::vec::Vec<super::public::Sort>,
    #[prost(message, optional, tag = "4")]
    pub filter: ::core::option::Option<ChapterFilter>,
}
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct ListChapterReply {
    #[prost(message, optional, tag = "1")]
    pub header: ::core::option::Option<super::public::ReplyHeader>,
    #[prost(message, optional, tag = "2")]
    pub pager: ::core::option::Option<super::public::Pager>,
    #[prost(message, repeated, tag = "3")]
    pub data: ::prost::alloc::vec::Vec<Chapter>,
}
/// Generated server implementations.
pub mod novel_service_server {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    /// Generated trait containing gRPC methods that should be implemented for use with NovelServiceServer.
    #[async_trait]
    pub trait NovelService: Send + Sync + 'static {
        async fn list_book(
            &self,
            request: tonic::Request<super::ListBookParams>,
        ) -> std::result::Result<tonic::Response<super::ListBookReply>, tonic::Status>;
        async fn list_chapter(
            &self,
            request: tonic::Request<super::ListChapterParams>,
        ) -> std::result::Result<
            tonic::Response<super::ListChapterReply>,
            tonic::Status,
        >;
    }
    #[derive(Debug)]
    pub struct NovelServiceServer<T: NovelService> {
        inner: _Inner<T>,
        accept_compression_encodings: EnabledCompressionEncodings,
        send_compression_encodings: EnabledCompressionEncodings,
        max_decoding_message_size: Option<usize>,
        max_encoding_message_size: Option<usize>,
    }
    struct _Inner<T>(Arc<T>);
    impl<T: NovelService> NovelServiceServer<T> {
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
    impl<T, B> tonic::codegen::Service<http::Request<B>> for NovelServiceServer<T>
    where
        T: NovelService,
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
                "/novel_service.NovelService/ListBook" => {
                    #[allow(non_camel_case_types)]
                    struct ListBookSvc<T: NovelService>(pub Arc<T>);
                    impl<
                        T: NovelService,
                    > tonic::server::UnaryService<super::ListBookParams>
                    for ListBookSvc<T> {
                        type Response = super::ListBookReply;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::ListBookParams>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move { (*inner).list_book(request).await };
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
                        let method = ListBookSvc(inner);
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
                "/novel_service.NovelService/ListChapter" => {
                    #[allow(non_camel_case_types)]
                    struct ListChapterSvc<T: NovelService>(pub Arc<T>);
                    impl<
                        T: NovelService,
                    > tonic::server::UnaryService<super::ListChapterParams>
                    for ListChapterSvc<T> {
                        type Response = super::ListChapterReply;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::ListChapterParams>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                (*inner).list_chapter(request).await
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
                        let method = ListChapterSvc(inner);
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
    impl<T: NovelService> Clone for NovelServiceServer<T> {
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
    impl<T: NovelService> Clone for _Inner<T> {
        fn clone(&self) -> Self {
            Self(Arc::clone(&self.0))
        }
    }
    impl<T: std::fmt::Debug> std::fmt::Debug for _Inner<T> {
        fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
            write!(f, "{:?}", self.0)
        }
    }
    impl<T: NovelService> tonic::server::NamedService for NovelServiceServer<T> {
        const NAME: &'static str = "novel_service.NovelService";
    }
}
