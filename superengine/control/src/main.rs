use clap::{command, Parser, Subcommand};

mod builds;

/// Super控制器
#[derive(Parser, Debug)]
#[command(name = "control")]
#[command(author, version, about, long_about = None)]
struct Control {
    #[command(subcommand)]
    command: Commands,
}

#[derive(Debug, Subcommand)]
enum Commands {
    /// 构建命令
    #[command(arg_required_else_help = true)]
    Build {
        /// 构建对象
        code: builds::Code,
    }
}


fn main() {
    let args = Control::parse();
    match args.command {
        Commands::Build { code } => builds::run(code),
    }
}