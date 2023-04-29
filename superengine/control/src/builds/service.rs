use std::process::Command;


pub fn build() {
    let cmd = "cargo build --release --package service";
    let output = Command::new("bash")
        .arg("-c")
        .arg(cmd)
        .output().unwrap();
    if !output.status.success() {
        println!("{}", String::from_utf8(output.stderr).unwrap());
    }
}