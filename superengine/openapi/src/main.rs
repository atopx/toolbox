use std::net::SocketAddr;
use std::str::FromStr;

use axum::{Router, routing::get};

mod meta;

#[tokio::main]
async fn main() {
    // initialize tracing
    tracing_subscriber::fmt()
        .with_max_level(tracing::Level::INFO)
        .with_test_writer()
        .init();

    // build our application with a route
    let app = Router::new()
        .route("/ping", get(ping))
        .route("/enum/:name", get(meta::meta_handler));

    // run our app with hyper
    let config = common::config::Config::load();
    let addr = SocketAddr::from_str(config.api_addr.as_str()).unwrap();
    tracing::info!("listening on {}", addr);

    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();
}


// basic handler that responds with a static string
async fn ping() -> &'static str {
    "pong"
}
