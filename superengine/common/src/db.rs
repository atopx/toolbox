use std::time::Duration;

use sea_orm::ConnectOptions;
use sea_orm::Database;
use sea_orm::DatabaseConnection;
use sea_orm::DbErr;

pub async fn new_client(dsn: &str) -> Result<DatabaseConnection, DbErr> {
    let mut opt = ConnectOptions::new(dsn.to_string());
    opt.max_connections(20)
        .min_connections(5)
        .connect_timeout(Duration::from_secs(16))
        .acquire_timeout(Duration::from_secs(16))
        .idle_timeout(Duration::from_secs(3600))
        .max_lifetime(Duration::from_secs(3600));
    Database::connect(opt).await
}
