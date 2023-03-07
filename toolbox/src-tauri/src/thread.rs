use std::net::{IpAddr, TcpStream};
use std::process;
use std::str::FromStr;
use std::sync::mpsc::{channel, Sender};
use std::thread;
use std::time::Instant;

use clap::Parser;

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

fn scan(tx: Sender<u16>, start_port: u16, addr: IpAddr, num_threads: u16) {
    let mut port: u16 = start_port + 1;
    loop {
        match TcpStream::connect((addr, port)) {
            Ok(_) => tx.send(port).unwrap(),
            Err(_) => {}
        }
        if MAX - port <= num_threads {
            break;
        }
        port += num_threads;
    }
}

fn main() {
    let args = Arguments::parse();
    let addr = match IpAddr::from_str(&args.target) {
        Ok(ip) => ip,
        Err(_) => {
            eprintln!("target not a valid IP address, must be IPv4 or IPv6");
            process::exit(0);
        }
    };
    println!("running with {} threads, please wait...", args.num);
    let start = Instant::now();
    let num_threads = args.num;
    let (sender, receiver) = channel();
    for i in 0..num_threads {
        let tx = sender.clone();
        thread::spawn(move || scan(tx, i, addr, num_threads));
    }

    let mut out = vec![];
    drop(sender);
    for p in receiver {
        out.push(p);
    }
    println!("success, open {} ports, cost is {}ms", out.len(), start.elapsed().as_millis());
    out.sort();
    for v in out {
        println!("{} is open", v);
    }
}
