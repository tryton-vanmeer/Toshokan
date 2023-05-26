use std::io;

use anyhow::Result;
use clap::{Command, CommandFactory, Parser, Subcommand};
use clap_complete::{generate, Shell};
use colored::Colorize;
use toshokan::{get_games, get_game};

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

    /// View info for specified game
    Info {
        /// Game to view info about
        appid: u32,
    },
}

fn generate_completions(shell: Shell, cmd: &mut Command) -> Result<()> {
    generate(shell, cmd, cmd.get_name().to_string(), &mut io::stdout());

    Ok(())
}

fn list() -> Result<()> {
    println!(
        "{}",
        format!("{:<8} {}{}", "AppID".bold(), "Name".bold(), " ".repeat(32)).underline()
    );
    for game in get_games()? {
        println!("{:<8} {}", game.appid.to_string().green(), game.name);
    }

    Ok(())
}

fn info(appid: u32) -> Result<()> {
    let game = get_game(appid)?;

    println!("{:#?}", game);

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

        Commands::Info { appid } => info(appid)?,
    }

    Ok(())
}
