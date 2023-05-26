use std::env;

use anyhow::{Context, Ok, Result};
use steamlocate::{SteamApp, SteamDir};

pub struct Game {
    pub appid: u32,
    pub name: String,
    path: String,
}

impl Game {
    fn from_steamapp(id: &u32, app: &SteamApp) -> Self {
        Self {
            appid: *id,
            name: app.name.as_ref().unwrap().clone(),
            path: app.path.display().to_string(),
        }
    }

    fn path(&self) -> String {
        let home = env::var("HOME").unwrap();

        if self.path.starts_with(&home) {
            return self.path.replace(&home, "~");
        }

        self.path.clone()
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
