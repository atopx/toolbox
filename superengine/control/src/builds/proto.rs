use std::{env, fs, path::Path, vec};
use std::process::Command;

pub fn build() {
    build_rs();
    build_go();
}

fn build_go() {
    let cmd = "rm -rf domain && protoc --proto_path ../protocol/proto --go-grpc_out=../ --go_out=../ `find ../protocol/ -name *.proto`";
    let output = Command::new("bash")
        .current_dir(Path::new("../superserver/"))
        .arg("-c")
        .arg(cmd)
        .output().unwrap();
    if !output.status.success() {
        println!("{}", String::from_utf8(output.stderr).unwrap());
    }
}


fn build_rs() {
    let current_dir = env::current_dir().unwrap();
    println!("{:?}", current_dir);
    let base_path = Path::new(&current_dir).join("../protocol/proto");

    let mut proto_files = vec![];
    let mut proto_paths = vec![];
    proto_paths.push(base_path.clone());

    for entry in fs::read_dir(&base_path).unwrap() {
        let entry = entry.unwrap();
        let md = entry.metadata().unwrap();
        if md.is_dir() {
            proto_paths.push(entry.path());
            continue;
        }
    }

    for proto_path in proto_paths {
        for entry in fs::read_dir(&proto_path).unwrap() {
            let entry = entry.unwrap();
            if entry.metadata().unwrap().is_file() && entry.path().extension().unwrap() == "proto" {
                proto_files.push(entry.path().as_os_str().to_os_string())
            }
        }
    }

    tonic_build::configure()
        .out_dir("domain/src") // 生成代码的存放目录
        .build_client(false) // 不生成客户端
        .build_server(true)
        .compile(
            proto_files.as_slice(), // 欲生成的 proto 文件列表
            &[&base_path],         // proto 依赖所在的根目录
        )
        .unwrap();
}
