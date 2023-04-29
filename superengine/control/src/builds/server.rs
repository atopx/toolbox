use std::path::Path;
use std::process::Command;

pub fn build() {
    let cmd = "go build -o superserver -tags=jsoniter -ldflags '-w -s' && chmod +x superserver";
    let output = Command::new("bash")
        .current_dir(Path::new("../superserver/"))
        .arg("-c")
        .arg(cmd)
        .output().unwrap();
    if !output.status.success() {
        println!("{}", String::from_utf8(output.stderr).unwrap());
    }
}