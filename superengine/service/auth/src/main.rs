use domain::auth_service::auth_service_server::AuthServiceServer;

pub mod business;

mod server;

#[tokio::main]
async fn main() {
    tracing_subscriber::fmt()
        .with_max_level(tracing::Level::INFO)
        .with_test_writer()
        .init();

    let config = common::config::Config::load();
    let db = common::db::new_client(config.db_url.to_owned())
        .await
        .unwrap();
    let service = server::AuthService::new(db);
    tonic::transport::Server::builder()
        .add_service(AuthServiceServer::new(service))
        .serve(config.auth_srv.parse().unwrap())
        .await
        .unwrap();
}
