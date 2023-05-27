mod cli;
mod protondb;

use anyhow::Result;

fn main() -> Result<()> {
    cli::run()
}
