use clap::{Parser, ValueEnum};

mod proto;
mod enums;
mod swagger;
mod service;
mod server;

#[derive(Parser, Debug, ValueEnum, Clone)]
pub enum Code {
    All,
    Proto,
    Enums,
    Swagger,
    Server,
    Service,
}

pub fn run(code: Code) {
    match code {
        Code::All => {
            proto::build();
            enums::build();
            swagger::build();
            server::build();
            service::build();
        }
        Code::Proto => proto::build(),
        Code::Enums => enums::build(),
        Code::Swagger => swagger::build(),
        Code::Server => server::build(),
        Code::Service => service::build(),
    }
}