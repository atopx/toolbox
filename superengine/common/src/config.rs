pub struct Config {
    pub auth_srv: String,
    pub db_url: String,
}

impl Config {
    pub fn load() -> Self {
        dotenvy::dotenv().ok();
        return Self {
            auth_srv: std::env::var("AUTH_SERVICE").expect("DATABASE_URL is not set in .env file"),
            db_url: std::env::var("DATABASE_URL").expect("DATABASE_URL is not set in .env file"),
        };
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_load() {
        let cfg = Config::load();
        assert_eq!(cfg.auth_srv, "127.0.0.1:18001");
    }
}
