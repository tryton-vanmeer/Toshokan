use std::io;

use anyhow::Result;
use clap::{Command, CommandFactory, Parser, Subcommand};
use clap_complete::{generate, Shell as CompletionShell};

#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
struct Cli {
    #[command(subcommand)]
    cmd: Commands,
}

#[derive(Subcommand, Debug)]
enum Commands {
    /// Generate toshokan shell completions for your shell to stdout
    Completions {
        #[clap(value_enum)]
        shell: CompletionShell,
    },
    /// List installed games in your Steam library
    List,
}

fn generate_completions(shell: CompletionShell, cmd: &mut Command) -> Result<()> {
    generate(shell, cmd, cmd.get_name().to_string(), &mut io::stdout());

    Ok(())
}

fn list() -> Result<()> {
    println!("{}", toshokan::get_games());
    Ok(())
}

pub fn run() -> Result<()> {
    let args = Cli::parse();

    match args.cmd {
        Commands::Completions { shell } => generate_completions(shell, &mut Cli::command())?,
        Commands::List => list()?,
    }

    Ok(())
}
