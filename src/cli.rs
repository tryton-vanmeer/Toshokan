use std::io;

use anyhow::Result;
use clap::{Command, CommandFactory, Parser, Subcommand};
use clap_complete::{generate, Shell};
use toshokan::get_games;

#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
struct Cli {
    #[command(subcommand)]
    cmd: Commands,
}

#[derive(Subcommand, Debug)]
enum Commands {
    /// Generate shell completions. Default to current shell
    Completions {
        /// Infer current shell when missing, fallback to bash
        #[clap(value_enum)]
        shell: Option<Shell>,
    },

    /// List installed games in your Steam library
    List,
}

fn generate_completions(shell: Shell, cmd: &mut Command) -> Result<()> {
    generate(shell, cmd, cmd.get_name().to_string(), &mut io::stdout());

    Ok(())
}

fn list() -> Result<()> {
    for game in get_games()? {
        println!("{}", game.name);
    }

    Ok(())
}

pub fn run() -> Result<()> {
    let args = Cli::parse();

    match args.cmd {
        Commands::Completions { shell } => {
            let gen = match shell {
                Some(s) => s,
                None => Shell::from_env().unwrap_or(Shell::Bash),
            };

            generate_completions(gen, &mut Cli::command())?
        }

        Commands::List => list()?,
    }

    Ok(())
}
