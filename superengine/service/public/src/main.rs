mod server;
pub mod business;

use domain::public_service::public_service_server::PublicServiceServer;

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
    let service = server::PublicService::new(db);
    tonic::transport::Server::builder()
        .add_service(PublicServiceServer::new(service))
        .serve(config.auth_srv.parse().unwrap())
        .await
        .unwrap();
}
