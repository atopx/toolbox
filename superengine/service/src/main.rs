use domain::auth_service::auth_service_server::AuthServiceServer;
use domain::public_service::public_service_server::PublicServiceServer;

mod auth;
mod public;

#[tokio::main]
async fn main() {
    tracing_subscriber::fmt()
        .with_max_level(tracing::Level::INFO)
        .with_test_writer()
        .init();

    let config = common::config::Config::load();

    let auth_srv = auth::AuthService::new(config.db_url.as_str()).await;
    let public_srv = public::PublicService::new(config.db_url.as_str()).await;


    tonic::transport::Server::builder()
        .add_service(AuthServiceServer::new(auth_srv))
        .add_service(PublicServiceServer::new(public_srv))
        .serve(config.srv_addr.parse().unwrap())
        .await
        .unwrap();
}
