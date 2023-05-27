mod cli;
mod protondb;
mod steam;

use anyhow::Result;

fn main() -> Result<()> {
    cli::run()
}
