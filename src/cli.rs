use crate::protondb;
use crate::steam;

use std::io;

use anyhow::Result;
use clap::{Command, CommandFactory, Parser, Subcommand};
use clap_complete::{generate, Shell};
use colored::Colorize;

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
        "{}\t{}",
        "AppID".bold().underline(),
        "Name".bold().underline()
    );
    for game in steam::get_games()? {
        println!("{}\t{}", game.appid.to_string().bright_green(), game.name);
    }

    Ok(())
}

fn print_info(game: steam::Game) {
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

        let protondb = protondb::Summary::from_appid(game.appid);

        match protondb {
            Ok(summary) => println!("{:<width$} {}", "proton rating".blue().bold(), summary),
            Err(err) => eprint!("{}", err),
        };
    }
}

fn info(appid: u32) -> Result<()> {
    print_info(steam::get_game(appid)?);

    Ok(())
}

fn info_all() -> Result<()> {
    let games = steam::get_games()?;
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
