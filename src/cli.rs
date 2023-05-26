use anyhow::Result;
use clap::Parser;

#[derive(Parser, Debug)]
#[command(author, version)]
/// Tool for interacting with your Steam library.
struct Args {}

pub fn run() -> Result<()> {
    let args = Args::parse();

    Ok(())
}
