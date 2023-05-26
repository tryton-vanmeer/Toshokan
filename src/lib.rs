use std::{env, path::PathBuf};

use anyhow::{Context, Ok, Result};
use steamlocate::{SteamApp, SteamDir};

trait DisplayHomeAsTilde {
    fn display_home_as_tilde(&self) -> String;
}

impl DisplayHomeAsTilde for PathBuf {
    fn display_home_as_tilde(&self) -> String {
        let path = self.display().to_string();

        let home = env::var("HOME").unwrap();
        if path.starts_with(&home) {
            return path.replace(&home, "~");
        }

        path
    }
}

#[derive(Debug)]
pub struct Game {
    pub appid: u32,
    pub name: String,
    pub proton: bool,
    path: PathBuf,
}

impl Game {
    fn from_steamapp(id: &u32, app: &SteamApp) -> Self {
        let proton = match app.user_config.get("platform_override_source") {
            Some(platform) => platform == "windows",
            _ => false,
        };

        Self {
            appid: *id,
            name: app.name.as_ref().unwrap().clone(),
            proton,
            path: app.path.clone(),
        }
    }

    fn should_filter(&self) -> bool {
        ![
            "Proton 6.3",
            "Proton 7.0",
            "Proton 8.0",
            "Proton EasyAntiCheat Runtime",
            "Proton Experimental",
            "Proton Hotfix",
            "Steamworks Common Redistributables",
            "Steam Linux Runtime - Sniper",
            "Steam Linux Runtime - Soldier",
        ]
        .contains(&self.name.as_ref())
    }

    pub fn path(&self) -> String {
        self.path.display_home_as_tilde()
    }
}

pub fn get_games() -> Result<Vec<Game>> {
    let mut steam = SteamDir::locate().context("unable to find steamdir")?;

    let mut games: Vec<Game> = steam
        .apps()
        .iter()
        .map(|(id, app)| Game::from_steamapp(id, app.as_ref().unwrap()))
        .filter(|game| game.should_filter())
        .collect::<Vec<Game>>();

    games.sort_by_key(|game| game.name.to_string());

    Ok(games)
}

pub fn get_game(appid: u32) -> Result<Game> {
    let mut steam = SteamDir::locate().context("unable to find steamdir")?;

    Ok(Game::from_steamapp(
        &appid,
        steam
            .app(&appid)
            .context(format!("unable to find game with appid {}", appid))?,
    ))
}
