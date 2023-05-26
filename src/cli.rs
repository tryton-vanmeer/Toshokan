use std::io;

use anyhow::{Ok, Result};
use clap::{Command, CommandFactory, Parser, Subcommand};
use clap_complete::{generate, Shell};
use colored::Colorize;
use toshokan::{get_game, get_games, Game};

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
        #[arg(required_unless_present = "all")]
        appid: Option<u32>,
        /// Show info for all games
        #[arg(short, long)]
        all: bool,
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
        println!("{:<8} {}", game.appid.to_string().green().bold(), game.name);
    }

    Ok(())
}

fn print_info(game: Game) {
    let width = 13;

    println!("{}", game.name.purple().bold());

    println!(
        "{:<width$} https://store.steampowered.com/app/{}",
        "store".blue().bold(),
        game.appid
    );

    println!("{:<width$} {}", "install dir".blue().bold(), game.path());

    if game.proton {
        println!(
            "{:<width$} {}",
            "proton prefix".blue().bold(),
            game.proton_prefix()
        );
    }
}

fn info(appid: u32) -> Result<()> {
    print_info(get_game(appid)?);

    Ok(())
}

fn info_all() -> Result<()> {
    let games = get_games()?;
    let mut games_peeker = games.iter().peekable();

    while let Some(game) = games_peeker.next() {
        print_info(game.to_owned());

        if games_peeker.peek().is_some() {
            println!();
        }
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

        Commands::Info { appid, all } => {
            if all {
                info_all()?
            } else {
                info(appid.unwrap())?
            }
        }
    }

    Ok(())
}
