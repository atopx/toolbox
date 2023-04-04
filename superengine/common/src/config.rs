pub struct Config {
    pub db_url: String,
    pub auth_srv: String,
    pub mainframe_srv: String,
    pub public_srv: String,
}

impl Config {
    pub fn load() -> Self {
        dotenvy::dotenv().ok();
        return Self {
            db_url: std::env::var("DATABASE_URL").expect("DATABASE_URL is not set in .env file"),
            auth_srv: std::env::var("AUTH_SERVICE").expect("DATABASE_URL is not set in .env file"),
            mainframe_srv: std::env::var("MAINFRAME_SERVICE").expect("DATABASE_URL is not set in .env file"),
            public_srv: std::env::var("PUBLIC_SERVICE").expect("DATABASE_URL is not set in .env file"),
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

    #[test]
    fn test_split() {
        let addr = "D8BBC1DFBEFC";
        for i in 0..6 {
            println!(
                "{}, {:?}",
                i,
                u8::from_str_radix(&addr[i * 2..i + 2], 16).unwrap()
            )
        }
    }
}
// 187 132
// 216
