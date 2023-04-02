use reqwest::blocking::Client;
use std::{
    fs,
    io::{BufRead, BufReader, Error, Write},
    path,
};
use super::super::config::config::{UA, BASE_DIR};


#[derive(Clone, serde::Serialize)]
pub struct NovelMessge {
    has_error: bool,
    content: String,
}

impl NovelMessge {
    fn new(has_error: bool, content: String) -> Self {
        return Self { has_error, content };
    }
}

#[tauri::command]
pub fn novel_fetch(link: &str) -> NovelMessge {
    let client = Client::new();
    let resp = client.get(link).header("user-agent", UA).send();
    match resp {
        Ok(data) => NovelMessge::new(false, data.text().unwrap()),
        Err(e) => NovelMessge::new(true, e.to_string()),
    }
}

#[tauri::command]
pub fn novel_compose(id: &str, index: i32, title: &str, content: &str) {
    fs::create_dir_all(format!("{BASE_DIR}/cache/{id}")).unwrap();
    // 清理广告
    let mut result = content.trim().replace("\u{3000}", "\n");
    result = result.replace("天才一秒记住本站地址：", "");
    result = result.replace("/最快更新！无广告！", "");
    result = result.replace("本站笔趣阁最新域名：www.1biqug.net", "");
    result = result.replace("chaptererror();", "");
    // 写入章节文件
    let mut output = fs::File::create(format!("{BASE_DIR}/cache/{id}/{index:0>6}.txt")).unwrap();
    write!(output, "{}\n{}\n", title, result).unwrap();
}

#[tauri::command]
pub fn novel_save(id: &str) -> NovelMessge {
    let target_dir_name = format!("{BASE_DIR}/cache/{id}");
    let target_dir_path = path::Path::new(target_dir_name.as_str());
    if !target_dir_path.exists() {
        return NovelMessge::new(true, format!("not such dir with {target_dir_name}"));
    };

    // 获取文件列表
    let mut entries = target_dir_path
        .read_dir()
        .unwrap()
        .map(|res| res.map(|e| e.path()))
        .collect::<Result<Vec<_>, Error>>()
        .unwrap();
    // 文件列表排序
    entries.sort();

    // 写入主文件
    let target_file_name = format!("{BASE_DIR}/{id}.txt");
    let mut file = fs::File::create(&target_file_name).unwrap();
    for entrie in entries {
        let buffer = BufReader::new(fs::File::open(entrie).unwrap());
        for line in buffer.lines() {
            write!(file, "{}\n", line.unwrap()).unwrap();
        }
    }
    return NovelMessge::new(false, target_file_name);
}

#[tauri::command]
pub fn novel_clean(id: &str) {
    fs::remove_dir_all(format!("{BASE_DIR}/cache/{id}")).unwrap();
    // fs::remove_file(format!("{BASE_DIR}/{id}.txt")).unwrap();
}
