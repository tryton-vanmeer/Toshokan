mod cli;
mod steam;
mod protondb;

use anyhow::Result;

fn main() -> Result<()> {
    cli::run()
}
