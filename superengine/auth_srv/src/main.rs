use domain::auth_service::auth_service_server::AuthServiceServer;
use sqlx::mysql::MySqlPoolOptions;

use std::time::Duration;

pub mod server;

#[tokio::main]
async fn main() {
    std::env::set_var("AUTH_ADDR", "127.0.0.1:18001");
    std::env::set_var(
        "DB_DSN",
        "mysql://root:123456@127.0.0.1:3306/superserver",
    );

    let addr = std::env::var("AUTH_ADDR").unwrap();
    let db_dsn = std::env::var("DB_DSN").unwrap();

    let dbpool = MySqlPoolOptions::new()
        .min_connections(10)
        .max_connections(100)
        .idle_timeout(Duration::from_secs(60))
        .max_lifetime(Duration::from_secs(3600))
        .connect(&db_dsn)
        .await
        .unwrap();

    let service = server::AuthService::new(dbpool);

    tonic::transport::Server::builder()
        .add_service(AuthServiceServer::new(service))
        .serve(addr.parse().unwrap())
        .await
        .unwrap()
}
