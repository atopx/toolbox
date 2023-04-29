use std::path::Path;
use std::process::Command;

pub fn build() {
    let cmd = "rm -rf ../protocol/api/swagger.yaml && swag init --output ../protocol/api --outputTypes yaml --parseInternal";
    let output = Command::new("bash")
        .current_dir(Path::new("../superserver/"))
        .arg("-c")
        .arg(cmd)
        .output().unwrap();
    if !output.status.success() {
        println!("{}", String::from_utf8(output.stderr).unwrap());
    }
}