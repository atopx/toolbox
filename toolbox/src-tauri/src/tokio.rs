use std::net::IpAddr;
use std::str::FromStr;

use std::process;
use tokio::net::TcpStream;
use clap::Parser;
use tokio::runtime::Runtime;
use tokio::time::Instant;


/// this is a multi-thread ip port sniffer.
#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
struct Arguments {
    /// [num] threads
    #[arg(short, long, default_value_t = 1)]
    num: u16,

    /// target [ipaddr]
    #[arg(short, long)]
    target: String,
}

const MAX: u16 = 65535;



fn main() {
    let args = Arguments::parse();
    let addr = match IpAddr::from_str(&args.target) {
        Ok(ip) => ip,
        Err(_) => {
            eprintln!("target not a valid IP address, must be IPv4 or IPv6");
            process::exit(0);
        }
    };

    let rt = Runtime::new().unwrap();
    
    let start = Instant::now();
    let num_threads = args.num;
    for i in 0..num_threads {
        rt.block_on(async move {
            let mut port = i + 1;
            loop {
                match TcpStream::connect((addr, port)).await {
                    Ok(_) => { println!("open port {}", port) },
                    Err(_) => {}
                }
                if MAX - port <= num_threads {
                    break;
                }
                port += num_threads
            }
        });
    }
    println!("success, cost: {}ms", start.elapsed().as_millis());
}
